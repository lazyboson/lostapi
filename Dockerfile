#Build stage
FROM golang:1.18.4-alpine3.16 AS builder
WORKDIR /build
COPY . .
RUN go build -o restService cmd/restService.go

#Run stage
FROM alpine:3.16
WORKDIR /build
COPY --from=builder /build/restService .
COPY cmd/rest.env .
EXPOSE 3232
CMD ["/entrypoint.sh"]
