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

load(
    "//web/internal:shared.bzl",
    "path",)


def _web_test_archive_impl(ctx):
  content = '{\n  "webTestFiles": [{\n    "archiveFile": "%s"' % path(
      ctx, ctx.file.archive)
  if ctx.attr.named_files:
    content += ',\n    "namedFiles": {'
    first = True
    for k, v in ctx.attr.named_files.items():
      if first:
        first = False
      else:
        content += ","
      content += '\n      "%s": "%s"' % (k, v)
    content += "\n    }"
  content += "\n  }]\n}"

  ctx.file_action(
      output=ctx.outputs.web_test_metadata, content=content, executable=False)
  return struct(
      runfiles=ctx.runfiles(
          files=[ctx.file.archive, ctx.outputs.web_test_metadata]),
      web_test_metadata=ctx.outputs.web_test_metadata)


web_test_archive = rule(
    implementation=_web_test_archive_impl,
    attrs={
        "archive":
            attr.label(
                allow_single_file=True, cfg="data", mandatory=True),
        "named_files":
            attr.string_dict(mandatory=True)
    },
    outputs={"web_test_metadata": "%{name}.gen.json"},)
"""Specifies an archive file with named files in it.

The archive will be unzipped only if Web Test Launcher wants one the named
files in the archive.

Args:
  archive: label referring to the archive file.
  named_files: a map of names used by Web Test Launcher to path in the archive.
"""
