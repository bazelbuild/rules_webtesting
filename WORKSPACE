git_repository(
    name = "io_bazel_rules_go",
    commit = "ae8ea32be1af991eef77d6347591dc8ba56c40a2",
    remote = "https://github.com/bazelbuild/rules_go.git",
)

load("@io_bazel_rules_go//go:def.bzl", "go_repositories")

go_repositories()

new_git_repository(
    name = "com_github_golang_net",
    build_file = "BUILD.net",
    commit = "f315505cf3349909cdf013ea56690da34e96a451",
    remote = "https://github.com/golang/net.git",
)
