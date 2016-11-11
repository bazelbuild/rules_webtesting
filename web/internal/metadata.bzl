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
  """Produces a merged web test metadata file.

  Args:
    ctx: a Skylark rule context.
    merger: the WTL metadata merger executable.
    output: a File object for the output file.
    inputs: a sequence of File objects. These files are in order of priority; 
      i.e. values in the first file will take precedence over values in the 
      second file, etc.
  """
  inputs = list(inputs)
  paths = [i.path for i in reversed(inputs)]
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
  """Generates a web_test metadata file with specified contents.

  Args:
    ctx: a Skylark rule context.
    output: File object. The file to write the metadata to.
    environment: string. The Web Test Launcher environment name.
    browser_label: Label. The label for a browser rule.
    test_label: Label. The label for the test being executed.
    web_test_files: sequence of web_test_file structs. The named files needed
      for this configuration.
  """
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


def _web_test_files_to_string(ctx, files_struct, indent=''):
  """Converts a web_test_files struct to a JSON string.

  Args:
    ctx: a Skylark rule context.
    files_struct: a web_test_files struct.
    indent: string. the base indentation that should be used when writing
      generating the JSON-encoded string.
  Returns:
    a JSON-encoded string representing the web_test_files object.
  """
  result = indent + '{\n'
  if files_struct.archive_file:
    result += indent + '  "archiveFile": "' + files.long_path(
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
      result += files.long_path(ctx, f)
    result += '"'
  result += '\n' + indent + '  }\n' + indent + '}'
  return result


def _get_files(ctx, attr_names):
  """Finds all of the web_test_metadata files in given attributes.

  It ensures that all web_test_metadata files position in the returned list is
  based on the last attribute in attr_names that they appear in.

  Args:
    ctx: a Skylark rule context.
    attr_names: sequence of strings. The names of attributes to get 
      web_test_metadata files from. Attributes earlier in the list will
      return files earlier in the result.
  Returns:
    A sequence of all of the web_test_metadata files in the specificied 
    attributes.
  """
  metadata_files = set()

  for attr_name in reversed(attr_names):
    attr = getattr(ctx.attr, attr_name)
    if lists.is_list_like(attr):
      for a in reversed(attr):
        mf = _get_metadata_if_present(a)
        if mf:
          metadata_files = metadata_files | set([mf])
    else:
      mf = _get_metadata_if_present(attr)
      if mf:
        metadata_files = metadata_files | set([mf])

  return metadata_files


def _get_metadata_if_present(attr):
  """Gets attr.web_test.metadata if it exists.

  Args:
    attr: the object to get the web_test.metadata from.
  Returns:
    The File object that web_test.metadata refers to if present.
    Otherwise None.
  """
  if hasattr(attr, 'web_test') and hasattr(attr.web_test, 'metadata'):
    return attr.web_test.metadata
  return None


def _web_test_files(archive_file=None, named_files={}):
  """Build a web_test_files struct.

  Args:
    archive_file: a File object. The archive file where the named_files will be
      found. If absent, the named_files are located directly in the runfiles.
    named_files: a dict of strings to strings or File objects. The mapping of
      names to file path. If archive_file is absent, the values should be
      File objects for files that will be in the runfiles of the test. If 
      archive_file is present, the values should be string paths referencing
      files in archive_file.
  Returns:
    A web_test_files struct.
  """
  return struct(archive_file=archive_file, named_files=named_files)


metadata = struct(
    create_file=_create_file,
    get_files=_get_files,
    merge_files=_merge_files,
    web_test_files=_web_test_files)
