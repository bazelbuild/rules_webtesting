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


def _web_test_impl(ctx):
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

printenv

$TEST_SRCDIR/%s --metadata=%s --test=%s
""" % (_path(ctx, ctx.executable._launcher), _path(ctx, ctx.attr.browser.json),
       _path(ctx, ctx.executable.test)),
      output=ctx.outputs.executable,
      executable=True)
  return struct(runfiles=ctx.runfiles(
      files=ctx.files.test + ctx.files._launcher + ctx.files.browser,
      collect_default=True,
      collect_data=True))


web_test = rule(
    implementation=_web_test_impl,
    test=True,
    attrs={
        "test": attr.label(
            executable=True, mandatory=True, cfg=DATA_CFG),
        "browser": attr.label(
            mandatory=True, cfg=DATA_CFG),
        "config": attr.label(
            default=Label("//external:web_test_default_config"), providers=[]),
        "data": attr.label_list(
            allow_files=True, cfg=DATA_CFG),
        "_merger": attr.label(
            executable=True,
            cfg=HOST_CFG,
            default=Label("//external:web_test_merger")),
        "_launcher": attr.label(
            executable=True, default=Label("//external:web_test_launcher")),
    },)


def _path(ctx, file):
  if file.owner and file.owner.workspace_root:
    return file.owner.workspace_root + "/" + file.short_path
  else:
    return ctx.workspace_name + "/" + file.short_path
