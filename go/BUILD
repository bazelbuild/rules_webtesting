package(
    default_testonly = 1,
    default_visibility = ["//visibility:public"],
)

load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")
load("//web:web.bzl", "web_test")

go_library(
    name = "browser",
    srcs = ["browser.go"],
    deps = [
        "//metadata",
        "//util:bazel",
        "@com_github_tebeka_selenium//:selenium",
    ],
)

go_test(
    name = "browser_test",
    srcs = ["browser_test.go"],
    library = ":browser",
    deps = [
        "@com_github_tebeka_selenium//:selenium",
    ],
)

web_test(
    name = "browser_test_external",
    browser = "//browsers:external",
    local = 1,
    test = ":browser_test",
)
