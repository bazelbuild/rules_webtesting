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

load(":metadata.bzl", "metadata")
load(":provider.bzl", "WebTestInfo")
load(":runfiles.bzl", "runfiles")

def _browser_impl(ctx):
  """Implementation of the browser rule."""
  patch = ctx.new_file("%s.tmp.json" % ctx.label.name)
  metadata.create_file(ctx=ctx, output=patch, browser_label=ctx.label)
  metadata_files = [
      patch,
      ctx.file.metadata,
  ] + [dep[WebTestInfo].metadata for dep in ctx.attr.deps]

  metadata.merge_files(
      ctx=ctx,
      merger=ctx.executable.merger,
      output=ctx.outputs.web_test_metadata,
      inputs=metadata_files)

  return [
      DefaultInfo(runfiles=runfiles.collect(ctx)),
      WebTestInfo(
          disabled=ctx.attr.disabled,
          environment=ctx.attr.environment,
          execution_requirements=ctx.attr.execution_requirements,
          metadata=ctx.outputs.web_test_metadata,
          required_tags=ctx.attr.required_tags),
  ]

browser = rule(
    attrs = {
        "data": attr.label_list(
            doc = "Runtime dependencies for this configuration.",
            allow_files = True,
            cfg = "data",
        ),
        "deps": attr.label_list(
            doc = "Other web_test-related rules that this rule depends on.",
            providers = [WebTestInfo],
        ),
        "disabled": attr.string(
            doc =
                "If set then a no-op test will be run when using this browser.",
        ),
        "environment": attr.string_dict(doc = "Map of environment variables-values to set."),
        "execution_requirements": attr.string_dict(
            doc = "Map of execution requirements for this browser.",
        ),
        "merger": attr.label(
            doc = "Metadata merger executable.",
            default = Label("//go/metadata/main"),
            executable = True,
            cfg = "host",
        ),
        "metadata": attr.label(
            doc = "The web_test metadata file that defines how this browser is " +
                  "launched and default capabilities for this browser.",
            mandatory = True,
            allow_single_file = [".json"],
        ),
        "required_tags": attr.string_list(
            doc =
                "A list of tags that web_tests using this browser should have.",
        ),
    },
    doc = "Defines a browser configuration for use with web_test.",
    outputs = {"web_test_metadata": "%{name}.gen.json"},
    implementation = _browser_impl,
)
