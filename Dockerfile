FROM sluggard/myfilebase:latest as builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

# RUN apt-get update -y && apt install -y clang
# RUN ln -s /usr/include/asm-generic /usr/include/asm
WORKDIR /app 

COPY . .

RUN go build 

FROM alpine:latest

WORKDIR /myfile

COPY --from=builder /app/myfile .
#COPY conf/application.json /myfile/conf/application

ENTRYPOINT [ "./myfile" ]