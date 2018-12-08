package say

import (
	"github.com/zhwei820/appconfig/pbservice/pb/appconfig/say"
)

func SayHello(in []byte) ([]byte, error) {
	args := &say.SayInput{}
	args.Unmarshal(in)
	output := &say.SayOutput{Url: "say OK", Title: args.Query,}

	return output.Marshal()
}
