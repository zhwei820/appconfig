// Code generated by protoc-gen-hprose , DO NOT EDIT.
// source: hp.proto

package hprose



type SingService struct {
	Hello func(input []byte) (out  []byte, err error) `simple:"true"`
}

