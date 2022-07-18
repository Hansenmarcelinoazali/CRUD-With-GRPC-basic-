How to Crete Proto:
protoc \
      -I. \
      -I/usr/local/include \
      -I${GOPATH}/src \
      -I${GOPATH}/src/latihan/proto \
      -I${GOPATH}/src/latihan/proto/lib \
      --go_out=plugins=grpc:$GOPATH/src \
      --grpc-gateway_out=logtostderr=true:$GOPATH/src \
      v1/books/books.proto

How to run :
go run main.go
