FROM golang:alpine AS builder

# RUN apt-get update -y && apt install -y clang
# RUN ln -s /usr/include/asm-generic /usr/include/asm
WORKDIR /app 

COPY . .
# RUN sh build.sh

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o myfile .

FROM alpine:latest

WORKDIR /myfile

COPY --from=builder /app/myfile .
COPY .conf/*.yaml ./conf/

ENTRYPOINT [ "./myfile", "start"]