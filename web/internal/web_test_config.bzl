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

The web_test_config rules provides configuration information
such as additional capabilities.

DO NOT load this file. Use "@io_bazel_rules_web//web:web.bzl".
"""

load(":metadata.bzl", "metadata")
load(":provider.bzl", "WebTestInfo")

def _web_test_config_impl(ctx):
    """Implementation of the web_test_config rule."""
    metadata_files = []

    if ctx.attr.metadata:
        metadata_files = [ctx.file.metadata]

    metadata_files += [dep[WebTestInfo].metadata for dep in ctx.attr.deps]

    if metadata_files:
        metadata.merge_files(
            ctx = ctx,
            merger = ctx.executable.merger,
            output = ctx.outputs.web_test_metadata,
            inputs = metadata_files,
        )
    else:
        metadata.create_file(ctx = ctx, output = ctx.outputs.web_test_metadata)

    return [
        DefaultInfo(
            runfiles = ctx.runfiles(collect_data = True, collect_default = True),
        ),
        WebTestInfo(metadata = ctx.outputs.web_test_metadata),
    ]

web_test_config = rule(
    attrs = {
        "data": attr.label_list(
            doc = "Runtime dependencies for this configuration.",
            allow_files = True,
        ),
        "deps": attr.label_list(
            doc = "Other web_test-related rules that this rule depends on.",
            providers = [WebTestInfo],
        ),
        "merger": attr.label(
            doc = "The metadata merger binary.",
            default = Label("//go/metadata/main"),
            allow_files = True,
            cfg = "host",
            executable = True,
        ),
        "metadata": attr.label(
            doc = "A web_test metadata file.",
            allow_single_file = [".json"],
        ),
    },
    doc = "A configuration that can be used across multiple web_tests.",
    outputs = {"web_test_metadata": "%{name}.gen.json"},
    implementation = _web_test_config_impl,
)
