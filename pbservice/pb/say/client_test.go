// Code generated by protoc-gen-hprose , DO NOT EDIT.
// source: say.proto

package say

import (
	"log"
	"testing"
)



func TestCallSayHello(t *testing.T) {
	out, err := SayRpc(SayClientRpc, "SayHello", SayHelloInputBytes)
	if err != nil {
		log.Fatal("error:", err.Error())
	}
	log.Print(string(out))
}



