# Copyright 2016 Google Inc. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
"""Defines a go_repository rule that downloads and sets up a go package."""


def go_repository(name,
                  urls,
                  sha256,
                  license,
                  import_name,
                  strip_prefix=None,
                  excluded_srcs=None,
                  license_file='LICENSE',
                  deps=None,
                  testonly=None):
  """go_repository downloads a go package and makes it available for import.

  The bazel dep for the go package is "@name".
  The go import for the go package is "import_name".

  Args:
      name: Name of the repository rule. BUILD files can depend on the package
        as "@name"
      urls: A list of URLs where the package can be downloaded from.
      sha256: SHA256 for downloaded file.
      license: A string indicating the licenses for the package.
        e.g. "licenses([\"notice\"])  # MIT.".
      import_name: The value for go_prefix. Go files can import this package
        as import "import_name".
      strip_prefix: Any directories to strip of the prefix when extracting the
        downloaded file.
      excluded_srcs: Any .go file in the package that should be excluded.
        **/*test.go is always excluded.
      license_file: The name of the license file. By default "LICENSE".
      deps: Any dependencies for the go package.
      testonly: Whether the rules should be marked testonly.
  """
  if deps:
    deps_str = '"' + '",\n        "'.join(deps) + '"'
  else:
    deps_str = ''

  if excluded_srcs:
    excludes_str = '"' + '",\n            "'.join(excluded_srcs) + '",'
  else:
    excludes_str = ''

  build_file = """
load("@io_bazel_rules_go//go:def.bzl", "go_prefix", "go_library")

{license}

exports_files(["{license_file}"])

go_prefix("{import_name}")

go_library(
    name = "go_default_library",
    srcs = glob(
        ["*.go"],
        exclude = [
            "*test.go",
            {excludes}
        ],
    ),
    testonly = {testonly},
    deps = [
        {deps}
    ],
)

alias(
    name = "{name}",
    actual = ":go_default_library",
    testonly = {testonly},
    visibility = ["//visibility:public"],
)
""".format(
      name=name,
      license=license,
      license_file=license_file,
      import_name=import_name,
      testonly=testonly,
      deps=deps_str,
      excludes=excludes_str)

  native.new_http_archive(
      name=name,
      build_file_content=build_file,
      sha256=sha256,
      urls=urls,
      strip_prefix=strip_prefix)
