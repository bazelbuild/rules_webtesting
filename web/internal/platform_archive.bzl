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

def _platform_archive_impl(ctx):
    """Implementation of the `platform_archive` rule."""
    urls = ctx.attr.urls
    sha256 = ctx.attr.sha256

    # If the URL resolves to a `.dmg` file, then we need to convert the file
    # to a zip so that we can extract the actual binaries. We use the `convert_dmg`
    # script provided by the webtesting Bazel rules.
    if urls[0].endswith(".dmg"):
        download_file_name = "_download_file_%s.dmg" % ctx.attr.name
        result_zip_name = "_converted_file_%s.zip" % ctx.attr.name

        ctx.download(urls, download_file_name, sha256)
        ctx.execute([ctx.path(Label("@io_bazel_rules_webtesting//web/internal:convert_dmg.sh")), download_file_name, result_zip_name])
        ctx.extract(result_zip_name)

        ctx.delete(result_zip_name)
        ctx.delete(download_file_name)
    else:
        ctx.download_and_extract(
            url = urls,
            sha256 = sha256,
        )

    # The platform archive has been downloaded and extracted. We now generate a repository
    # `BUILD.bazel` file that exposes the archive files, together with the specified
    # named files using the `platform_metadata` rule.
    ctx.file("BUILD.bazel", content = """
load("@io_bazel_rules_webtesting//web:web.bzl", "platform_metadata")

licenses(%s)

platform_metadata(
  name = "metadata",
  files = glob(["**/*"]),
  named_files = %s,
  visibility = ["//visibility:public"],
)
""" % (str(ctx.attr.licenses), str(ctx.attr.named_files)))

"""
  Rule that can be used to download and unpack a platform-specific archive (e.g. for a browser)
  a dedicated Bazel repository. Additionally, files within the archive can be denoted with an
  unique name so that web tests can access these files in a platform-agnostic way, regardless of
  which `platform_archive` has been selected as a dependency (i.e. through constraints and select)

  As an example for the concept of denoting archive files with an unique name, consider a case
  where a a web test decides conditionally based on the current exec platform which
  `platform_archive` repository is used (e.g. mac, windows or linux). The archives are different
  for each platform. The test usually would need to determine the current platform, and know how
  each archive is structured in order to access the browser binary within the repository. By
  defining named files though, the web test could just pull a named file called `BINARY` that
  always resolves to the browser binary in a platform-agnostic way.

  Note #1: This rule exists as an alternative to the previous `platform_http_file` concept
  because this rule does not extract the archive directly, but relies on later build actions
  to perform the unpacking. The old `platform_http_file` approach resulted in less efficient
  caching because build actions are likely to be invalidated more frequently (e.g. `bazel clean).
  We also noticed that the extraction within RBE containers is rather unstable, and extracting
  the archives as part of a Bazel repository mitigates this (as extractions happens on the host).

  Note #2: Additionally with the previous `platform_http_file` approach, a single repository
  had to be defined for all platforms. Only the archive for for the current host platform
  was pulled. This broke cross-compilation because the wrong platform archive would be use
  for web tests that run in the exec platform (and not in the host platform).
"""
platform_archive = repository_rule(
    implementation = _platform_archive_impl,
    attrs = {
        "urls": attr.string_list(
            doc = "URLs used for downloading the archive. Multiple URLs can be serve as fallback.",
            mandatory = True,
        ),
        "sha256": attr.string(
            doc = "SHA256 checksum for the archive.",
            mandatory = True,
        ),
        "licenses": attr.string_list(
            mandatory = True,
            allow_empty = False,
            doc = """
              Licenses that apply to the archive. Will be passed to a `licenses` invocation
              within the repository. https://docs.bazel.build/versions/0.24.0/be/functions.html#licenses.
            """,
        ),
        "named_files": attr.string_dict(
            doc = """
              Dictionary that maps files to unique identifiers. This is useful
              if browser archives are different on different platforms and the web
              tests would not want to care about archive-specific paths. e.g. targets
              expect a `CHROMIUM` key to point to the Chromium browser binary.
            """,
            mandatory = True,
        ),
    },
)
