load("//web/internal:collections.bzl", "lists")

def _web_test_metadata_aspect_impl(target, ctx):
    transitive_files = get_files(ctx.rule, ["data", "deps", "srcs"])

    if hasattr(target, "web_test_metadata"):
        for m in target.web_test_metadata:
            lists.ensure_at_end_of_list(transitive_files, m)

    print(ctx.label)
    print(dir(ctx.rule))
    print([t.short_path for t in transitive_files])
    return struct(transitive_web_test_metadata=transitive_files)


web_test_metadata_aspect = aspect(
    implementation =_web_test_metadata_aspect_impl,
    attr_aspects = ["deps", "data", "srcs", "browser", "config"],
)


def get_files(ctx, attr_names):
  """Finds all of the web_test_metadata files in given attributes.

  It ensures that all web_test_metadata files position in the returned list is
  based on the last attribute in attr_names that they appear in.

  Args:
    ctx: the context object for this rule.
    attr_names: the names of attributes to get web_test_metadata files from.
  Returns:
    a list of all of the web_test_metadata files in the specificied attributes.
  """
  metadata_files = []

  for attr_name in attr_names:
    if hasattr(ctx.attr, attr_name):
        attr = getattr(ctx.attr, attr_name)
        if hasattr(attr, 'transitive_web_test_metadata'):
            for l in attr.transitive_web_test_metadata:
                lists.ensure_at_end_of_list(metadata_files, l)
        elif lists.is_list_like(attr):
          for a in attr:
            if hasattr(a, 'transitive_web_test_metadata'):
                for l in a.transitive_web_test_metadata:
                    lists.ensure_at_end_of_list(metadata_files, l)

  return metadata_files
