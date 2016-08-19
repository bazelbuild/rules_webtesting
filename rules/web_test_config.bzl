"""web_test_config.bzl defines the web_test_config rule.

The web_test_config rules provides configuration information such as
whether to record video to web_test.

DO NOT load this file. Use //testing/web/build_defs:web.bzl
"""

load("//rules:metadata.bzl", "create_file", "merge_files")


def _web_test_config_impl(ctx):
  """Implementation of the web_test_config rule."""
  files_to_merge = []
  record = ""
  for config in ctx.attr.configs:
    if config.record:
      record = config.record
    files_to_merge += [config.json]

  patch = ctx.new_file("%s.tmp.json" % ctx.label.name)
  create_file(ctx=ctx, output=patch, record_video=ctx.attr.record)
  files_to_merge += [patch]

  if ctx.attr.record:
    record = ctx.attr.record

  merge_files(
      ctx=ctx,
      merger=ctx.executable._merger,
      output=ctx.outputs.json,
      inputs=files_to_merge)

  return struct(record=record, json=ctx.outputs.json)


web_test_config = rule(
    implementation=_web_test_config_impl,
    attrs={
        "configs": attr.label_list(providers=["json", "record"]),
        "record": attr.string(
            default="", values=["", "never", "failed", "always"]),
        "_merger": attr.label(
            executable=True,
            cfg=HOST_CFG,
            default=Label("//external:web_test_merger")),
    },
    outputs={"json": "%{name}.gen.json"},)