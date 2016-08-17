package(
    default_testonly = 1,
    default_visibility = ["//visibility:public"],
)

load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_binary(
    name = "merger",
    srcs = ["merger.go"],
    deps = [
        ":metadata",
    ],
)

go_library(
    name = "metadata",
    srcs = ["metadata.go"],
    deps = [
        ":capabilities",
    ],
)

go_library(
    name = "capabilities",
    srcs = ["capabilities.go"],
)

go_test(
    name = "metadata_test",
    srcs = ["metadata_test.go"],
    library = ":metadata",
    deps = [
        "//util:bazel",
    ],
    data = [
        "//metadata/testdata:all-fields.json",
        "//metadata/testdata:android-browser-gingerbread-nexus-s.json",
        "//metadata/testdata:android-browser-gingerbread-nexus-s_faked_on_chrome-linux.json",
        "//metadata/testdata:chrome-linux.json",
    ]
)

sh_test(
    name = "merger_test",
    srcs = ["merger_test.sh"],
    data = [
        ":merger",
        "//metadata/testdata:android-browser-gingerbread-nexus-s.json",
        "//metadata/testdata:android-browser-gingerbread-nexus-s_faked_on_chrome-linux.json",
        "//metadata/testdata:chrome-linux.json",
    ]
)