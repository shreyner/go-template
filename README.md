# Go template

docker-compose
```shell
$ docker compose -p go-tempalte up -d
```

build docker image
```shell
$ docker build -t go-restapi:latest .
```

start container
```shell
$ docker run --rm -p 8080:8080 -e DATABASE_URL="postgres://postgres:postgres@pg:5432/develop?sslmode=disable" --net=go-tempalte_default --name go-restapi go-restapi:latest
```

ENV
```text
PORT=8080
DATABASE_URL="postgres://postgres:postgres@pg:5432/develop?sslmode=disable"
```

migration 

Up
```shell
$ docker run --rm --net=go-tempalte_default -v $(pwd)/db/migrations:/flyway/sql -v $(pwd)/flyway.conf:/flyway/conf/flyway.conf flyway/flyway migrate  
```

