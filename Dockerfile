FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on 
RUN go get github.com/aydinsercan/GolangRestApi/main
RUN cd /build && git clone https://github.com/aydinsercan/GolangRestApi.git

RUN cd /build/GolangRestApi/main && go build

EXPOSE 8888

ENTRYPOINT [ "/build/GolangRestApi/main" ]