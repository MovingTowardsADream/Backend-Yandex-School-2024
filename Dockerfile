FROM golang:1.22

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o yandex-lavka ./cmd/main.go

CMD ["./yandex-lavka"]