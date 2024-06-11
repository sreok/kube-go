FROM golang:1.22.0
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY="https://goproxy.cn,direct"
	
COPY ./build/main .kube-go
EXPOSE 8080
ENTRYPOINT ["./kube-go"]