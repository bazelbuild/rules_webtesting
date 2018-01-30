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
"""Implementation of the web_test bazel rule

DO NOT load this file. Use "@io_bazel_rules_web//web:web.bzl".
"""

load("//web/internal:collections.bzl", "maps")
load("//web/internal:files.bzl", "files")
load("//web/internal:metadata.bzl", "metadata")
load("//web/internal:provider.bzl", "WebTestInfo")


def _web_test_impl(ctx):
  if ctx.attr.browser[WebTestInfo].disabled:
    return _generate_noop_test(
        ctx, """The browser configuration you requested is temporarily disabled.

Disabled browser: %s

Why was this browser disabled?
%s""" % (ctx.attr.browser.label, ctx.attr.browser[WebTestInfo].disabled))

  data_labels = [data.label for data in ctx.attr.data]

  if ctx.attr.test.label not in data_labels:
    fail("Test %s must be in data." % ctx.attr.test.label, "data")

  if ctx.attr.browser.label not in data_labels:
    fail("Browser %s must be in data." % ctx.attr.browser.label, "data")

  if ctx.attr.config.label not in data_labels:
    fail("Browser %s must be in data." % ctx.attr.browser.label, "data")

  missing_tags = [
      tag for tag in ctx.attr.browser[WebTestInfo].required_tags
      if (tag not in ctx.attr.tags) and (tag != "local" or not ctx.attr.local)
  ]

  if missing_tags:
    fail("Browser {browser} requires tags {tags} that are missing.".format(
        browser=ctx.attr.browser.label, tags=missing_tags))

  return _generate_default_test(ctx)


def _generate_noop_test(ctx, reason, status=0):
  """Generates a no-op test.

  Args:
    ctx: the ctx object for this rule.
    reason: string, a description of why the no-op test is being used.
    status: int, the exit code the test should return
  Returns:
    an empty struct for this rule.
  """
  if status:
    success = "fails"
  else:
    success = "passes"

  metadata.create_file(ctx, output=ctx.outputs.web_test_metadata)

  ctx.template_action(
      template=ctx.file.noop_web_test_template,
      output=ctx.outputs.executable,
      substitutions={
          "%TEMPLATED_success%": success,
          "%TEMPLATED_reason%": reason,
          "%TEMPLATED_status%": str(status),
      },
      executable=True)

  return []


def _generate_default_test(ctx):
  patch = ctx.new_file("%s.tmp.json" % ctx.label.name)
  metadata.create_file(
      ctx=ctx,
      output=patch,
      config_label=ctx.attr.config.label,
      label=ctx.label,
      test_label=ctx.attr.test.label)

  metadata.merge_files(
      ctx=ctx,
      merger=ctx.executable.merger,
      output=ctx.outputs.web_test_metadata,
      inputs=[
          patch,
          ctx.attr.config[WebTestInfo].metadata,
          ctx.attr.browser[WebTestInfo].metadata,
      ])

  env_vars = ""
  env = maps.clone(ctx.attr.browser[WebTestInfo].environment)
  env["WEB_TEST_METADATA"] = files.long_path(ctx, ctx.outputs.web_test_metadata)
  for k, v in env.items():
    env_vars += "export %s=%s\n" % (k, v)

  ctx.template_action(
      template=ctx.file.web_test_template,
      output=ctx.outputs.executable,
      substitutions={
          "%TEMPLATED_env_vars%":
              env_vars,
          "%TEMPLATED_launcher%":
              files.long_path(ctx, ctx.executable.launcher),
          "%TEMPLATED_metadata%":
              files.long_path(ctx, ctx.outputs.web_test_metadata),
          "%TEMPLATED_test%":
              files.long_path(ctx, ctx.executable.test),
      },
      executable=True)

  return [
      DefaultInfo(
          runfiles=ctx.runfiles(
              collect_data=True,
              collect_default=True,
              files=[ctx.outputs.web_test_metadata]))
  ]


web_test = rule(
    attrs={
        "browser":
            attr.label(mandatory=True, providers=[WebTestInfo]),
        "config":
            attr.label(mandatory=True, providers=[WebTestInfo]),
        "data":
            attr.label_list(allow_files=True, cfg="data"),
        "launcher":
            attr.label(cfg="target", executable=True),
        "merger":
            attr.label(
                cfg="host",
                executable=True,
                default=Label("//go/metadata/main")),
        "noop_web_test_template":
            attr.label(
                allow_single_file=True,
                default=Label("//web/internal:noop_web_test.sh.template")),
        "test":
            attr.label(cfg="target", executable=True, mandatory=True),
        "web_test_template":
            attr.label(
                allow_single_file=True,
                default=Label("//web/internal:web_test.sh.template")),
    },
    outputs={
        "web_test_metadata": "%{name}.gen.json",
    },
    test=True,
    implementation=_web_test_impl)
"""Runs a provided test against a provided browser configuration.

Args:
  browser: A browser configuration that defines the type of browser used for
    this test.
  config: Additional configuration that overrides the configuration in browser.
  data: Additional runtime dependencies for the test.
  launcher: The web test launcher binary.
  merger: The metadata merger binary.
  noop_web_test_template: Shell template used to launch test when browser is
    disabled.
  test: The test that will be run against the provided browser.
  web_test_template: Shell template used to launch test when browser is not
    disabled.
"""
