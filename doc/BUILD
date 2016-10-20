load("@io_bazel_skydoc//skylark:skylark.bzl", "skylark_doc")

skylark_doc(
    name = "web-docs",
    srcs = [
        "//web:repositories.bzl",
        "//web:web.bzl",
        "//web/internal:browser.bzl",
        "//web/internal:web_test.bzl",
        "//web/internal:web_test_config.bzl",
        "//web/internal:web_test_named_executable.bzl",
    ],
    tags = ["manual"],
)
