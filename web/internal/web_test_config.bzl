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

load("//web/internal:files.bzl", "files")
load("//web/internal:metadata.bzl", "metadata")


def _web_test_config_impl(ctx):
  """Implementation of the web_test_config rule."""
  if ctx.file.metadata:
    m = [ctx.file.metadata]
  else:
    m = []
  return struct(runfiles=files.runfiles(ctx=ctx), web_test_metadata=m)


web_test_config = rule(
    attrs={
        "configs":
            attr.label_list(providers=["web_test_metadata"]),
        "metadata":
            attr.label(allow_single_file=True),
        "data":
            attr.label_list(
                allow_files=True, cfg="data", aspects=[metadata.aspect]),
    },
    implementation=_web_test_config_impl)
"""A browser-independent configuration that can be used across multiple web_tests.

Args:
  configs: A list of web_test_config rules that this rule inherits from.
    Configuration in rules later in the list will override configuration
    earlier in the list.
  metadata: A web_test metadata file with configuration that will override
    all other configuration.
  data: Additional files that this web_test_config depends on at runtime.
"""
