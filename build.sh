set -xe

# install packages and dependencies
go get github.com/aws/aws-sdk-go
go get -u github.com/gorilla/mux

# build command
go build -o bin/application application.go