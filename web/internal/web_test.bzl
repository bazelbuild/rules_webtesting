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

load(":files.bzl", "files")
load(":metadata.bzl", "metadata")
load(":provider.bzl", "WebTestInfo")
load(":runfiles.bzl", "runfiles")

def _web_test_impl(ctx):
    if ctx.attr.browser[WebTestInfo].disabled:
        return _generate_noop_test(
            ctx,
            """The browser configuration you requested is temporarily disabled.

Disabled browser: %s

Why was this browser disabled?
%s""" % (ctx.attr.browser.label, ctx.attr.browser[WebTestInfo].disabled),
        )

    missing_tags = [
        tag
        for tag in ctx.attr.browser[WebTestInfo].required_tags
        if (tag not in ctx.attr.tags) and (tag != "local" or not ctx.attr.local)
    ]

    if missing_tags:
        fail("Browser {browser} requires tags {tags} that are missing.".format(
            browser = ctx.attr.browser.label,
            tags = missing_tags,
        ))

    return _generate_default_test(ctx)

def _generate_noop_test(ctx, reason, status = 0):
    """Generates a no-op test.

    Args:
        ctx: the ctx object for this rule.
        reason: string, a description of why the no-op test is being used.
        status: int, the exit code the test should return.

    Returns:
        an empty struct for this rule.
    """
    if status:
        success = "fails"
    else:
        success = "passes"

    metadata.create_file(ctx, output = ctx.outputs.web_test_metadata)

    test = ctx.actions.declare_file(ctx.label.name)
    ctx.actions.expand_template(
        template = ctx.file.noop_web_test_template,
        output = test,
        substitutions = {
            "%TEMPLATED_success%": success,
            "%TEMPLATED_reason%": reason,
            "%TEMPLATED_status%": str(status),
        },
        is_executable = True,
    )

    return [DefaultInfo(executable = test)]

def _generate_default_test(ctx):
    patch = ctx.actions.declare_file("%s.tmp.json" % ctx.label.name)
    metadata.create_file(
        ctx = ctx,
        output = patch,
        config_label = ctx.attr.config.label,
        label = ctx.label,
        test_label = ctx.attr.test.label,
    )

    metadata.merge_files(
        ctx = ctx,
        merger = ctx.executable.merger,
        output = ctx.outputs.web_test_metadata,
        inputs = [
            patch,
            ctx.attr.config[WebTestInfo].metadata,
            ctx.attr.browser[WebTestInfo].metadata,
        ],
    )

    env_vars = ""
    env = {}
    env.update(ctx.attr.browser[WebTestInfo].environment)
    env["WEB_TEST_METADATA"] = files.long_path(ctx, ctx.outputs.web_test_metadata)
    for k, v in env.items():
        env_vars += "export %s=%s\n" % (k, v)

    test = ctx.actions.declare_file(ctx.label.name)
    ctx.actions.expand_template(
        template = ctx.file.web_test_template,
        output = test,
        substitutions = {
            "%TEMPLATED_env_vars%": env_vars,
            "%TEMPLATED_launcher%": files.long_path(ctx, ctx.executable.launcher),
            "%TEMPLATED_metadata%": files.long_path(ctx, ctx.outputs.web_test_metadata),
            "%TEMPLATED_test%": files.long_path(ctx, ctx.executable.test),
        },
        is_executable = True,
    )

    return [
        DefaultInfo(
            executable = test,
            runfiles = runfiles.collect(
                ctx = ctx,
                files = [ctx.outputs.web_test_metadata],
                targets = [
                    ctx.attr.browser,
                    ctx.attr.config,
                    ctx.attr.launcher,
                    ctx.attr.test,
                ],
            ),
        ),
        testing.ExecutionInfo(env),
        testing.TestEnvironment(ctx.attr.browser[WebTestInfo].environment),
    ]

web_test = rule(
    attrs = {
        "browser": attr.label(
            doc = "The browser configuration to use for this test.",
            mandatory = True,
            providers = [WebTestInfo],
        ),
        "config": attr.label(
            doc = "Additional configuration for this test.",
            mandatory = True,
            providers = [WebTestInfo],
        ),
        "data": attr.label_list(
            doc = "Additional runtime dependencies for the test.",
            allow_files = True,
        ),
        "launcher": attr.label(
            doc = "The web test launcher binary.",
            allow_files = True,
            cfg = "target",
            executable = True,
        ),
        "merger": attr.label(
            doc = "The metadata merger binary.",
            default = Label("//go/metadata/main"),
            allow_files = True,
            cfg = "host",
            executable = True,
        ),
        "noop_web_test_template": attr.label(
            doc =
                "Shell template used to launch test when browser is disabled.",
            default = Label("//web/internal:noop_web_test.sh.template"),
            allow_single_file = True,
        ),
        "test": attr.label(
            doc = "The test that will be run against the provided browser.",
            cfg = "target",
            executable = True,
            mandatory = True,
        ),
        "web_test_template": attr.label(
            doc = "Shell template used to launch test.",
            default = Label("//web/internal:web_test.sh.template"),
            allow_single_file = True,
        ),
    },
    doc = "Runs a provided test against a provided browser configuration.",
    outputs = {
        "web_test_metadata": "%{name}.gen.json",
    },
    test = True,
    implementation = _web_test_impl,
)
