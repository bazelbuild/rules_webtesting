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

load("//web/internal:shared.bzl",
     "build_runfiles",
     "get_metadata_files",
     "merge_metadata_files",
     "path",)


def _web_test_named_executable_impl(ctx):
  name = ctx.attr.alt_name or ctx.label.name

  metadata_files = get_metadata_files(ctx, ["data"])

  if metadata_files:
    patch = ctx.new_file("%s.tmp.json" % ctx.label.name)
  else:
    patch = ctx.outputs.web_test_metadata

  content = """{
  "namedFiles": {"%s": "%s"}
}""" % (name, path(ctx, ctx.executable.executable))

  ctx.file_action(output=patch, content=content, executable=False)

  if metadata_files:
    metadata_files += [patch]
    merge_metadata_files(
        ctx=ctx,
        merger=ctx.executable._merger,
        output=ctx.outputs.web_test_metadata,
        inputs=metadata_files,)

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
            allow_files=True, executable=True, cfg="data", mandatory=True),
        "data": attr.label_list(
            allow_files=True, cfg="data"),
        "_merger": attr.label(
            executable=True,
            cfg="host",
            default=Label("//external:web_test_merger")),
    },
    outputs={"web_test_metadata": "%{name}.gen.json"},)
"""Defines a executable that can be located by name.

Args:
  alt_name: If supplied, is used instead of name to lookup the executable.
  executable: The executable that will be returned for name or alt_name.
  data: Runtime dependencies for the executable.
"""
