FROM golang

WORKDIR /go/src/github.com/pkliczewski/provider-pod
COPY . /go/src/github.com/pkliczewski/provider-pod

RUN go install

ENTRYPOINT /go/bin/provider-pod