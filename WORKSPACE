load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "2f01d6eb4121eae9fd7df78ac2b3fa910ad873e857a8b3560f78a585994a14d4",
    strip_prefix = "rules_go-14edd3aaf0c7c0a0cd24fa47b8d0832778770570",
    urls = ["https://github.com/bazelbuild/rules_go/archive/14edd3aaf0c7c0a0cd24fa47b8d0832778770570.zip"],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains(
    version = "1.19.3",
    experiments = ["boringcrypto"],
)