export GOPATH=$PWD

go get github.com/gorilla/mux

go install
./bin/hello
