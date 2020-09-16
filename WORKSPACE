workspace(name = "istio_ecosystem_wasm_extensions")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "proxy_wasm_cpp_sdk",
    sha256 = "06f0f386dc8111082062f01e74e0c297e4a83857585519adb8727a3e7170f3b7",
    strip_prefix = "proxy-wasm-cpp-sdk-558d45a2f496e3039e50584cf8304ae535ca73de",
    url = "https://github.com/proxy-wasm/proxy-wasm-cpp-sdk/archive/558d45a2f496e3039e50584cf8304ae535ca73de.tar.gz",
)

http_archive(
    name = "emscripten_toolchain",
    build_file = "//bazel/wasm:emscripten-toolchain.BUILD",
    patch_cmds = [
        "./emsdk install 1.39.6-upstream",
        "./emsdk activate --embedded 1.39.6-upstream",
    ],
    strip_prefix = "emsdk-1.39.6",
    url = "https://github.com/emscripten-core/emsdk/archive/1.39.6.tar.gz",
)
