FROM golang:1.17-alpine3.15 as builder

RUN apk --no-cache add build-base

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.io,direct

# RUN apt-get update -y && apt install -y clang
# RUN ln -s /usr/include/asm-generic /usr/include/asm
WORKDIR /app 

COPY . .
RUN rm application.json

# RUN CC=clang CXX=clang++ GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 go build .
# RUN GOARCH=amd64 CGO_ENABLED=1 go build -ldflags="-w -s" -o .
# RUN CGO_ENABLED=1 GOARCH=amd64 go build -ldflags="--extldflags" .
RUN go build 


FROM alpine:latest

WORKDIR /myfile

COPY --from=builder /app/myfile .

ENTRYPOINT [ "./myfile", "start" ]