package pbservice

import (
	. "github.com/zhwei820/appconfig/pb/appconfig"
)

func Say(in []byte) ([]byte, error) {
	args := &SayInput{}
	args.Unmarshal(in)
	output := &SayOutput{Url: "OK", Title: args.Query,}

	return output.Marshal()
}
