git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    tag = "0.0.3",
)

load("@io_bazel_rules_go//go:def.bzl", "go_repositories")

go_repositories()

new_git_repository(
    name = "com_github_golang_net",
    build_file = "BUILD.net",
    commit = "62685c2d7ca23c807425dca88b11a3e2323dab41",
    remote = "https://github.com/golang/net.git",
)
