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
"""A rule for defining files that can be located by name.

DO NOT load this file. Use "@io_bazel_rules_web//web:web.bzl".
"""

load(":metadata.bzl", "metadata")
load(":provider.bzl", "WebTestInfo")

def _web_test_named_file_impl(ctx):
    name = ctx.attr.alt_name or ctx.label.name

    metadata.create_file(
        ctx = ctx,
        output = ctx.outputs.web_test_metadata,
        web_test_files = [
            metadata.web_test_files(ctx = ctx, named_files = {name: ctx.file.file}),
        ],
    )

    return [
        DefaultInfo(
            runfiles = ctx.runfiles(
                collect_data = True,
                collect_default = True,
                files = [ctx.file.file],
            ),
        ),
        WebTestInfo(metadata = ctx.outputs.web_test_metadata),
    ]

web_test_named_file = rule(
    attrs = {
        "alt_name": attr.string(doc = "If supplied, is used instead of name."),
        "file": attr.label(
            doc = "The file that will be returned for name.",
            allow_single_file = True,
            mandatory = True,
        ),
    },
    doc = "Defines a file that can be located by name.",
    outputs = {"web_test_metadata": "%{name}.gen.json"},
    implementation = _web_test_named_file_impl,
)
