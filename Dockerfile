FROM golang:alpine AS build-env

RUN apk add --no-cache git

ADD . /src

RUN cd /src && go mod tidy && CGO_ENABLED=0 go build -o myapp cmd/app/main.go

# final stage
FROM alpine

RUN apk add --no-cache ca-certificates

COPY --from=build-env /src/myapp /app/

ENTRYPOINT ["/app/myapp"]
