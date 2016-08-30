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

load("//rules:shared.bzl", "build_runfiles")
load("//rules:metadata.bzl", "create_file", "merge_files")


def _web_test_named_executable_impl(ctx):
  name = ctx.attr.alt_name or ctx.label.name
  create_file(
      ctx=ctx,
      output=ctx.outputs.web_test_metadata,
      named_executables={name: ctx.executable.executable},)

  return struct(
      runfiles=build_runfiles(
          ctx=ctx,
          files=[ctx.outputs.web_test_metadata],
          deps_attrs=["executable"],),
      web_test_metadata=ctx.outputs.web_test_metadata,)


web_test_named_executable = rule(
    implementation=_web_test_named_executable_impl,
    attrs={
        "alt_name": attr.string(),
        "executable": attr.label(
            allow_files=True, executable=True, cfg=DATA_CFG),
        "data": attr.label_list(
            allow_files=True, cfg=DATA_CFG),
    },
    outputs={"web_test_metadata": "%{name}.gen.json"},)
