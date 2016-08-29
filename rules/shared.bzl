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

def browser_attr(mandatory):
  """browser_attr returns a configured browser attribute.

  Args:
    mandatory: bool; whether the browser attribute is required or not.
  Returns:
    an attr.label for a browser attribute.
  """
  return attr.label(
      mandatory=mandatory,
      cfg=DATA_CFG,
      providers=[
          "disabled",
          "environment",
          "required_tags",
          "web_test_metadata",
      ],)


def browser_struct(runfiles,
                   web_test_metadata,
                   disabled="",
                   environment=None,
                   required_tags=None,):
  """Creates a struct for the return value of browser rule.

  Args:
    runfiles: a runfiles object.
    web_test_metadata: File; the web test metadata file for this
      browser.
    enabled: bool; whether this browser is enabled or not.
    environment; dict; a map of environment variable names to values that
      should be set before starting Web Test Launcher.
    required_tags: sequence of Strings; tags that must be applied to
      any web_test using this browser.
  Returns:
    a struct
  """
  environment = environment or {}
  required_tags = required_tags or []
  return struct(
      disabled=disabled,
      environment=environment,
      required_tags=required_tags,
      runfiles=runfiles,
      web_test_metadata=web_test_metadata,)


def config_attr():
  """config_attr returns a configured config attribute.

  Returns:
    an attr.label for a config attribute.
  """
  return attr.label(
      cfg=DATA_CFG,
      default=Label("//external:web_test_default_config"),
      providers=[
          "web_test_metadata",
      ],)


def config_struct(runfiles,
                  web_test_metadata,):
  """Creates a struct for the return value of web_test_config rule.

  Args:
    runfiles: a runfiles object.
    web_test_metadata: File; the web test metadata file for this
      browser.
  Returns:
    a struct
  """
  return struct(
      runfiles=runfiles,
      web_test_metadata=web_test_metadata,)


def build_runfiles(ctx, deps_attrs=[], files=[]):
  """Creates a runfiles object.

  Args:
    ctx: the ctx object to create the runfiles object from.
    dep_attrs: sequence of strings; the attributes on ctx to
      transitively collect runfiles from. Note: runfiles from
      srcs, data, and deps attributes are always included.
    files: sequence of Files; list of files to be added to runfiles.
  Returns:
    A configured runfiles object.
  """
  transitive_files = set()
  for attr in deps_attrs:
    transitive_files += _get_transitive_files(getattr(ctx.attr, attr))
    if hasattr(ctx.files, attr):
      transitive_files += getattr(ctx.files, attr)
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
  if type(attr_val) == type([]) or type(attr_val) == type(set()):
    for val in attr_val:
      transitive_files += _get_transitive_files(val)
    return transitive_files
  if hasattr(attr_val, "data_runfiles"):
    transitive_files += attr_val.data_runfiles.files
  if hasattr(attr_val, "default_runfiles"):
    transitive_files += attr_val.default_runfiles.files
  return transitive_files


def path(ctx, file):
  if file.owner and file.owner.workspace_root:
    return file.owner.workspace_root + "/" + file.short_path
  else:
    return ctx.workspace_name + "/" + file.short_path
