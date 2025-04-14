FROM golang:1.22 as builder

WORKDIR /app
COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bootstrap ./cmd/ioManager/main.go

FROM scratch
COPY --from=builder /app/bootstrap /bootstrap

ENTRYPOINT ["/bootstrap"]