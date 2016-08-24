load("@io_bazel_rules_go//go:def.bzl", "go_repositories")


def web_test_bindings(launcher="@web_test_rules//launcher:main",
                      merger="@web_test_rules//metadata:merger",
                      default_config="@web_test_rules//rules:default"):
  native.bind(
      name="web_test_launcher",
      actual=launcher,)

  native.bind(
      name="web_test_merger",
      actual=merger,)

  native.bind(
      name="web_test_default_config",
      actual=default_config,)

  go_repositories()
