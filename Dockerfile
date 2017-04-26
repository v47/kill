FROM golang:1.8

ADD ./ /src
WORKDIR /src

RUN cd plugin && go build -buildmode=plugin -o myplugin.so plugin.go
RUN go build -o plugin_service

EXPOSE 8181:8181 6060:6060

ENTRYPOINT ["/src/plugin_service"]