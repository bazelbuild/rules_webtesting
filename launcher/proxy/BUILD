package(
    default_testonly = 1,
    default_visibility = ["//visibility:public"],
)

load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "proxy",
    srcs = ["proxy.go"],
    deps = [
        ":driverhub",
        "//launcher:errors",
        "//launcher:healthreporter",
        "//launcher/environments:environment",
        "//util:httphelper",
        "//util:portpicker",
    ],
)

go_library(
    name = "driverhub",
    srcs = [
        "driver_hub.go",
        "driver_responses.go",
        "driver_session.go",
    ],
    deps = [
        ":webdriver",
        "//launcher:errors",
        "//launcher:healthreporter",
        "//launcher/environments:environment",
        "//metadata:capabilities",
        "//util:httphelper",
        "@com_github_gorilla_mux//:mux",
    ],
)

go_library(
    name = "webdriver",
    srcs = ["webdriver.go"],
    deps = [
        "//launcher:errors",
        "//launcher:healthreporter",
    ],
)
