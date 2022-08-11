# Go template


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

