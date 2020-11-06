load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

package(default_visibility = ["//visibility:public"])

go_proto_library(
    name = "options_go_proto",
    importpath = "github.com/apstndb/go-zetasql-genproto/zetasql/proto/optionspb",
    protos = ["@com_google_zetasql//zetasql/proto:options_proto"],
    deps = [
        "//bazel/zetasql/public:builtin_function_go_proto",
        "//bazel/zetasql/public:options_go_proto",
        "//bazel/zetasql/public:type_go_proto",
        "//bazel/zetasql/resolved_ast:resolved_node_kind_go_proto",
    ],
)

go_proto_library(
    name = "simple_catalog_go_proto",
    importpath = "github.com/apstndb/go-zetasql-genproto/zetasql/proto/simplecatalogpb",
    protos = ["@com_google_zetasql//zetasql/proto:simple_catalog_proto"],
    deps = [
        ":function_go_proto",
        ":options_go_proto",
        "//bazel/zetasql/public:simple_constant_go_proto",
        "//bazel/zetasql/public:simple_table_go_proto",
        "//bazel/zetasql/public:type_go_proto",
    ],
)

go_proto_library(
    name = "function_go_proto",
    importpath = "github.com/apstndb/go-zetasql-genproto/zetasql/proto/functionpb",
    protos = ["@com_google_zetasql//zetasql/proto:function_proto"],
    deps = [
        "//bazel/zetasql/public:deprecation_warning_go_proto",
        "//bazel/zetasql/public:function_go_proto",
        "//bazel/zetasql/public:options_go_proto",
        "//bazel/zetasql/public:parse_location_range_go_proto",
        "//bazel/zetasql/public:parse_resume_location_go_proto",
        "//bazel/zetasql/public:type_go_proto",
        "//bazel/zetasql/public:value_go_proto",
        "//bazel/zetasql/resolved_ast:serialization_go_proto",
    ],
)
