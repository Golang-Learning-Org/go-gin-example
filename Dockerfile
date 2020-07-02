# FROM golang:latest

# ENV GOPROXY https://goproxy.cn,direct
# WORKDIR $GOPATH/src/github.com/EvanXzj/go-gin-example
# COPY . .

# RUN go build .

# EXPOSE 8000
# ENTRYPOINT [ "./go-gin-example" ]

### 记得挂载conf目录, link MySQL
FROM golang:alpine AS builder
ADD . /src
RUN cd /src && go build -o bin/go-gin-example main.go

FROM alpine:latest

WORKDIR /app
COPY ./conf /app/
COPY --from=builder /src/bin/go-gin-example /app/

EXPOSE 8000
ENTRYPOINT [ "/app/go-gin-example" ]

# FROM scratch

# WORKDIR $GOPATH/src/github.com/EDDYCJY/go-gin-example
# COPY . $GOPATH/src/github.com/EDDYCJY/go-gin-example

# EXPOSE 8000
# CMD ["./go-gin-example"]


