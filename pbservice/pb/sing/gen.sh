GOGO_ROOT=${GOPATH}/src/github.com/gogo/protobuf

protoc -I.:${GOPATH}/src  --hprose_out=.  --gogofaster_out=. sing.proto
