# Copyright 2024 Google Inc.
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
"""Module extension for introducing browsers dependencies."""

load("//web/versioned:browsers-0.3.1.bzl", browser_repositories_0_3_1 = "browser_repositories")
load("//web/versioned:browsers-0.3.2.bzl", browser_repositories_0_3_2 = "browser_repositories")
load("//web/versioned:browsers-0.3.3.bzl", browser_repositories_0_3_3 = "browser_repositories")
load("//web/versioned:browsers-0.3.4.bzl", browser_repositories_0_3_4 = "browser_repositories")

browser_versions = {
    "0.3.1": browser_repositories_0_3_1,
    "0.3.2": browser_repositories_0_3_2,
    "0.3.3": browser_repositories_0_3_3,
    "0.3.4": browser_repositories_0_3_4,
}

def parse_version(version):
    if not version:
        return [0, 0, 0]
    return [int(s) for s in version.split(".")]

def _browser_repositories_extension(ctx):
    version = None

    # Get the override version from the root module if specified.
    root_module = ctx.modules[0]
    if len(root_module.tags.override_version) > 1:
        fail("Only one override_version tag is allowed in the root module.")
    override_version = root_module.tags.override_version[0].version if root_module.tags.override_version else None

    if override_version:
        version = override_version
    else:
        # Go through modules in the dependency graph and choose the highest version.
        for mod in ctx.modules:
            for tag in mod.tags.install:
                if parse_version(tag.version) > parse_version(version):
                    version = tag.version

    if not version:
        fail("No version of browser_repositories is specified in the dependency graph, supported versions are %s." % ', '.join(browser_versions.keys()))

    browser_repositories = browser_versions.get(version)
    if not browser_repositories:
        fail("Unsupported version %s of browser_repositories, supported versions are %s." % (version, ', '.join(browser_versions.keys())))

    browser_repositories()

    # Mark this module extension as reproducible, so it doesn't appear in lockfile.
    return ctx.extension_metadata(reproducible = True)

# The `install` tag can be used by any module to specify the version of browsers version.
# If multiple versions are specified, the highest one is chosen.
install = tag_class(
    attrs = {
        "version": attr.string(mandatory = True),
    },
)

# The `override_version` tag can be used by root module to override the version of browsers version.
override_version = tag_class(
    attrs = {
        "version": attr.string(mandatory = True),
    },
)

browser_repositories_extension = module_extension(
    implementation = _browser_repositories_extension,
    tag_classes = {
        "install": install,
        "override_version": override_version,
    }
)
