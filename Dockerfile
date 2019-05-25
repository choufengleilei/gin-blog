
#####################################
#   STEP 1 build executable binary  #
#####################################
#FROM choufengleilei/golang-build:v1 AS build
#
#WORKDIR /go/src/gin-blog
#
#COPY . .
#
#RUN go mod download
#
#RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o gin-blog

#alpine 测试环境
#scratch
FROM alpine

WORKDIR $GOPATH/src/gin-blog

COPY gin-blog .

COPY docs ./docs

COPY conf ./conf

COPY runtime/logs ./runtime/logs

EXPOSE 8000

ENTRYPOINT ["./gin-blog"]

