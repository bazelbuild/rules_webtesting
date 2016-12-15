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
"""A rule for defining executables that can be located by name.

DO NOT load this file. Use "@io_bazel_rules_web//web:web.bzl".
"""

load("//web/internal:metadata.bzl", "metadata")


def _web_test_named_executable_impl(ctx):
  if ctx.attr.executable.label not in [data.label for data in ctx.attr.data]:
    fail("Executable %s must be in data." % ctx.attr.executable.label, "data")

  name = ctx.attr.alt_name or ctx.label.name

  patch = ctx.new_file("%s.tmp.json" % ctx.label.name)
  metadata.create_file(
      ctx=ctx,
      output=patch,
      web_test_files=[
          metadata.web_test_files(
              ctx=ctx, named_files={name: ctx.executable.executable}),
      ])
  metadata_files = [patch] + [dep.web_test.metadata for dep in ctx.attr.deps]

  metadata.merge_files(
      ctx=ctx,
      merger=ctx.executable.merger,
      output=ctx.outputs.web_test_metadata,
      inputs=metadata_files)

  return struct(
      runfiles=ctx.runfiles(
          collect_data=True, collect_default=True),
      web_test=struct(metadata=ctx.outputs.web_test_metadata))


web_test_named_executable = rule(
    attrs={
        "alt_name":
            attr.string(),
        "data":
            attr.label_list(
                allow_files=True, cfg="data"),
        "deps":
            attr.label_list(providers=["web_test"]),
        "executable":
            attr.label(
                allow_files=True, executable=True, cfg="data", mandatory=True),
        "merger":
            attr.label(
                executable=True,
                cfg="host",
                default=Label("//go/metadata:merger")),
    },
    outputs={"web_test_metadata": "%{name}.gen.json"},
    implementation=_web_test_named_executable_impl,)
"""Defines a executable that can be located by name.

Args:
  alt_name: If supplied, is used instead of name to lookup the executable.
  data: Runtime dependencies for the executable.
  deps: Other web_test-related rules that this rule depends on.
  executable: The executable that will be returned for name or alt_name.
  merger: Metadata merger executable.
"""
