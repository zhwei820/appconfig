// Code generated by protoc-gen-hprose , DO NOT EDIT.
// source: sing.proto

package sing



type SingService struct {
	SingHello func(input []byte) (out  []byte, err error) `simple:"true"`
}

