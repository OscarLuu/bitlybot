from golang:1.13-alpine as builder
copy ${PWD} /bitlybot/
workdir /bitlybot
run apk add --no-cache --update make \
  && go get -u github.com/mitchellh/gox \
  && make build

from scratch
copy --from=builder /bitlybot/bin/bitlybot_linux_amd64 .
entrypoint ["./bitlybot_linux_amd64"]
