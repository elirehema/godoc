FROM golang:1.19

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
# RUN go install github.com/elirehema/tchsam
RUN cd /build && git clone https://github.com/elirehema/godoc.git
COPY .env .
RUN cd /build/godoc && CGO_ENABLED=0 go build -o /main
RUN ls -al godoc/
EXPOSE 8080
ENTRYPOINT ["/main"]