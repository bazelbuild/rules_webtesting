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

load(
    "//web/internal:shared.bzl",
    "build_runfiles",
    "create_metadata_file",
    "get_metadata_files",
    "merge_metadata_files",)


def _browser_impl(ctx):
  """Implementation of the browser rule."""
  patch = ctx.new_file("%s.tmp.json" % ctx.label.name)
  create_metadata_file(
      ctx=ctx,
      output=patch,
      browser_label=ctx.label,)

  metadata_files = get_metadata_files(ctx,
                                      ["data"]) + [ctx.file.metadata, patch]

  merge_metadata_files(
      ctx=ctx,
      merger=ctx.executable.merger,
      output=ctx.outputs.web_test_metadata,
      inputs=metadata_files,)

  return struct(
      disabled=ctx.attr.disabled,
      environment=ctx.attr.environment,
      required_tags=ctx.attr.required_tags,
      runfiles=build_runfiles(
          ctx, files=[ctx.outputs.web_test_metadata]),
      web_test_metadata=ctx.outputs.web_test_metadata,)


browser = rule(
    attrs={
        "metadata":
            attr.label(
                mandatory=True,
                allow_single_file=True,
                cfg="data",),
        "data":
            attr.label_list(
                allow_files=True,
                cfg="data",),
        "disabled":
            attr.string(),
        "environment":
            attr.string_dict(default={}),
        "required_tags":
            attr.string_list(default=[]),
        "merger":
            attr.label(
                executable=True,
                cfg="host",
                default=Label("//go/metadata:merger"),),
    },
    outputs={"web_test_metadata": "%{name}.gen.json"},
    implementation=_browser_impl,)
"""Defines a browser configuration for use with web_test.

Args:
  metadata: The web_test metadata file that defines how this browser
    is launched and default capabilities for this browser.
  data: Runtime dependencies needed for this browser.
  disabled: If set, then a no-op test will be run for all tests using
    this browser.
  environment: Map of environment variables-values to set for this browser.
  required_tags: A list of tags that all web_tests using this browser
    should have. Examples include "requires-network", "local", etc.
"""
