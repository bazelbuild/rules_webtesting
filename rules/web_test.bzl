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


def _web_test_impl(ctx):
  """TODO: implement"""
  ctx.file_action(
      content="""#!/bin/bash

 printenv

 $TEST_SRCDIR/$TEST_WORKSPACE/%s --metadata=$TEST_WORKSPACE/%s --test=$TEST_WORKSPACE/%s
 """ % (ctx.executable._launcher.short_path, ctx.attr.browser.json.short_path,
        ctx.executable.test.short_path),
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
