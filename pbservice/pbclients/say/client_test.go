package say

import (
	"log"
	"testing"
)

func TestCallSayHello(t *testing.T) {
	out, err := SayRpc(SayClientRpc, "SayHello", []byte("tessssssss"))
	if err != nil {
		log.Fatal("error:", err.Error())
	}
	log.Print(string(out))
}
