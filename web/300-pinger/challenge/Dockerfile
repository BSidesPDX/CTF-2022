FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/src/github.com/BSidesPDX/CTF-2022/web/300-no-idea/challenge
COPY . .
RUN go get -d -v
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /frontend

FROM busybox
COPY --from=builder ./frontend /frontend
RUN adduser -D sysop
RUN mkdir -p /home/sysop/.ssh
RUN chmod 700 /home/sysop/.ssh
COPY ./dropbear_priv /home/sysop/.ssh/id_rsa
RUN chmod 600 /home/sysop/.ssh/id_rsa
RUN chown -R sysop:sysop /home/sysop
COPY dropbearmulti /usr/bin/ssh
RUN chmod +s /bin/ping
USER 1000
ENTRYPOINT ["/frontend", "/frontend"]
