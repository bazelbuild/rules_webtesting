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
"""A rule for configuring archives that contain named files.

DO NOT load this file. Use "@io_bazel_rules_web//web:web.bzl".
"""

load("//web/internal:metadata.bzl", "metadata")
load("//web/internal:provider.bzl", "WebTestInfo")
load("//web/internal:runfiles.bzl", "runfiles")


def _web_test_archive_impl(ctx):
  metadata.create_file(
      ctx=ctx,
      output=ctx.outputs.web_test_metadata,
      web_test_files=[
          metadata.web_test_files(
              ctx=ctx,
              archive_file=ctx.file.archive,
              named_files=ctx.attr.named_files),
      ])

  return [
      DefaultInfo(runfiles=ctx.runfiles(files=[ctx.file.archive])),
      WebTestInfo(metadata=ctx.outputs.web_test_metadata),
  ]


web_test_archive = rule(
    doc="""Specifies an archive file with named files in it.

The archive will be unzipped only if Web Test Launcher wants one the named
files in the archive.""",
    implementation=_web_test_archive_impl,
    attrs={
        "archive":
            attr.label(
                doc="Archive file that contains named files.",
                allow_single_file=[
                    ".deb",
                    ".tar",
                    ".tar.bz2",
                    ".tbz2",
                    ".tar.gz",
                    ".tgz",
                    ".tar.Z",
                    ".zip",
                ],
                mandatory=True),
        "named_files":
            attr.string_dict(
                doc="A map of names to paths in the archive.", mandatory=True),
    },
    outputs={"web_test_metadata": "%{name}.gen.json"},
)
