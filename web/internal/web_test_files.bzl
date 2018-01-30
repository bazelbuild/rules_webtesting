# Copyright 2017 Google Inc.
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
"""A rule for configuring a set of named files.

DO NOT load this file. Use "@io_bazel_rules_web//web:web.bzl".
"""

load("//web/internal:metadata.bzl", "metadata")
load("//web/internal:provider.bzl", "WebTestInfo")


def _web_test_files_impl(ctx):
  named_files = {}
  files = depset()

  for target, name in ctx.attr.files.items():
    if name in named_files:
      fail("%s appears multiple times." % name, "files")
    if len(target.files) != 1:
      fail("%s refers to multiple files." % target.label, "files")
    named_files[name] = target.files.to_list()[0]
    files = files + target.files

  patch = ctx.new_file("%s.tmp.json" % ctx.label.name)
  metadata.create_file(
      ctx=ctx,
      output=patch,
      web_test_files=[
          metadata.web_test_files(ctx=ctx, named_files=named_files),
      ])

  metadata_files = [patch
                   ] + [dep[WebTestInfo].metadata for dep in ctx.attr.deps]

  metadata.merge_files(
      ctx=ctx,
      merger=ctx.executable.merger,
      output=ctx.outputs.web_test_metadata,
      inputs=metadata_files)

  return [
      DefaultInfo(
          runfiles=ctx.runfiles(
              collect_data=True, collect_default=True, files=files.to_list())),
      WebTestInfo(metadata=ctx.outputs.web_test_metadata),
  ]


web_test_files = rule(
    implementation=_web_test_files_impl,
    attrs={
        "data":
            attr.label_list(allow_files=True, cfg="data"),
        "deps":
            attr.label_list(providers=[WebTestInfo]),
        "merger":
            attr.label(
                executable=True,
                cfg="host",
                default=Label("//go/metadata/main")),
        "files":
            attr.label_keyed_string_dict(
                mandatory=True, allow_files=True, allow_empty=False),
    },
    outputs={"web_test_metadata": "%{name}.gen.json"},
)
"""Specifies a set of named files.

Args:
  data: Runtime dependencies for this rule.
  deps: Other web_test-related rules that this rule depends on.
  merger: Metadata merger executable.
  files: A map of files to names.
"""
