# terminal 1
go run receive.go "a.b" "c.d"

# terminal 2
go run emit.go "a.b" "hello a & b"
go run emit.go "c.d" "hello c & d"
