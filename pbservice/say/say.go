package say

import (
	. "github.com/zhwei820/appconfig/pb/appconfig/say"
)

func SayHello(in []byte) ([]byte, error) {
	args := &SayInput{}
	args.Unmarshal(in)
	output := &SayOutput{Url: "say OK", Title: args.Query,}

	return output.Marshal()
}
