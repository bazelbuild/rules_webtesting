package(
    default_testonly = 1,
    default_visibility = ["//visibility:public"],
)

load("//web:web.bzl", "browser")

browser(
    name = "firefox-external",
    browser = "firefox",
    metadata = "firefox-external.json",
)

browser(
    name = "chrome-external",
    browser = "chrome",
    metadata = "chrome-external.json",
)

browser(
    name = "firefox-native",
    browser = "firefox",
    data = ["//java:SeleniumServer"],
    metadata = "firefox-native.json",
)

browser(
    name = "chrome-native",
    browser = "chrome",
    data = ["//java:SeleniumServer"],
    metadata = "chrome-native.json",
)
