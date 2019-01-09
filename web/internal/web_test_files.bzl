# Copyright 2017 Google Inc.
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
"""A rule for configuring a set of named files.

DO NOT load this file. Use "@io_bazel_rules_web//web:web.bzl".
"""

load(":metadata.bzl", "metadata")
load(":provider.bzl", "WebTestInfo")

def _web_test_files_impl(ctx):
    named_files = {}
    runfiles = depset()

    for target, name in ctx.attr.files.items():
        if name in named_files:
            fail("%s appears multiple times." % name, "files")
        if len(target.files) != 1:
            fail("%s refers to multiple files." % target.label, "files")
        named_files[name] = target.files.to_list()[0]
        runfiles = depset(transitive = [target.files, runfiles])

    metadata.create_file(
        ctx = ctx,
        output = ctx.outputs.web_test_metadata,
        web_test_files = [
            metadata.web_test_files(ctx = ctx, named_files = named_files),
        ],
    )

    return [
        DefaultInfo(
            runfiles = ctx.runfiles(
                collect_data = True,
                collect_default = True,
                files = runfiles.to_list(),
            ),
        ),
        WebTestInfo(metadata = ctx.outputs.web_test_metadata),
    ]

web_test_files = rule(
    attrs = {
        "files": attr.label_keyed_string_dict(
            doc = "A map of files to names.",
            mandatory = True,
            allow_files = True,
            allow_empty = False,
        ),
    },
    doc = "Specifies a set of named files.",
    outputs = {"web_test_metadata": "%{name}.gen.json"},
    implementation = _web_test_files_impl,
)
