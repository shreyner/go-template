FROM golang:1.19.2 as module

ADD go.mod go.sum /m/
RUN cd /m && go mod download

FROM golang:1.19.2 as builder

COPY --from=module /go/pkg /go/pkg

RUN useradd -u 10001 application

RUN mkdir -p /application
ADD . /application
WORKDIR /application

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
    go build -o ./bin/application ./cmd/app/main.go

FROM scratch

COPY --from=builder /etc/passwd /etc/passwd
USER application

COPY --from=builder /application/bin/application /application

CMD ["/application"]
