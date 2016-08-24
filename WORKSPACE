workspace(name = "web_test_rules")

load("//web:repositories.bzl", "web_test_repositories")

web_test_repositories(prefix = "")

load("//web:bindings.bzl", "web_test_bindings")

web_test_bindings(
    default_config = "//rules:default",
    launcher = "//launcher:main",
    merger = "//metadata:merger",
)

new_git_repository(
    name = "com_github_tebeka_selenium",
    build_file = "BUILD.selenium_go",
    remote = "https://github.com/tebeka/selenium.git",
    tag = "v0.8.5",
)

maven_jar(
    name = "com_google_guava_guava",
    artifact = "com.google.guava:guava:19.0",
)

new_http_archive(
    name = "org_seleniumhq_selenium",
    build_file = "BUILD.selenium_java",
    url = "http://selenium-release.storage.googleapis.com/3.0-beta2/selenium-java-3.0.0-beta2.zip",
)

maven_jar(
    name = "org_json_json",
    artifact = "org.json:json:20160810",
)

maven_jar(
    name = "com_google_code_findbugs_jsr305",
    artifact = "com.google.code.findbugs:jsr305:3.0.1",
)

maven_jar(
    name = "junit_junit",
    artifact = "junit:junit:4.12",
)

maven_jar(
    name = "com_google_truth_truth",
    artifact = "com.google.truth:truth:0.29",
)
