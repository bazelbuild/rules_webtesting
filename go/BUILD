package(
    default_testonly = 1,
    default_visibility = ["//visibility:public"],
)

load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")
load("//web:web.bzl", "web_test_suite")

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
    name = "browser_test_wrapped",
    srcs = ["browser_test.go"],
    library = ":browser",
    tags = ["manual"],
    deps = [
        "@com_github_tebeka_selenium//:selenium",
    ],
)

web_test_suite(
    name = "browser_test",
    browsers = [
        "//browsers:chrome-external",
        "//browsers:chrome-native",
        "//browsers:firefox-external",
        "//browsers:firefox-native",
    ],
    local = 1,
    test = ":browser_test_wrapped",
)
