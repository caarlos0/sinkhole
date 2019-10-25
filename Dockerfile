FROM golang:1.13-alpine as build
WORKDIR /src
COPY . .
ENV CGO_ENABLED=0
RUN go build -o server main.go

FROM alpine
COPY --from=build /src/server server
ENTRYPOINT ["./server"]
