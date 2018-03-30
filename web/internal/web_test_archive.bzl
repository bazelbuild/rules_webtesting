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
"""A rule for configuring archives that contain named files.

DO NOT load this file. Use "@io_bazel_rules_web//web:web.bzl".
"""

load(":files.bzl", "files")
load(":metadata.bzl", "metadata")
load(":provider.bzl", "WebTestInfo")
load(":runfiles.bzl", "runfiles")


def _web_test_archive_impl(ctx):
  if ctx.attr.extract == "run":
    return _web_test_archive_run_impl(ctx)
  if ctx.attr.extract == "build":
    return _web_test_archive_build_impl(ctx)

  fail("unknown value %s" % ctx.attr.extract, attr="extract")


def _web_test_archive_run_impl(ctx):
  metadata.create_file(
      ctx=ctx,
      output=ctx.outputs.web_test_metadata,
      web_test_files=[
          metadata.web_test_files(
              ctx=ctx,
              archive_file=ctx.file.archive,
              named_files=ctx.attr.named_files,
              strip_prefix=ctx.attr.strip_prefix),
          metadata.web_test_files(
              ctx=ctx,
              named_files={
                  "EXTRACT_EXE": ctx.executable.extract_exe_target,
              }),
      ])

  return [
      DefaultInfo(
          runfiles=runfiles.collect(
              ctx=ctx,
              files=[ctx.file.archive],
              targets=[ctx.attr.extract_exe_target])),
      WebTestInfo(metadata=ctx.outputs.web_test_metadata),
  ]


def _web_test_archive_build_impl(ctx):
  out_dir = ctx.actions.declare_directory(ctx.label.name + ".out")
  ctx.actions.run(
      executable=ctx.executable.extract_exe_host,
      arguments=[ctx.file.archive.path, out_dir.path, ctx.attr.strip_prefix],
      mnemonic="WebTestArchive",
      progress_message="Extracting %s" % ctx.file.archive.short_path,
      use_default_shell_env=True,
      inputs=[ctx.file.archive],
      outputs=[out_dir])

  out_dir_path = files.long_path(ctx, out_dir)
  named_files = {}

  for n, p in ctx.attr.named_files.items():
    named_files[n] = out_dir_path + "/" + p

  metadata.create_file(
      ctx=ctx,
      output=ctx.outputs.web_test_metadata,
      web_test_files=[
          metadata.web_test_files(ctx=ctx, named_files=named_files),
      ])

  return [
      DefaultInfo(runfiles=runfiles.collect(ctx=ctx, files=[out_dir])),
      WebTestInfo(metadata=ctx.outputs.web_test_metadata),
  ]


web_test_archive = rule(
    doc="""Specifies an archive file with named files in it.

        If extract=="run", then the archive will only be extracted if WTL wants one
        of the named files in it.""",
    implementation=_web_test_archive_impl,
    attrs={
        "archive":
            attr.label(
                doc="Archive file that contains named files.",
                allow_single_file=[
                    ".deb",
                    ".tar",
                    ".tar.bz2",
                    ".tbz2",
                    ".tar.gz",
                    ".tgz",
                    ".tar.Z",
                    ".zip",
                ],
                mandatory=True),
        "named_files":
            attr.string_dict(
                doc="A map of names to paths in the archive.", mandatory=True),
        "extract":
            attr.string(
                doc="When the archive shoud be extracted.",
                default="run",
                values=["build", "run"]),
        "strip_prefix":
            attr.string(doc="""Prefix to strip when archive is extracted. 
                BASH-style globbing is allowed."""),
        "extract_exe_host":
            attr.label(
                doc="""Executable to extract files if extract = build.
                    Should accept three positional parameters: 
                      archive out_dir strip_prefix""",
                allow_files=True,
                cfg="host",
                default=Label("//web/internal:extract.sh"),
                executable=True),
        "extract_exe_target":
            attr.label(
                doc="""Executable to extract files if extract = run.
                    Should accept three positional parameters: 
                      archive out_dir strip_prefix""",
                allow_files=True,
                cfg="target",
                default=Label("//web/internal:extract.sh"),
                executable=True),
    },
    outputs={"web_test_metadata": "%{name}.gen.json"},
)
