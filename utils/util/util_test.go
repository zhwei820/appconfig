package util

import (
	"testing"
	"fmt"
)

func TestGetOutboundIP(t *testing.T) {
	res := GetOutboundIP()
	fmt.Printf("%v", res)
}
