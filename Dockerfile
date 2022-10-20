FROM golang:1.19.2-alpine3.16 AS builder
WORKDIR /go/src
COPY server/go ./go
COPY server/main.go .
COPY server/go.mod .
COPY server/go.sum .


ARG VERSION="latest"
RUN CGO_ENABLED=0 go build -ldflags="-w -s -X main.version=${VERSION}" -o ./dist/jumia .

FROM scratch

# copy the binary that already built and all required files
COPY --from=builder /go/src/dist /
#COPY --from=builder /go/src/env/sample.config /env/config
EXPOSE 8085/tcp
ENTRYPOINT ["/jumia"]

