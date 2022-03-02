
FROM golang:rc-alpine3.13

RUN mkdir -p /go/src/autom
WORKDIR /go/src/autom
COPY go.mod /go/src/autom/
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download
COPY . /go/src/autom/
RUN go build -o server.bin main.go

EXPOSE 20153
CMD /go/src/autom/server.bin