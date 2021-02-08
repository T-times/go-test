FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get github.com/T-times/go-test/app
RUN cd /build && git clone https://github.com/T-times/go-test.git

RUN cd /build/go-test/app && go build

EXPOSE 9999

ENTRYPOINT [ "/build/go-test/app/app" ]