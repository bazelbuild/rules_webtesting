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

load("//web/internal:metadata.bzl", "metadata")


def _browser_impl(ctx):
  """Implementation of the browser rule."""
  patch = ctx.new_file("%s.tmp.json" % ctx.label.name)
  metadata.create_file(ctx=ctx, output=patch, browser_label=ctx.label)
  metadata_files = [
      patch,
      ctx.file.metadata,
  ] + [dep.web_test.metadata for dep in ctx.attr.deps]

  metadata.merge_files(
      ctx=ctx,
      merger=ctx.executable.merger,
      output=ctx.outputs.web_test_metadata,
      inputs=metadata_files)

  return struct(
      runfiles=ctx.runfiles(collect_data=True, collect_default=True),
      web_test=struct(
          disabled=ctx.attr.disabled,
          environment=ctx.attr.environment,
          metadata=ctx.outputs.web_test_metadata,
          required_tags=ctx.attr.required_tags))


browser = rule(
    attrs={
        "data":
            attr.label_list(allow_files=True, cfg="data"),
        "deps":
            attr.label_list(providers=["web_test"]),
        "disabled":
            attr.string(),
        "environment":
            attr.string_dict(),
        "merger":
            attr.label(
                executable=True,
                cfg="host",
                default=Label("//go/metadata/main")),
        "metadata":
            attr.label(mandatory=True, allow_single_file=[".json"], cfg="data"),
        "required_tags":
            attr.string_list(),
    },
    outputs={"web_test_metadata": "%{name}.gen.json"},
    implementation=_browser_impl)
"""Defines a browser configuration for use with web_test.

Args:
  data: Runtime dependencies for this configuration.
  deps: Other web_test-related rules that this rule depends on.
  disabled: If set, then a no-op test will be run for all tests using
    this browser.
  environment: Map of environment variables-values to set for this browser.  
  merger: Metadata merger executable.  
  metadata: The web_test metadata file that defines how this browser
    is launched and default capabilities for this browser.  
  required_tags: A list of tags that all web_tests using this browser
    should have. Examples include "requires-network", "local", etc.
"""
