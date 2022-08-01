FROM golang:latest

RUN mkdir /build 
WORKDIR /build

RUN export GO111MODULE=on
RUN go install github.com/dm-uporov/HelloGo@latest
RUN cd / build && git clone https://github.com/dm-uporov/HelloGo.git

RUN cd /build/HelloGo && go build

EXPOSE 9090

ENTRYPOINT [ "/build/HelloGo/main" ]
