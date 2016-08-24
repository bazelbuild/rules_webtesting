def web_test_repositories(prefix="@web_test_rules", java=False, go=False):
  native.git_repository(
      name="io_bazel_rules_go",
      commit="ae8ea32be1af991eef77d6347591dc8ba56c40a2",
      remote="https://github.com/bazelbuild/rules_go.git",)

  native.new_git_repository(
      name="com_github_gorilla_mux",
      build_file=prefix + "//:BUILD.gorilla_mux",
      commit="cf79e51a62d8219d52060dfc1b4e810414ba2d15",
      remote="https://github.com/gorilla/mux.git",)

  native.http_jar(
      name="org_seleniumhq_server",
      sha256="df874ce5b9508ac9f4ee0a3f50290836915c837b68975066a3841e839bc39804",
      url="http://selenium-release.storage.googleapis.com/3.0-beta2/selenium-server-standalone-3.0.0-beta2.jar",
  )

  if java:
    native.new_http_archive(
        name="org_seleniumhq_java",
        build_file=prefix + "//:BUILD.selenium_java",
        url="http://selenium-release.storage.googleapis.com/3.0-beta2/selenium-java-3.0.0-beta2.zip",
    )

    native.maven_jar(
        name="org_json_json",
        artifact="org.json:json:20160810",)

    native.maven_jar(
        name="com_google_code_findbugs_jsr305",
        artifact="com.google.code.findbugs:jsr305:3.0.1",)

    native.maven_jar(
        name="com_google_guava_guava",
        artifact="com.google.guava:guava:19.0",)

  if go:
    native.new_git_repository(
        name="com_github_tebeka_selenium",
        build_file=prefix + "//:BUILD.selenium_go",
        remote="https://github.com/tebeka/selenium.git",
        tag="v0.8.5",)
