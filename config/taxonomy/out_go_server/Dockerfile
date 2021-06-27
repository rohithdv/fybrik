FROM golang:1.10 AS build
WORKDIR /go/src
COPY openapiserver ./openapiserver
COPY main.go .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o openapiserver .

FROM scratch AS runtime
COPY --from=build /go/src/openapiserver ./
EXPOSE 8080/tcp
ENTRYPOINT ["./openapiserver"]
