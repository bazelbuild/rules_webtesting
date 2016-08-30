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

load('//rules:shared.bzl', 'path')


def merge_files(ctx, merger, output, inputs):
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


def create_file(ctx,
                output,
                capabilities=None,
                form_factor=None,
                browser_name=None,
                environment=None,
                browser_label=None,
                test_label=None,
                crop_screenshots=None,
                record_video=None,
                named_executables=None):
  """Generates a web_test metadata file with specified contents."""
  content = '{\n  "_comment": "generated file for %s"' % ctx.label
  if capabilities:
    content += ',\n  "capabilities": ' + capabilities
  if form_factor:
    content += ',\n  "formFactor": "' + form_factor + '"'
  if browser_name:
    content += ',\n  "browserName": "' + browser_name + '"'
  if environment:
    content += ',\n  "environment": "' + environment + '"'
  if browser_label:
    content += ',\n  "browserLabel": "%s"' % browser_label
  if test_label:
    content += ',\n  "testLabel": "%s"' % test_label
  if crop_screenshots == True:
    content += ',\n  "cropScreenshots": true'
  elif crop_screenshots == False:
    content += ',\n  "cropScreenshots": false'
  if record_video:
    content += ',\n  "recordVideo": "' + record_video + '"'
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
