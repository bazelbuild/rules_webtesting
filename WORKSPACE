git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    tag = "0.0.4",
)

load("@io_bazel_rules_go//go:def.bzl", "go_repositories")

go_repositories()

new_git_repository(
    name = "com_github_golang_net",
    build_file = "BUILD.net",
    commit = "f315505cf3349909cdf013ea56690da34e96a451",
    remote = "https://github.com/golang/net.git",
)

new_git_repository(
    name = "com_github_tebeka_selenium",
    build_file = "BUILD.selenium",
    remote = "https://github.com/tebeka/selenium.git",
    tag = "v0.8.5"
)
