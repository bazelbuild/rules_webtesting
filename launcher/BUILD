package(
    default_testonly = 1,
    default_visibility = ["//visibility:public"],
)

load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")
load("//web:web.bzl", "web_test")

go_binary(
    name = "main",
    srcs = ["main.go"],
    deps = [
        ":cmdhelper",
        "//launcher/environments:environment",
        "//launcher/environments:external",
        "//launcher/environments:native",
        "//launcher/proxy",
        "//launcher/services:selenium",
        "//metadata",
        "//util:bazel",
    ],
)

go_library(
    name = "errors",
    srcs = ["errors.go"],
)

go_test(
    name = "errors_test",
    srcs = ["errors_test.go"],
    library = ":errors",
)

go_library(
    name = "healthreporter",
    srcs = ["health_reporter.go"],
    deps = [":errors"],
)

go_test(
    name = "healthreporter_test",
    srcs = ["health_reporter_test.go"],
    library = ":healthreporter",
)

go_library(
    name = "cmdhelper",
    srcs = ["cmd_helper.go"],
    deps = [":errors"],
)

sh_test(
    name = "example_test",
    srcs = ["example_test.sh"],
)

web_test(
    name = "test",
    browser = "//browsers:external",
    test = ":example_test",
)
