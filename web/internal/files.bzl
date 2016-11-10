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
"""Utility functions for working with runfiles."""

load('//web/internal:collections.bzl', 'lists')


def _runfiles(ctx, deps_attrs=[], files=[]):
  """Creates a runfiles object.

  Args:
    ctx: the ctx object to create the runfiles object from.
    deps_attrs: sequence of strings; the attributes on ctx to
      transitively collect runfiles from. Note: runfiles from
      srcs, data, and deps attributes are always included.
    files: sequence of Files; list of files to be added to runfiles.
  Returns:
    A configured runfiles object.
  """
  transitive_files = set()
  for attr in deps_attrs:
    transitive_files = transitive_files | _get_transitive_files(
        getattr(ctx.attr, attr))
    if hasattr(ctx.files, attr):
      transitive_files = transitive_files | getattr(ctx.files, attr)
  return ctx.runfiles(
      files=files,
      transitive_files=transitive_files,
      collect_data=True,
      collect_default=True,)


def _get_transitive_files(attr_val):
  """Collects the set of transitive runfiles for the given attribute value.

  If the attribute is a sequence type, then it recurses into the values of
  sequence.

  Args:
    attr_value: value fo the attribute to collect runfiles from.
  Returns:
    set of Files.
  """
  transitive_files = set()
  if lists.is_list_like(attr_val):
    for val in attr_val:
      transitive_files = transitive_files | _get_transitive_files(val)
    return transitive_files
  if hasattr(attr_val, 'data_runfiles'):
    transitive_files = transitive_files | attr_val.data_runfiles.files
  if hasattr(attr_val, 'default_runfiles'):
    transitive_files = transitive_files | attr_val.default_runfiles.files
  return transitive_files


def _runfiles_path(ctx, file):
  """Constructs a path relative to TEST_SRCDIR for accessing the file."""
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


files = struct(
    runfiles=_runfiles,
    runfiles_path=_runfiles_path,)
