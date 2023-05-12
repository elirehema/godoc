FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
# RUN go install github.com/elirehema/tchsam
RUN cd /build && git clone https://github.com/elirehema/godoc.git

RUN cd /build/godoc && go build

EXPOSE 8080
ENTRYPOINT ["/build/godoc/main"]