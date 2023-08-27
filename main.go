package main

//go:generate go run github.com/99designs/gqlgen generate

//go:generate protoc -I /usr/local/include/ -I $GOPATH/src/ -I ./schema/src/proto --go-grpc_out=./schema/surveypb/ --go_out=./schema/surveypb/ ./schema/src/proto/survey.proto --go-grpc_opt=paths=source_relative --go_opt=paths=source_relative

//go:generate go run ./schema/src/proto/tagparser/tagparser.go ./schema/surveypb/
