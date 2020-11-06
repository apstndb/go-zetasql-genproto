```
$ docker run -p 50051:50051 ghcr.io/apstndb/zetasql-local-service-docker:v0.1.0
$ go run ./ "SELECT * FROM (SELECT 1.0 AS a)"  
```
