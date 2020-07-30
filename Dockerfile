FROM golang:alpine as build
WORKDIR /src
COPY . .
RUN go build -o server main.go

FROM alpine
COPY --from=build /src/server server
ENTRYPOINT ["./server"]
