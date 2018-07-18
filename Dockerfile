FROM golang:1.9.4

RUN go get -u github.com/golang/protobuf/proto && \
    go get firebase.google.com/go   && \ 
    go get firebase.google.com/go/messaging  && \ 
    go get google.golang.org/api/option  && \ 
    go get github.com/satori/go.uuid && \
    go get google.golang.org/grpc  && \
    go get golang.org/x/net/context

RUN mkdir $GOPATH/src/github.com && \
    mkdir $GOPATH/src/github.com/Multy-io && \
    cd $GOPATH/src/github.com/Multy-io && \ 
    git clone https://github.com/Multy-io/Multy-back.git && \ 
    cd Multy-back && \ 
    git checkout release_1.1.1-prod && \  
    git pull origin release_1.1.1-prod

RUN cd $GOPATH/src/github.com/golang/protobuf && \
    make all

RUN apt-get update && \
    apt-get install -y protobuf-compiler

RUN cd $GOPATH/src/github.com/Multy-io/Multy-back && \ 
    rm -r $GOPATH/src/github.com/Multy-io/Multy-back/vendor/github.com/golang/protobuf && \
    make build

WORKDIR /go/src/github.com/Multy-io/Multy-back/cmd

RUN echo "VERSION 02"

ENTRYPOINT $GOPATH/src/github.com/Multy-io/Multy-back/cmd/multy
