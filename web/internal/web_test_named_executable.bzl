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
    "build_runfiles",
    "get_metadata_files",
    "merge_metadata_files",
    "path",)
load("//web/internal:web_test_metadata_aspect.bzl", "web_test_metadata_aspect")


def _web_test_named_executable_impl(ctx):
  name = ctx.attr.alt_name or ctx.label.name

  content = """{
  "webTestFiles": [{"namedFiles": {"%s": "%s"} }]
}""" % (name, path(ctx, ctx.executable.executable))

  ctx.file_action(output=ctx.outputs.web_test_metadata, content=content, executable=False)

  return struct(
      runfiles=build_runfiles(ctx=ctx, deps_attrs=["executable"]),
      web_test_metadata=[ctx.outputs.web_test_metadata])


web_test_named_executable = rule(
    attrs={
        "alt_name":
            attr.string(),
        "executable":
            attr.label(
                allow_files=True, executable=True, cfg="data", mandatory=True),
        "data":
            attr.label_list(
                allow_files=True, cfg="data", aspects=[web_test_metadata_aspect]),
    },
    outputs={"web_test_metadata": "%{name}.gen.json"},
    implementation=_web_test_named_executable_impl,)
"""Defines a executable that can be located by name.

Args:
  alt_name: If supplied, is used instead of name to lookup the executable.
  executable: The executable that will be returned for name or alt_name.
  data: Runtime dependencies for the executable.
"""
