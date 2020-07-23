# compile the service
FROM golang:1.14 as builder
COPY . /cohack-golang
WORKDIR /cohack-golang
ENV GO111MODULE=on
# for using go module in china
ENV GOPROXY=https://goproxy.cn,direct
RUN CGO_ENABLED=0 GOOS=linux go build -o todo-service

# use scratch for minimal image size
FROM scratch
WORKDIR /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /cohack-golang/todo-service .
ENTRYPOINT ["./todo-service"]