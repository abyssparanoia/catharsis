FROM golang:1.13.0-alpine3.10 AS builder

ARG SERVICE_NAME=authentication

WORKDIR /go/src/github.com/abyssparanoia/catharsis/

ENV CGO_ENABLED=0

COPY . .

RUN go install -v -tags netgo -ldflags '-extldflags "-static"' ./cmd/authentication/


FROM alpine AS server

RUN apk add ca-certificates
COPY --from=builder /go/bin/authentication /bin/server

WORKDIR /go/src/github.com/abyssparanoia/catharsis/

ENV PORT 50051
EXPOSE 50051

ENTRYPOINT ["/bin/server"]