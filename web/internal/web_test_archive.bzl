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


def _web_test_archive_impl(ctx):
  patch = ctx.new_file("%s.tmp.json" % ctx.label.name)
  metadata.create_file(
      ctx=ctx,
      output=patch,
      web_test_files=[
          metadata.web_test_files(
              ctx=ctx,
              archive_file=ctx.file.archive,
              named_files=ctx.attr.named_files),
      ])

  metadata_files = [patch] + [dep.web_test.metadata for dep in ctx.attr.deps]

  metadata.merge_files(
      ctx=ctx,
      merger=ctx.executable.merger,
      output=ctx.outputs.web_test_metadata,
      inputs=metadata_files)

  return struct(
      runfiles=ctx.runfiles(
          collect_data=True, collect_default=True, files=[ctx.file.archive]),
      web_test=struct(metadata=ctx.outputs.web_test_metadata))


web_test_archive = rule(
    implementation=_web_test_archive_impl,
    attrs={
        "archive":
            attr.label(
                allow_single_file=True, cfg="data", mandatory=True),
        "data":
            attr.label_list(
                allow_files=True, cfg="data"),
        "deps":
            attr.label_list(providers=["web_test"]),
        "merger":
            attr.label(
                executable=True,
                cfg="host",
                default=Label("//go/metadata:merger")),
        "named_files":
            attr.string_dict(mandatory=True),
    },
    outputs={"web_test_metadata": "%{name}.gen.json"},)
"""Specifies an archive file with named files in it.

The archive will be unzipped only if Web Test Launcher wants one the named
files in the archive.

Args:
  archive: Archive file that contains named files.
  data: Runtime dependencies for this rule.
  deps: Other web_test-related rules that this rule depends on.
  merger: Metadata merger executable.
  named_files: A map of names to paths in the archive.
"""
