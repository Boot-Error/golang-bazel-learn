load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "build",
    srcs = glob([
        "*.go",
    ]),
    importpath = "github.com/Boot-Error/golang-bazel-learn/lib-logging",
    visibility = ["//visibility:public"],
    deps = [
        "//lib-logging/constants:build",
    ],
)