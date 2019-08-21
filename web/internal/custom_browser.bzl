# Copyright 2018 Google Inc.
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
"""The custom_browser rule is used to extend existing web_test browsers.

DO NOT load this file. Use "@io_bazel_rules_web//web:web.bzl".
"""

load(":metadata.bzl", "metadata")
load(":provider.bzl", "WebTestInfo")
load(":runfiles.bzl", "runfiles")
load("@bazel_skylib//lib:dicts.bzl", "dicts")

def _custom_browser_impl(ctx):
    """Implementation of the custom_browser rule."""
    metadata_files = [
        dep[WebTestInfo].metadata
        for dep in reversed(ctx.attr.deps)
    ] + [ctx.attr.browser[WebTestInfo].metadata]

    if ctx.files.metadata:
        metadata_files.append(ctx.file.metadata)

    patch = ctx.actions.declare_file("%s.tmp.json" % ctx.label.name)
    metadata.create_file(ctx = ctx, output = patch, browser_label = ctx.label)
    metadata_files.append(patch)

    metadata.merge_files(
        ctx = ctx,
        merger = ctx.executable.merger,
        output = ctx.outputs.web_test_metadata,
        inputs = metadata_files,
    )

    env = dicts.add(
        ctx.attr.browser[WebTestInfo].environment,
        ctx.attr.environment,
    )

    return [
        DefaultInfo(runfiles = runfiles.collect(ctx, targets = [ctx.attr.browser])),
        WebTestInfo(
            metadata = ctx.outputs.web_test_metadata,
            required_tags = ctx.attr.browser[WebTestInfo].required_tags,
            disabled = ctx.attr.disabled or ctx.attr.browser[WebTestInfo].disabled,
            environment = env,
            execution_requirements = ctx.attr.browser[WebTestInfo].execution_requirements,
        ),
    ]

custom_browser = rule(
    implementation = _custom_browser_impl,
    attrs = {
        "browser": attr.label(
            doc = "The browser configuration this browser extends.",
            mandatory = True,
            providers = [WebTestInfo],
        ),
        "deps": attr.label_list(
            doc = "Other web_test-related rules that this rule depends on.",
            providers = [WebTestInfo],
        ),
        "data": attr.label_list(
            doc = "Runtime dependencies for this configuration.",
            allow_files = True,
        ),
        "disabled": attr.string(
            doc =
                "If set then a no-op test will be run when using this browser.",
        ),
        "metadata": attr.label(
            doc = "The web_test metadata file that defines how this browser is " +
                  "launched and the default capabilities for this browser.",
            allow_single_file = True,
        ),
        "merger": attr.label(
            doc = "The metadata merger binary.",
            default = Label("//go/metadata/main"),
            allow_files = True,
            cfg = "host",
            executable = True,
        ),
        "environment": attr.string_dict(doc = "Map of environment variables-values to set."),
    },
    outputs = {"web_test_metadata": "%{name}.gen.json"},
)
