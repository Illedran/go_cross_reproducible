load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "platforms",
    sha256 = "218efe8ee736d26a3572663b374a253c012b716d8af0c07e842e82f238a0a7ee",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/platforms/releases/download/0.0.10/platforms-0.0.10.tar.gz",
        "https://github.com/bazelbuild/platforms/releases/download/0.0.10/platforms-0.0.10.tar.gz",
    ],
)

http_archive(
    name = "bazel_features",
    sha256 = "5d7e4eb0bb17aee392143cd667b67d9044c270a9345776a5e5a3cccbc44aa4b3",
    strip_prefix = "bazel_features-1.13.0",
    url = "https://github.com/bazel-contrib/bazel_features/releases/download/v1.13.0/bazel_features-v1.13.0.tar.gz",
)

http_archive(
    name = "toolchains_protoc",
    sha256 = "3898f8e621ca5b3c7b94300c5ae19075d5bb28349d6e6f407490a466ba1e2544",
    strip_prefix = "toolchains_protoc-0.3.1",
    url = "https://github.com/aspect-build/toolchains_protoc/releases/download/v0.3.1/toolchains_protoc-v0.3.1.tar.gz",
)

http_archive(
    name = "rules_proto",
    sha256 = "6fb6767d1bef535310547e03247f7518b03487740c11b6c6adb7952033fe1295",
    strip_prefix = "rules_proto-6.0.2",
    url = "https://github.com/bazelbuild/rules_proto/releases/download/6.0.2/rules_proto-6.0.2.tar.gz",
)

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "b2038e2de2cace18f032249cb4bb0048abf583a36369fa98f687af1b3f880b26",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.48.1rules_go-v0.48.1.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.48.1/rules_go-v0.48.1.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    integrity = "sha256-12v3pg/YsFBEQJDfooN6Tq+YKeEWVhjuNdzspcvfWNU=",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.37.0/bazel-gazelle-v0.37.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.37.0/bazel-gazelle-v0.37.0.tar.gz",
    ],
)


load("@toolchains_protoc//protoc:repositories.bzl", "rules_protoc_dependencies")

rules_protoc_dependencies()

load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies")

rules_proto_dependencies()

load("@bazel_features//:deps.bzl", "bazel_features_deps")

bazel_features_deps()

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains(version = "1.22.5")

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

gazelle_dependencies()

load("@toolchains_protoc//protoc:toolchain.bzl", "protoc_toolchains")

protoc_toolchains(
    name = "protoc_toolchains",
    version = "v25.3",
)

load("@rules_proto//proto:setup.bzl", "rules_proto_setup")

rules_proto_setup()

load("@rules_proto//proto:toolchains.bzl", "rules_proto_toolchains")

rules_proto_toolchains()
