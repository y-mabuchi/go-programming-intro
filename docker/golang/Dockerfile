FROM golang:latest

RUN mkdir /go/src/work

WORKDIR /go/src/work

ADD . .

RUN go get github.com/julienschmidt/httprouter \
    && go get github.com/oxequa/realize \
    && go get golang.org/x/net/http2 \
    && go get github.com/lib/pq

CMD [ "realize", "start", "--run" ]
