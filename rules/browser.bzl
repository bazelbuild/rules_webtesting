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
"""The browser rule is used to define web_test browsers.

DO NOT load this file. Use "@io_bazel_rules_web//web:web.bzl".
"""

load("//rules:shared.bzl", "browser_struct", "build_runfiles")
load("//rules:metadata.bzl", "create_file", "merge_files")


def _browser_impl(ctx):
  """Implementation of the browser rule."""
  patch = ctx.new_file("%s.tmp.json" % ctx.label.name)
  create_file(ctx=ctx, output=patch, browser_label=ctx.label)
  merge_files(
      ctx=ctx,
      merger=ctx.executable._merger,
      output=ctx.outputs.web_test_metadata,
      inputs=[ctx.file.metadata, patch])

  required_tags = set(ctx.attr.required_tags)
  required_tags += ["browser-" + ctx.label.name]

  return browser_struct(
      disabled=ctx.attr.disabled,
      required_tags=required_tags,
      runfiles=build_runfiles(
          ctx, files=[ctx.outputs.web_test_metadata]),
      web_test_metadata=ctx.outputs.web_test_metadata,)


browser = rule(
    implementation=_browser_impl,
    attrs={
        "metadata": attr.label(
            mandatory=True, allow_files=True, single_file=True, cfg=DATA_CFG),
        "data": attr.label_list(
            allow_files=True, cfg=DATA_CFG),
        "disabled": attr.string(),
        "required_tags": attr.string_list(default=[]),
        "_merger": attr.label(
            executable=True,
            cfg=HOST_CFG,
            default=Label("//external:web_test_merger")),
    },
    outputs={"web_test_metadata": "%{name}.gen.json"},)
