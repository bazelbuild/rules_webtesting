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
"""A library of functions for working with runfiles."""

load('//web/internal:collections.bzl', 'lists')


def _long_path(ctx, file):
  """Constructs a path relative to TEST_SRCDIR for accessing the file.
  
  Args:
    ctx: a Skylark rule context.
    file: a File object. The file should appear in the runfiles for the
      test.
  Returns:
    A string path relative to TEST_SRCDIR suitable for use in tests and
    testing infrastructure.
  """
  if file.short_path[:3] == '../':
    # sometimes a file's short_path is ../<workspace_root>/<file_path>
    # then we just need to trim the ../
    return file.short_path[3:]
  if file.owner and file.owner.workspace_root:
    # if the file has an owner and that owner has a workspace_root,
    # prepend it.
    return file.owner.workspace_root + '/' + file.short_path
  # otherwise assume the file is in the same workspace as the current rule.
  return ctx.workspace_name + '/' + file.short_path


files = struct(long_path=_long_path)
