load("@io_bazel_rules_go//go:def.bzl", "go_binary")

go_binary(
    name = "build",
    srcs = glob([
        "*.go",
    ]),
    importpath = "github.com/Boot-Error/golang-bazel-learn/service-apiserver",
    visibility = ["//visibility:public"],
    deps = [
        "//lib-logging:build",
    ],
)