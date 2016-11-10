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

load("//web/internal:collections.bzl", "lists")
load("//web/internal:files.bzl", "files")
load("//web/internal:metadata.bzl", "metadata")


def _web_test_impl(ctx):
  if ctx.attr.browser.disabled:
    return _generate_noop_test(
        ctx, """the browser configuration you requested is temporarily disabled.

Disabled browser: %s

Why was this browser disabled?
%s""" % (ctx.attr.browser.label, ctx.attr.browser.disabled))

  missing_tags = [
      tag for tag in ctx.attr.browser.required_tags
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

  metadata.create_file(
      ctx=ctx,
      output=ctx.outputs.web_test_metadata,
      test_label=ctx.attr.test.label)

  ctx.template_action(
      template=ctx.file.noop_web_test_template,
      output=ctx.outputs.executable,
      substitutions={
          "%TEMPLATED_success%": success,
          "%TEMPLATED_reason%": reason,
          "%TEMPLATED_status%": str(status),
      },
      executable=True)

  return struct()


def _generate_default_test(ctx):
    """Generates the web_test.

  Args:
    ctx: the ctx object for this rule.
  Returns:
    the struct for this rule.
  """
  metadata_files = metadata.get_files(ctx,
                                      ["data", "test", "browser", "config"])

  patch = ctx.new_file("%s.tmp.json" % ctx.label.name)
  metadata.create_file(ctx=ctx, output=patch, test_label=ctx.attr.test.label)
  lists.ensure_at_end_of_list(metadata_files, patch)

  metadata.merge_files(
      ctx=ctx,
      merger=ctx.executable.merger,
      output=ctx.outputs.web_test_metadata,
      inputs=metadata_files)

  env_vars = ""
  for k, v in ctx.attr.browser.environment.items():
    env_vars += "export %s=%s\n" % (k, v)

  ctx.template_action(
      template=ctx.file.web_test_template,
      output=ctx.outputs.executable,
      substitutions={
          "%TEMPLATED_env_vars%":
              env_vars,
          "%TEMPLATED_launcher%":
              files.runfiles_path(ctx, ctx.executable.launcher),
          "%TEMPLATED_metadata%":
              files.runfiles_path(ctx, ctx.outputs.web_test_metadata),
          "%TEMPLATED_test%":
              files.runfiles_path(ctx, ctx.executable.test),
      },
      executable=True)

  return struct(runfiles=files.runfiles(
      ctx=ctx,
      files=[ctx.outputs.web_test_metadata],
      deps_attrs=["browser", "config", "launcher", "test"]))


web_test = rule(
    attrs={
        "test":
            attr.label(
                cfg="data",
                executable=True,
                mandatory=True,
                aspects=[metadata.aspect]),
        "browser":
            attr.label(
                cfg="data",
                mandatory=True,
                providers=[
                    "disabled",
                    "environment",
                    "required_tags",
                ],
                aspects=[metadata.aspect]),
        "config":
            attr.label(
                cfg="data",
                default=Label("//web:default_config"),
                aspects=[metadata.aspect]),
        "data":
            attr.label_list(
                allow_files=True, cfg="data", aspects=[metadata.aspect]),
        "merger":
            attr.label(
                cfg="host",
                executable=True,
                default=Label("//go/metadata:merger")),
        "launcher":
            attr.label(
                cfg="data",
                executable=True,
                default=Label("//go/launcher:main")),
        "web_test_template":
            attr.label(
                allow_single_file=True,
                default=Label("//web/internal:web_test.sh.template")),
        "noop_web_test_template":
            attr.label(
                allow_single_file=True,
                default=Label("//web/internal:noop_web_test.sh.template")),
    },
    outputs={"web_test_metadata": "%{name}.gen.json",},
    test=True,
    implementation=_web_test_impl)
"""Runs a provided test against a provided browser configuration.

Args:
  test: The test that will be run against the provided browser.
  browser: A browser configuration that defines the type of browser used for
    this test.
  config: Additional configuration that overrides the configuration in browser.
  data: Additional runtime dependencies for the test.
  merger: A web test metadata merger binary.
  launcher: A web test launcher binary.
  web_test_template: The shell script that should be executed for running the
    web_test.
  noop_web_test_template: The shell script that should be executed if the
    browser is disabled.
"""
