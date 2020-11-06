load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

package(default_visibility = ["//visibility:public"])

go_proto_library(
    name = "builtin_function_go_proto",
    importpath = "github.com/apstndb/go-zetasql-genproto/zetasql/public/builtin_functionpb",
    protos = ["@com_google_zetasql//zetasql/public:builtin_function_proto"],
)

go_proto_library(
    name = "deprecation_warning_go_proto",
    importpath = "github.com/apstndb/go-zetasql-genproto/zetasql/public/deprecation_warningpb",
    protos = [
        "@com_google_zetasql//zetasql/public:deprecation_warning_proto",
    ],
    deps = [":error_location_go_proto"],
)

go_proto_library(
    name = "function_go_proto",
    importpath = "github.com/apstndb/go-zetasql-genproto/zetasql/public/functionpb",
    protos = ["@com_google_zetasql//zetasql/public:function_proto"],
)

go_proto_library(
    name = "parse_resume_location_go_proto",
    importpath = "github.com/apstndb/go-zetasql-genproto/zetasql/public/parse_resume_locationpb",
    protos = ["@com_google_zetasql//zetasql/public:parse_resume_location_proto"],
)

go_proto_library(
    name = "options_go_proto",
    importpath = "github.com/apstndb/go-zetasql-genproto/zetasql/public/optionspb",
    protos = ["@com_google_zetasql//zetasql/public:options_proto"],
)

go_proto_library(
    name = "simple_table_go_proto",
    importpath = "github.com/apstndb/go-zetasql-genproto/zetasql/public/simple_tablepb",
    protos = [
        "@com_google_zetasql//zetasql/public:simple_table_proto",
    ],
    deps = [":type_go_proto"],
)

go_proto_library(
    name = "simple_constant_go_proto",
    importpath = "github.com/apstndb/go-zetasql-genproto/zetasql/public/simple_constantpb",
    protos = [
        "@com_google_zetasql//zetasql/public:simple_constant_proto",
    ],
    deps = [
        ":type_go_proto",
        ":value_go_proto",
    ],
)

go_proto_library(
    name = "parse_location_range_go_proto",
    importpath = "github.com/apstndb/go-zetasql-genproto/zetasql/public/parse_location_rangepb",
    protos = ["@com_google_zetasql//zetasql/public:parse_location_range_proto"],
)

go_proto_library(
    name = "type_go_proto",
    importpath = "github.com/apstndb/go-zetasql-genproto/zetasql/public/typepb",
    protos = ["@com_google_zetasql//zetasql/public:type_proto"],
)

go_proto_library(
    name = "type_annotation_go_proto",
    importpath = "github.com/apstndb/go-zetasql-genproto/zetasql/public/type_annotationpb",
    protos = ["@com_google_zetasql//zetasql/public:type_annotation_proto"],
    deps = ["//bazel/zetasql/public/proto:type_annotation_go_proto"],
)

go_proto_library(
    name = "value_go_proto",
    importpath = "github.com/apstndb/go-zetasql-genproto/zetasql/public/valuepb",
    protos = ["@com_google_zetasql//zetasql/public:value_proto"],
)

go_proto_library(
    name = "error_location_go_proto",
    importpath = "github.com/apstndb/go-zetasql-genproto/zetasql/public/error_locationpb",
    protos = ["@com_google_zetasql//zetasql/public:error_location_proto"],
)
