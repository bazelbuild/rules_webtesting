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


def build_runfiles(ctx, deps_attrs=[], files=[]):
  """Creates a runfiles object.

  Args:
    ctx: the ctx object to create the runfiles object from.
    dep_attrs: sequence of strings; the attributes on ctx to
      transitively collect runfiles from. Note: runfiles from
      srcs, data, and deps attributes are always included.
    files: sequence of Files; list of files to be added to runfiles.
  Returns:
    A configured runfiles object.
  """
  transitive_files = set()
  for attr in deps_attrs:
    transitive_files += _get_transitive_files(getattr(ctx.attr, attr))
    if hasattr(ctx.files, attr):
      transitive_files += getattr(ctx.files, attr)
  return ctx.runfiles(
      files=files,
      transitive_files=transitive_files,
      collect_data=True,
      collect_default=True,)


def _get_transitive_files(attr_val):
  """Collects the set of transitive runfiles for the given attribute value.

  If the attribute is a sequence type, then it recurses into the values of
  sequence.

  Args:
    attr_value: value fo the attribute to collect runfiles from.
  Returns:
    set of Files.
  """
  transitive_files = set()
  if is_list_like(attr_val):
    for val in attr_val:
      transitive_files += _get_transitive_files(val)
    return transitive_files
  if hasattr(attr_val, 'data_runfiles'):
    transitive_files += attr_val.data_runfiles.files
  if hasattr(attr_val, 'default_runfiles'):
    transitive_files += attr_val.default_runfiles.files
  return transitive_files


def path(ctx, file):
  if file.owner and file.owner.workspace_root:
    return file.owner.workspace_root + '/' + file.short_path
  else:
    return ctx.workspace_name + '/' + file.short_path


def get_metadata_files(ctx, attr_names):
  metadata_files = []

  for attr_name in attr_names:
    attr = getattr(ctx.attr, attr_name)
    if hasattr(attr, 'web_test_metadata'):
      metadata_files += [attr.web_test_metadata]
    elif is_list_like(attr):
      metadata_files += [value.web_test_metadata for value in attr
                         if hasattr(value, 'web_test_metadata')]

  return metadata_files


def is_list_like(val):
  """Checks is val has list-like (list, set, tuple) value."""
  return type(val) in [type([]), type(set()), type(())]


def merge_metadata_files(ctx, merger, output, inputs):
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


def create_metadata_file(ctx,
                         output,
                         capabilities=None,
                         environment=None,
                         browser_label=None,
                         test_label=None,
                         named_executables=None):
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
  if named_executables:
    first = True
    content += ',\n  "namedExecutables": {'
    for k, v in named_executables.items():
      if first:
        first = False
      else:
        content += ','
      content += '\n    "' + k + '"' + ': "' + path(ctx, v) + '"'
    content += '\n  }'
  content += '\n}\n'

  ctx.file_action(output=output, content=content, executable=False)
