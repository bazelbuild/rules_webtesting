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
"""A library of functions for working with web_test metadata files."""

load(":files.bzl", "files")

def _merge_files(ctx, merger, output, inputs):
    """Produces a merged web test metadata file.

    Args:
        ctx: a Skylark rule context.
        merger: the WTL metadata merger executable.
        output: a File object for the output file.
        inputs: a list of File objects. These files are in order of priority;
          i.e. values in the first file will take precedence over values in the
          second file, etc.
    """
    paths = [i.path for i in reversed(inputs)]
    short_paths = [i.short_path for i in inputs]
    args = ["--output", output.path] + paths

    ctx.actions.run(
        outputs = [output],
        inputs = inputs,
        executable = merger,
        arguments = args,
        mnemonic = "METADATAMERGER",
        progress_message = "merging %s" % (", ".join(short_paths)),
    )

def _create_file(
        ctx,
        output,
        browser_label = None,
        capabilities = None,
        config_label = None,
        environment = None,
        label = None,
        test_label = None,
        web_test_files = None,
        extension = None):
    """Generates a web_test metadata file with specified contents.

    Args:
        ctx: a Skylark rule context.
        output: File object. The file to write the metadata to.
        browser_label: Label. The label for a browser rule.
        capabilities: struct; Browser capabilities.
        config_label: Label. The label for the web_test_config rule.
        environment: string. The Web Test Launcher environment name.
        label: Label. The label for this target.
        test_label: Label. The label for the test being executed.
        web_test_files: sequence of web_test_file structs.
        extension: map or struct defining additional fields that should be added
          metadata file.
    """
    fields = {}
    if browser_label:
        fields["browserLabel"] = str(browser_label)
    if capabilities:
        fields["capabilities"] = capabilities
    if config_label:
        fields["configLabel"] = str(config_label)
    if environment:
        fields["environment"] = environment
    if label:
        fields["label"] = str(label)
    if test_label:
        fields["testLabel"] = str(test_label)
    if web_test_files:
        fields["webTestFiles"] = web_test_files
    if extension:
        if type(extension) == type({}):
            extension = struct(**extension)
        fields["extension"] = extension

    ctx.actions.write(
        output = output,
        content = json.encode(struct(**fields)),
        is_executable = False,
    )

def _web_test_files(ctx, archive_file = None, named_files = None, strip_prefix = ""):
    """Build a web_test_files struct.

    Args:
        ctx: a Skylark rule context.
        archive_file: a File object. The archive file where the named_files will be
          found. If absent, the named_files are located directly in the runfiles.
        named_files: a dict of strings to strings or File objects. The mapping of
          names to file path. If archive_file is absent, the values should be
          File objects for files that will be in the runfiles of the test. If
          archive_file is present, the values should be string paths referencing
          files in archive_file.
        strip_prefix: string. A prefix to strip from the begining of paths in an archive.

    Returns:
        A web_test_files struct.
    """
    named_files = named_files or {}
    for k, v in named_files.items():
        if type(v) != type(""):
            named_files[k] = files.long_path(ctx, v)
    if archive_file:
        archive_file = files.long_path(ctx, archive_file)
    return struct(
        archiveFile = archive_file,
        namedFiles = struct(**named_files),
        stripPrefix = strip_prefix,
    )

metadata = struct(
    create_file = _create_file,
    merge_files = _merge_files,
    web_test_files = _web_test_files,
)
