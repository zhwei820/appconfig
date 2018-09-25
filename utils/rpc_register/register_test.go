package rpc_register

import (
	_ "github.com/zhwei820/appconfig/utils/gotests"
	"testing"
	"os"
	"os/signal"
	"syscall"
)

func TestDiscoveryRegister(t *testing.T) {
	DiscoveryRegister()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM) // 程序退出清理资源
	<-sigs

	Cancel()
	println("exit!")
}
