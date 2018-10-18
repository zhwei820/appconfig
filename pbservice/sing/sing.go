package sing

import (
	. "github.com/zhwei820/appconfig/pb/appconfig/say"
)

func SingHello(in []byte) ([]byte, error) {
	args := &SayInput{}
	args.Unmarshal(in)
	output := &SayOutput{Url: "sing OK", Title: args.Query,}

	return output.Marshal()
}
