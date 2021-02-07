FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get github.com/T-times/Interview-Prep/app
RUN cd /build && git clone https://github.com/T-times/go-test.git

RUN cd /build/Interview-Prep/app && go build

EXPOSE 9999

ENTRYPOINT [ "/build/Interview-Prep/app/app" ]