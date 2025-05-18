FROM golang:1.24-alpine AS build

WORKDIR /src

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app cmd/server.go

FROM alpine

WORKDIR /server

RUN apk --no-cache add ca-certificates

COPY --from=build src/app .
COPY --from=build src/migrations/ ./migrations

EXPOSE 8080

ENTRYPOINT [ "./app" ]

