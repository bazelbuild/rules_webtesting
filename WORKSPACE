git_repository(
    name = "io_bazel_rules_go",
    commit = "ae8ea32be1af991eef77d6347591dc8ba56c40a2",
    remote = "https://github.com/bazelbuild/rules_go.git",
)

load("@io_bazel_rules_go//go:def.bzl", "go_repositories")

go_repositories()

new_git_repository(
    name = "com_github_tebeka_selenium",
    build_file = "BUILD.selenium",
    remote = "https://github.com/tebeka/selenium.git",
    tag = "v0.8.5",
)

new_git_repository(
    name = "com_github_gorilla_mux",
    build_file = "BUILD.gorilla_mux",
    commit = "cf79e51a62d8219d52060dfc1b4e810414ba2d15",
    remote = "https://github.com/gorilla/mux.git",
)

bind(
    name = "web_test_launcher",
    actual = "//launcher:main",
)

bind(
    name = "web_test_merger",
    actual = "//metadata:merger",
)

bind(
    name = "web_test_default_config",
    actual = "//rules:default",
)

http_jar(
    name = "selenium_server",
    url = "http://goo.gl/2lZ46z",
)
