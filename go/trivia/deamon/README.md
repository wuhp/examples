# Build
export GOPATH=$PWD
go get github.com/VividCortex/godaemon     // 3d9f6e0b234fe7d17448b345b2e14ac05814a758
go install main

# Start
./bin/main
./bin/main -o /tmp/123.log
