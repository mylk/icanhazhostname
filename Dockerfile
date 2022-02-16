FROM alpine:latest

# configure Go
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH

RUN apk add --no-cache make musl-dev go && \
    mkdir -p ${GOPATH}/src ${GOPATH}/bin

ADD . $GOPATH/src
WORKDIR $GOPATH/src

CMD ["sh", "/go/src/run.sh"]

