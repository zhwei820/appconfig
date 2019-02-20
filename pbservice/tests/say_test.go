package tests

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/zhwei820/natsmicro/pb/natsmicro/say"
	"github.com/zhwei820/natsmicro/pbservice"
	_ "github.com/zhwei820/natsmicro/utils/gotests"
	"testing"
	"time"
)

func TestSay(t *testing.T) {
	for ii := 0; ii < 10000; ii++ {
		testfunc()
	}

}

func BenchmarkSay(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		testfunc()
	}
}

func testfunc() {

	out := say.SayOutput{}
	in := say.SayInput{Query: "testttt", PageNumber: 1, ResultPerPage: 10}
	inb, _ := in.Marshal()
	err := out.Unmarshal(NatsReq("sayhello", inb)) // 请求接口
	if err != nil {
		log.Error().Msg(fmt.Sprintf("Unmarshal SayOutput failed: %v", err.Error()))
	}

}

func NatsReq(subject string, inb []byte) ([]byte) {
	msg, err := pbservice.NatsConn.Request(subject, inb, 5000*time.Millisecond)
	if err != nil {
		log.Error().Msg(fmt.Sprintf("Request failed: %v", err.Error()))
		return []byte{}
	}
	log.Info().Str("subject", subject).Str("inb", string(inb)).Str("outb", string(msg.Data))
	return msg.Data
}
