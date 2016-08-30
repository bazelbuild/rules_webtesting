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
"""web_test_config.bzl defines the web_test_config rule.

The web_test_config rules provides configuration information such as
whether to record video to web_test.

DO NOT load this file. Use "@io_bazel_rules_web//web:web.bzl".
"""

load("//rules:metadata.bzl", "create_file", "merge_files")


def _web_test_config_impl(ctx):
  """Implementation of the web_test_config rule."""

  metadata_files = []
  for dep in ctx.attr.data:
    if hasattr(dep, "web_test_metadata"):
      metadata_files += [dep.web_test_metadata]

  for config in ctx.attr.configs:
    metadata_files += [config.web_test_metadata]

  patch = ctx.new_file("%s.tmp.json" % ctx.label.name)
  create_file(ctx=ctx, output=patch)
  metadata_files += [patch]

  merge_files(
      ctx=ctx,
      merger=ctx.executable._merger,
      output=ctx.outputs.web_test_metadata,
      inputs=metadata_files)

  return struct(web_test_metadata=ctx.outputs.web_test_metadata)


web_test_config = rule(
    implementation=_web_test_config_impl,
    attrs={
        "configs": attr.label_list(providers=["json", "record"]),
        "data": attr.label_list(
            allow_files=True, cfg=DATA_CFG),
        "_merger": attr.label(
            executable=True,
            cfg=HOST_CFG,
            default=Label("//external:web_test_merger")),
    },
    outputs={"web_test_metadata": "%{name}.gen.json"},)
