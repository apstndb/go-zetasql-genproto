package main

import (
	"context"
	"errors"
	"log"
	"os"

	local_servicepb "github.com/apstndb/go-zetasql-genproto/zetasql/local_service/local_servicepb"
	optionspb "github.com/apstndb/go-zetasql-genproto/zetasql/proto/optionspb"
	puboptionspb "github.com/apstndb/go-zetasql-genproto/zetasql/public/optionspb"
	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc"
)

func main() {
	if err := _main(); err != nil {
		log.Fatalln(err)
	}
}

func _main() error {
	address := "localhost:50051"
	conn, err := grpc.Dial(
		address,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		return err
	}

	if len(os.Args) != 2 {
		return errors.New("args must be formatsql [SQL]")
	}
	ctx := context.Background()
	client := local_servicepb.NewZetaSqlLocalServiceClient(conn)
	resp, err := client.Analyze(ctx, &local_servicepb.AnalyzeRequest{
		Options: &optionspb.AnalyzerOptionsProto{
			LanguageOptions: &optionspb.LanguageOptionsProto{
				ProductMode: puboptionspb.ProductMode_PRODUCT_EXTERNAL.Enum(),
			},
		},
		Target: &local_servicepb.AnalyzeRequest_SqlStatement{SqlStatement: os.Args[1]},
	})
	if err != nil {
		return err
	}

	var enc jsonpb.Marshaler
	return enc.Marshal(os.Stdout, resp.GetResolvedStatement())
}
