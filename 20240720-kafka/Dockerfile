FROM golang:1.22 AS buildStage
COPY go.mod go.sum /src_code/
WORKDIR /src_code
RUN go mod download
COPY . .
RUN go build -o /entrypoint .
RUN chmod +x /entrypoint


FROM debian:bookworm

COPY --from=buildStage /entrypoint /entrypoint

ENTRYPOINT ["/entrypoint"]
