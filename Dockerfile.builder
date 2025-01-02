FROM registry.cn-hangzhou.aliyuncs.com/117503445-mirror/sync:linux.amd64.docker.io.library.golang.1.23.4-alpine

RUN go install github.com/twitchtv/twirp/protoc-gen-twirp@latest && go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories

RUN apk add --no-cache protobuf

ENV CGO_ENABLED=0

RUN go env -w GOPROXY=https://goproxy.cn,direct

WORKDIR /workspace

ENTRYPOINT [ "/workspace/scripts/build-in-docker.sh" ]