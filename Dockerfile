FROM golang:1.16-alpine AS build

ENV GO111MODULE on
ENV CGO_ENABLED 0
ENV GOPROXY="https://goproxy.cn,direct"

WORKDIR /src/
# ADD go.mod .
# ADD go.sum .
# RUN go mod download
COPY . .
# RUN go build -ldflags="-s -w" -o /bin/zvos-edge-agent
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -mod=vendor -o /bin/zvos-edge-command-control main.go

FROM harbor.zvos.zoomlion.com/zvos-alpine/alpine-tz:v1.0

## replace the mirror
#RUN sed -i 's!https://dl-cdn.alpinelinux.org/!https://mirrors.ustc.edu.cn/!g' /etc/apk/repositories
#
## timezone
#RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata \
#    && ln -snf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
#    && echo "Asia/Shanghai" > /etc/timezone
#ENV TZ Asia/Shanghai

COPY --from=build /bin/zvos-edge-command-control /bin/zvos-edge-command-control
ENTRYPOINT ["/bin/zvos-edge-command-control"]
