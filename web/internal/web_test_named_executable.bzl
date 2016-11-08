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
"""Defines a executable that can be located by name."""

load("//web/internal:collections.bzl", "lists")
load("//web/internal:files.bzl", "files")
load("//web/internal:metadata.bzl", "metadata")


def _web_test_named_executable_impl(ctx):
  """Implementation of web_test_named_executable."""
  data_labels = [datum.label for datum in ctx.attr.data]

  if ctx.attr.executable.label not in data_labels:
    fail("Executable %s must be in data attribute" % ctx.attr.executable.label,
         "data")

  name = ctx.attr.alt_name or ctx.label.name
  metadata.create_file(
      ctx,
      output=ctx.outputs.web_test_metadata,
      web_test_files=[
          metadata.web_test_files({
              name: ctx.executable.executable
          })
      ])

  return struct(
      runfiles=files.runfiles(
          ctx=ctx, files=ctx.files.executable, deps_attrs=["executable"]),
      web_test_metadata=[ctx.outputs.web_test_metadata])


web_test_named_executable = rule(
    attrs={
        "alt_name":
            attr.string(),
        "executable":
            attr.label(
                allow_files=True,
                executable=True,
                cfg="data",
                mandatory=True,
                aspects=[metadata.aspect]),
        "data":
            attr.label_list(
                allow_files=True, cfg="data", aspects=[metadata.aspect]),
    },
    outputs={"web_test_metadata": "%{name}.gen.json"},
    implementation=_web_test_named_executable_impl,)
"""Defines a executable that can be located by name.

Args:
  alt_name: If supplied, is used instead of name to lookup the executable.
  executable: The executable that will be returned for name or alt_name.
  data: Runtime dependencies for the executable.
"""
