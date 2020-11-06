package main

import (
	"context"
	"fmt"
	"log"

	local_servicepb "github.com/apstndb/go-zetasql-genproto/zetasql/local_service/local_servicepb"
	"google.golang.org/grpc"
)

func String(v string) *string {
	return &v
}
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
	ctx := context.Background()
	client := local_servicepb.NewZetaSqlLocalServiceClient(conn)
	resp, err := client.FormatSql(ctx, &local_servicepb.FormatSqlRequest{Sql: String("SELECT * FROM `tbl`")})
	if err != nil {
		return err
	}
	fmt.Println(resp.GetSql())
	return nil
}
