package(
    default_testonly = 1,
    default_visibility = ["//visibility:public"],
)

java_library(
    name = "web",
    srcs = glob(["*.java"]),
    deps = [
        "//java/com/google/testing/util",
        "@com_google_code_findbugs_jsr305//jar",
        "@com_google_guava_guava//jar",
        "@org_json_json//jar",
        "@org_seleniumhq_selenium//:selenium",
    ],
)
