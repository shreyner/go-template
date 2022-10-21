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

# Migrate

https://github.com/jackc/tern

up
```shell
DATABASE_URI=postgres://postgres:postgres@localhost:5432/develop?sslmode=disable tern migrate -m migrate
```

down
```shell
DATABASE_URI=postgres://postgres:postgres@localhost:5432/develop?sslmode=disable  tern migrate --destination -1 -m migrate
```

create
```shell
DATABASE_URI=postgres://postgres:postgres@localhost:5432/develop?sslmode=disable  tern new user-and-add -m migrate
```


