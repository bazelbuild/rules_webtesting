package(
    default_testonly = 1,
    default_visibility = ["//visibility:public"],
)

load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "service",
    srcs = [
        "cmd.go",
        "server.go",
        "service.go",
    ],
    deps = [
        "//launcher:cmdhelper",
        "//launcher:errors",
        "//launcher:healthreporter",
        "//util:bazel",
        "//util:httphelper",
        "//util:portpicker",
    ],
)

go_library(
    name = "selenium",
    srcs = ["selenium.go"],
    deps = [":service"],
)
