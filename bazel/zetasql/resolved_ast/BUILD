load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

package(default_visibility = ["//visibility:public"])

go_proto_library(
    name = "resolved_node_kind_go_proto",
    importpath = "github.com/apstndb/go-zetasql-genproto/zetasql/resolved_ast/resolved_node_kindpb",
    protos = ["@com_google_zetasql//zetasql/resolved_ast:resolved_node_kind_proto"],
    visibility = ["//visibility:public"],
)

go_proto_library(
    name = "resolved_ast_enums_go_proto",
    importpath = "github.com/apstndb/go-zetasql-genproto/zetasql/resolved_ast/resolved_ast_enumspb",
    protos = ["@com_google_zetasql//zetasql/resolved_ast:resolved_ast_enums_proto"],
    visibility = ["//visibility:public"],
)

go_proto_library(
    name = "resolved_ast_go_proto",
    importpath = "github.com/apstndb/go-zetasql-genproto/zetasql/resolved_ast/resolved_astpb",
    protos = ["@com_google_zetasql//zetasql/resolved_ast:resolved_ast_proto"],
    visibility = ["//visibility:public"],
    deps = [
        "//bazel/zetasql/proto:function_go_proto",
        "//bazel/zetasql/public:type_go_proto",
        "//bazel/zetasql/public/proto:type_annotation_go_proto",
        "//bazel/zetasql/resolved_ast:resolved_ast_enums_go_proto",
        "//bazel/zetasql/resolved_ast:resolved_node_kind_go_proto",
        "//bazel/zetasql/resolved_ast:serialization_go_proto",
    ],
)

go_proto_library(
    name = "serialization_go_proto",
    importpath = "github.com/apstndb/go-zetasql-genproto/zetasql/resolved_ast/serializationpb",
    protos = ["@com_google_zetasql//zetasql/resolved_ast:serialization_proto"],
    deps = [
        "//bazel/zetasql/public:parse_location_range_go_proto",
        "//bazel/zetasql/public:type_go_proto",
        "//bazel/zetasql/public:value_go_proto",
    ],
)
