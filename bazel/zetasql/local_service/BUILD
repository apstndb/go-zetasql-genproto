load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

package(default_visibility = ["//visibility:public"])

go_proto_library(
    name = "local_service_go_proto",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "github.com/apstndb/go-zetasql-genproto/zetasql/local_service/local_servicepb",
    proto = "@com_google_zetasql//zetasql/local_service:local_service_proto",
    visibility = ["//visibility:public"],
    deps = [
        "//bazel/zetasql/proto:function_go_proto",
        "//bazel/zetasql/proto:options_go_proto",
        "//bazel/zetasql/proto:simple_catalog_go_proto",
        "//bazel/zetasql/public:options_go_proto",
        "//bazel/zetasql/public:parse_resume_location_go_proto",
        "//bazel/zetasql/public:simple_table_go_proto",
        "//bazel/zetasql/public:type_go_proto",
        "//bazel/zetasql/public:value_go_proto",
        "//bazel/zetasql/resolved_ast:resolved_ast_go_proto",
    ],
)
