FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/challenge/
COPY . .
RUN go get -d -v
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /challenge

FROM scratch
COPY --from=builder ./challenge /challenge
USER 65535
ENTRYPOINT ["/challenge", "/challenge"]