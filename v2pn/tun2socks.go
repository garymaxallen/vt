package v2pn

import (
	"os"
	"os/signal"
	"strconv"

	// "strconv"
	"syscall"

	"github.com/xjasonlyu/tun2socks/v2/engine"
	"github.com/xjasonlyu/tun2socks/v2/log"
)

func Tun2socks_start(socks5Addr string, fd int, externalFilesDir string) {
	// change this file
	// D:\tmp\vt\tun2socks\log\log.go
	// func init() {
	// 	// logrus.SetOutput(os.Stdout)
	// 	// logrus.SetLevel(logrus.DebugLevel)
	// }

	// func SetLogPath(logPath string) {
	// 	f, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// 	if err != nil {
	// 		log.Fatalf("error opening file: %v", err)
	// 	}
	// 	// defer f.Close()
	// 	logrus.SetLevel(logrus.DebugLevel)
	// 	logrus.SetOutput(f)
	// 	logrus.Println("This is a t2s test log entry")
	// }

	log.SetLogPath(externalFilesDir + "/t2sLog.txt")
	key := new(engine.Key)
	key = &engine.Key{
		Proxy:  socks5Addr,
		Device: "fd://" + strconv.Itoa(fd),
		//Device:   "wintun",
		LogLevel: "debug"}
	engine.Insert(key)
	engine.Start()
	defer engine.Stop()
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
}
