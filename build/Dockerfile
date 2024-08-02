FROM golang:1.11.1-alpine3.8

RUN apk add git make nodejs npm python g++ curl && \
    mkdir -p $GOPATH/src/MarsOccupied/plasma-core && \
    cd $GOPATH/src/MarsOccupied/plasma-core && \
    npm i -g truffle && \
    go get -u ethereum/go-ethereum && \
    cd $GOPATH/src/ethereum/go-ethereum && \
    go install ./cmd/abigen

CMD ["echo", "Done."]
