package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

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

	if len(os.Args) != 2 {
		return errors.New("args must be formatsql [SQL]")
	}
	ctx := context.Background()
	client := local_servicepb.NewZetaSqlLocalServiceClient(conn)
	resp, err := client.FormatSql(ctx, &local_servicepb.FormatSqlRequest{Sql: String(os.Args[1])})
	if err != nil {
		return err
	}
	fmt.Println(resp.GetSql())
	return nil
}
