package sing

import (
	"github.com/zhwei820/appconfig/pbservice/pb/appconfig/sing"
)
func SingHello(in []byte) ([]byte, error) {
	args := &sing.SingInput{}
	args.Unmarshal(in)
	output := &sing.SingOutput{Url: "sing OK", Title: args.Query,}

	return output.Marshal()
}
