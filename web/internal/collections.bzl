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
"""Collections module contains util functions for working with collections.

Usage:
  load("//testing/web/build_defs:collections.bzl", "lists", "maps")

  l = lists.clone(l)
  lists.ensure_contains(l, "//some:target")
"""


def _is_list_like(val):
  """Checks is val is a list-like (list, set, tuple) value."""
  return type(val) in [type([]), type(set()), type(())]


def _list_ensure_contains(lst, item):
  """Appends the specified item to the list if its not already a member."""
  if item not in lst:
    lst.append(item)


def _list_ensure_contains_all(lst, items):
  """Appends the specified items to the list if not already members."""
  for item in items:
    _list_ensure_contains(lst, item)


def _list_clone(original):
  """Create a new list with content of original."""
  if _is_list_like(original):
    return list(original)
  if original:
    fail('got "' + original + '", but expected none or a list-like value')
  return []


lists = struct(
    clone=_list_clone,
    ensure_contains=_list_ensure_contains,
    ensure_contains_all=_list_ensure_contains_all,
    is_list_like=_is_list_like)


def _map_clone(original):
  new_map = {}
  if original:
    new_map.update(original)
  return new_map


maps = struct(clone=_map_clone)
