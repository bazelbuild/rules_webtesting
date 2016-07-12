load("@io_bazel_rules_go//go:def.bzl", "go_binary")

go_binary(
    name = "main",
    srcs = ["main.go"],
    deps = [
        "@com_github_golang_net//:context",
        "@com_github_golang_net//:context/ctxhttp",
    ],
)
