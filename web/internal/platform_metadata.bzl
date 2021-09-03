# Copyright 2021 Google LLC.
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

load("@io_bazel_rules_webtesting//web/internal:metadata.bzl", "metadata")
load("@io_bazel_rules_webtesting//web/internal:provider.bzl", "WebTestInfo")

def _label_to_manifest_path(label):
    """Converts the specified label to a manifest path"""
    if label.package != "":
        return "%s/%s" % (label.workspace_name, label.package)
    return label.workspace_name

def _platform_metadata_impl(ctx):
    """Implementation of the `platform_metadata` rule."""
    named_files = {}
    base_dir = _label_to_manifest_path(ctx.label)

    # Update the named files to manifest paths that can be resolved
    # with Bazel runfile resolution in web tests.
    for n, p in ctx.attr.named_files.items():
        named_files[n] = base_dir + "/" + p

    # Create a web test metadata file that will be provided as part of
    # the `WebTestInfo` provider.
    metadata.create_file(
        ctx = ctx,
        output = ctx.outputs.web_test_metadata,
        web_test_files = [
            metadata.web_test_files(ctx = ctx, named_files = named_files),
        ],
    )

    return [
        DefaultInfo(runfiles = ctx.runfiles(files = ctx.files.files)),
        WebTestInfo(metadata = ctx.outputs.web_test_metadata),
    ]

"""
  Rule that is used in combination with the `platform_archive` rule. It captures a set
  of files which are needed for dealing with a browser or tool. Additionally, specific files
  for the browser/tool can be denoted with an unique name so that web tests can access the
  files in a platform-agnostic way, regardless of which platform repository has been selected.

  The specified files are exposed as runfiles of the target defined through this rule. The unique
  names with their associated files are captured within a metadata file that is exposed through a
  `WebTestInfo` provider. Web tests will be able to deal with this metadata file to resolve
  platform-specific files in a platform-agnostic way.

  More details on this can be found in the `platform_archive` repository rule.
"""
platform_metadata = rule(
    attrs = {
        "files": attr.label_list(
            mandatory = True,
            allow_files = True,
            doc = "List of files which are needed for the tool.",
        ),
        "named_files": attr.string_dict(
            doc = """
              Dictionary that maps files to unique identifiers. This is useful
              if browser or tool archives are different on different platforms and
              the web tests would not want to care about archive-specific paths.
              e.g. targets expect a `CHROMIUM` key to point to the Chromium browser binary.
            """,
            mandatory = True,
        ),
    },
    outputs = {"web_test_metadata": "%{name}.gen.json"},
    implementation = _platform_metadata_impl,
)
