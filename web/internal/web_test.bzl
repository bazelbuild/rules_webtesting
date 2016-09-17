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

load("//web/internal:shared.bzl", "build_runfiles", "get_metadata_files",
     "merge_metadata_files", "path")


def _web_test_impl(ctx):
  missing_tags = [
      tag for tag in ctx.attr.browser.required_tags
      if (tag not in ctx.attr.tags) and (tag != "local" or not ctx.attr.local)
  ]

  if missing_tags:
    fail("Browser {browser} requires tags {tags} that are missing.".format(
        browser=ctx.attr.browser.label, tags=missing_tags))

  metadata_files = get_metadata_files(ctx, ["data", "browser", "config"])

  merge_metadata_files(
      ctx=ctx,
      merger=ctx.executable._merger,
      output=ctx.outputs.web_test_metadata,
      inputs=metadata_files,)

  if ctx.attr.browser.disabled:
    return _generate_disabled_test(ctx)

  return _generate_default_test(ctx)


def _generate_disabled_test(ctx):
  ctx.file_action(
      content="""#!/bin/bash

cat << EOF
#####################################################################
This test always passes. Your test was not run.

This dummy test was inserted in place of the web test you intended
to run because the browser configuration you requested has been
temporarily disabled.

Disabled browser: {browser}

Why was this browser disabled?
{message}
#####################################################################
EOF

exit 0
""".format(
          browser=ctx.attr.browser.label, message=ctx.attr.disabled),
      output=ctx.outputs.executable,
      executable=True,)

  return struct()


def _generate_default_test(ctx):
  env_vars = ""
  for k, v in ctx.attr.browser.environment.items():
    env_vars += "export %s=%s\n" % (k, v)

  ctx.file_action(
      content="""#!/bin/bash

if [[ -z "$TEST_SRCDIR" ]]; then
  case "$0" in
    /*) self="$0" ;;
    *)  self="$PWD/$0" ;;
  esac

  if [[ -e "$self.runfiles" ]]; then
    export TEST_SRCDIR="$self.runfiles"
  else
    echo "Unable to determine runfiles location"
    exit -1
  fi
fi

if [[ -z "$TEST_TEMPDIR" ]]; then
  export TEST_TEMPDIR=$(mktemp -d test_tempdir.XXXXXX)
fi

if [ ! -e "/dev/shm" ]; then
  mkdir /dev/shm
fi

{env_vars}

printenv

$TEST_SRCDIR/{launcher} --metadata={metadata} --test={test}
""".format(
          env_vars=env_vars,
          launcher=path(ctx, ctx.executable._launcher),
          metadata=path(ctx, ctx.outputs.web_test_metadata),
          test=path(ctx, ctx.executable.test)),
      output=ctx.outputs.executable,
      executable=True)
  return struct(runfiles=build_runfiles(
      ctx,
      files=[ctx.outputs.web_test_metadata],
      deps_attrs=["_launcher", "browser", "config", "test"]))


web_test = rule(
    implementation=_web_test_impl,
    test=True,
    attrs={
        "test": attr.label(
            executable=True, mandatory=True, cfg="data"),
        "browser": attr.label(
            mandatory=True,
            cfg="data",
            providers=[
                "disabled",
                "environment",
                "required_tags",
                "web_test_metadata",
            ],),
        "config": attr.label(
            cfg="data",
            default=Label("//external:web_test_default_config"),
            providers=[
                "web_test_metadata",
            ],),
        "data": attr.label_list(
            allow_files=True, cfg="data"),
        "_merger": attr.label(
            executable=True,
            cfg="host",
            default=Label("//external:web_test_merger")),
        "_launcher": attr.label(
            executable=True, default=Label("//external:web_test_launcher")),
    },
    outputs={"web_test_metadata": "%{name}.gen.json"},)
"""Runs a provided test against a provided browser configuration.

Args:
  test: The test that will be run against the provided browser.
  browser: A browser configuration that defines the type of browser used for
    this test.
  config: Additional configuration that overrides the configuration in browser.
  data: Additional runtime dependencies for the test.
"""
