java_test(
    name = "BrowserTest",
    srcs = ["BrowserTest.java"],
    tags = ["manual"],
    deps = [
        "//java/com/google/testing/web",
        "@com_google_truth_truth//jar",
        "@junit_junit//jar",
        "@org_seleniumhq_selenium//:selenium",
    ],
)

load("//web:web.bzl", "web_test_suite")

web_test_suite(
    name = "BrowserTest_web",
    browsers = [
        "//browsers:chrome-external",
        "//browsers:chrome-native",
        "//browsers:firefox-external",
        "//browsers:firefox-native",
    ],
    local = 1,
    test = ":BrowserTest",
)
