# Copyright 2016 Google Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
"""A library of functions for working with web_test metadata files."""

load('//web/internal:collections.bzl', 'lists')
load('//web/internal:files.bzl', 'files')


def _merge_files(ctx, merger, output, inputs):
  """Takes a list of input metadata files, and produces a merged output file.

  Args:
    ctx: a skylark rule context
    merger: the WTL metadata merger executable
    output: a File object for the output file
    inputs: a list of File objects for the input files
  """
  paths = [i.path for i in inputs]
  short_paths = [i.short_path for i in inputs]
  args = ['--output', output.path] + paths

  ctx.action(
      outputs=[output],
      inputs=inputs,
      executable=merger,
      arguments=args,
      mnemonic='METADATAMERGER',
      progress_message='merging %s' % (', '.join(short_paths)))


def _create_file(ctx,
                 output,
                 capabilities=None,
                 environment=None,
                 browser_label=None,
                 test_label=None,
                 web_test_files=None):
  """Generates a web_test metadata file with specified contents."""
  content = '{\n  "_comment": "generated file for %s"' % ctx.label
  if capabilities:
    content += ',\n  "capabilities": ' + capabilities
  if environment:
    content += ',\n  "environment": "' + environment + '"'
  if browser_label:
    content += ',\n  "browserLabel": "%s"' % browser_label
  if test_label:
    content += ',\n  "testLabel": "%s"' % test_label
  if web_test_files:
    content += ',\n  "webTestFiles": ['
    first = True
    for f in web_test_files:
      if first:
        first = False
      else:
        content += ','
      content += '\n' + _web_test_files_to_string(ctx, f, '    ')
    content += '\n  ]'
  content += '\n}\n'

  ctx.file_action(output=output, content=content, executable=False)


def _get_files(ctx, attr_names):
  """Finds all of the web_test_metadata files in given attributes.

  It ensures that all web_test_metadata files position in the returned list is
  based on the last attribute in attr_names that they appear in.

  Args:
    ctx: the context object for this rule.
    attr_names: the names of attributes to get web_test_metadata files from.
  Returns:
    a list of all of the web_test_metadata files in the specificied attributes.
  """
  metadata_files = []

  for attr_name in attr_names:
    if hasattr(ctx.attr, attr_name):
      attr = getattr(ctx.attr, attr_name)
      if lists.is_list_like(attr):
        for item in attr:
          _get_files_helper(metadata_files, item)
      else:
        _get_files_helper(metadata_files, attr)
  return metadata_files


def _get_files_helper(metadata_files, r):
  if hasattr(r, 'transitive_web_test_metadata'):
    lists.ensure_all_at_end_of_list(metadata_files,
                                    r.transitive_web_test_metadata)
  if hasattr(r, 'web_test_metadata'):
    lists.ensure_all_at_end_of_list(metadata_files, r.web_test_metadata)


def _web_test_files(named_files={}, archive_file=None):
  return struct(archive_file=archive_file, named_files=named_files)


def _web_test_files_to_string(ctx, files_struct, indent=''):
  """Converts a web_test_files struct to a JSON string."""
  result = indent + '{\n'
  if files_struct.archive_file:
    result += indent + '  "archiveFile": "' + files.runfiles_path(
        ctx, files_struct.archive_file) + '",\n'
  result += indent + '  "namedFiles": {'
  first = True
  for n, f in files_struct.named_files.items():
    if first:
      first = False
    else:
      result += ','
    result += '\n' + indent + '    "' + n + '": "'
    if type(f) == type(''):
      result += f
    else:
      result += files.runfiles_path(ctx, f)
    result += '"'
  result += '\n' + indent + '  }\n' + indent + '}'
  return result


def _aspect_impl(target, ctx):
  transitive_files = _get_files(ctx.rule, ['data', 'deps', 'srcs'])

  if hasattr(target, 'web_test_metadata'):
    for m in target.web_test_metadata:
      lists.ensure_at_end_of_list(transitive_files, m)

  return struct(transitive_web_test_metadata=transitive_files)


_aspect = aspect(
    implementation=_aspect_impl,
    attr_aspects=['deps', 'data', 'srcs', 'browser', 'config'],)

metadata = struct(
    aspect=_aspect,
    create_file=_create_file,
    get_files=_get_files,
    merge_files=_merge_files,
    web_test_files=_web_test_files,)
