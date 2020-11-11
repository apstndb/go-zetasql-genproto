```
$ docker run -p 50051:50051 gcr.io/apstndb-sandbox/zetasql-local-service-docker:latest
$ go run ./ "SELECT * FROM (SELECT 1.0 AS a)"
```