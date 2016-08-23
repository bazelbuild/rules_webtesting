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

bind(
    name = "SeleniumServer",
    actual = "//java:SeleniumServer",
)

http_jar(
    name = "selenium_server",
    sha256 = "df874ce5b9508ac9f4ee0a3f50290836915c837b68975066a3841e839bc39804",
    url = "http://selenium-release.storage.googleapis.com/3.0-beta2/selenium-server-standalone-3.0.0-beta2.jar",
)
