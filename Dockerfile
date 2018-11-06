FROM golang:1.11.1 as build
WORKDIR /go/src/github.com/jessicagreben/health-check
COPY . .

RUN go get -d -v ./... && \
    go install -v ./...


FROM gcr.io/distroless/base
COPY --from=build /go/bin/health-check /
CMD ["/health-check"]