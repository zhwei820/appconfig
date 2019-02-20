package say

import (
	. "github.com/zhwei820/natsmicro/pb/natsmicro/say"
)

func SayHello(in []byte) ([]byte, error) {
	args := &SayInput{}
	args.Unmarshal(in)
	output := &SayOutput{Url: "yyyyyyyyyyysay OKxxxxxxxxxxxxxxxxxxx", Title: args.Query,}

	return output.Marshal()
}
