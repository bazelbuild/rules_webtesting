package(
    default_testonly = 1,
    default_visibility = ["//visibility:public"],
)

load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "environment",
    srcs = ["environment.go"],
    deps = [
        "//launcher:errors",
        "//launcher:healthreporter",
        "//metadata",
        "//metadata:capabilities",
    ],
)

go_library(
    name = "external",
    srcs = ["external.go"],
    deps = [
        ":environment",
        "//launcher:errors",
        "//metadata",
    ],
)

go_library(
    name = "native",
    srcs = ["native.go"],
    deps = [
        ":environment",
        "//launcher/services:selenium",
        "//metadata",
    ],
)
