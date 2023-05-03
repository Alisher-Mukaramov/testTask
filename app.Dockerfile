FROM golang:1.19-buster as builder
WORKDIR /build
COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 go build -a -o ./.bin/app ./cmd/main.go


FROM alpine
RUN apk --no-cache add ca-certificates
COPY --from=builder /build/.bin/app /bin/app

ENTRYPOINT ["/bin/app"]
