workspace(name = "web_test_rules")

load("//web:repositories.bzl", "web_test_repositories")

web_test_repositories(
    go = True,
    java = True,
    prefix = "",
)

load("//web:bindings.bzl", "web_test_bindings")

web_test_bindings(
    default_config = "//rules:default",
    launcher = "//launcher:main",
    merger = "//metadata:merger",
)

maven_jar(
    name = "junit_junit",
    artifact = "junit:junit:4.12",
)

maven_jar(
    name = "com_google_truth_truth",
    artifact = "com.google.truth:truth:0.29",
)
