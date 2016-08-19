"""The browser rule is used to define web_test browsers.

DO NOT load this file. Use //testing/web/build_defs:web.bzl
"""
load("//rules:metadata.bzl", "create_file", "merge_files")


def _browser_impl(ctx):
  """Implementation of the browser rule."""
  patch = ctx.new_file("%s.tmp.json" % ctx.label.name)
  create_file(ctx=ctx, output=patch, browser_label=ctx.label)
  merge_files(
      ctx=ctx,
      merger=ctx.executable._merger,
      output=ctx.outputs.json,
      inputs=[ctx.file.metadata, patch])

  return struct(
      runfiles=ctx.runfiles(collect_default=True, collect_data=True),
      browser=ctx.attr.browser,
      name=ctx.label.name,
      disabled=ctx.attr.disabled,
      json=ctx.outputs.json)


browser = rule(
    implementation=_browser_impl,
    attrs={
        "browser": attr.string(mandatory=True),
        "metadata": attr.label(
            mandatory=True, allow_files=True, single_file=True, cfg=DATA_CFG),
        "data": attr.label_list(
            allow_files=True, cfg=DATA_CFG),
        "disabled": attr.string(),
        "_merger": attr.label(
            executable=True,
            cfg=HOST_CFG,
            default=Label("//external:web_test_merger")),
    },
    outputs={"json": "%{name}.gen.json"},)