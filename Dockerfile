FROM golang:1.22.0
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY="https://goproxy.cn,direct"
WORKDIR /opt/kube-go
COPY . .
RUN go mod tidy
RUN go build -o app ./main
EXPOSE 8080
ENTRYPOINT ["/opt/kube-go/app"]