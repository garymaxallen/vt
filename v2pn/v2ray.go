package v2pn

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	core "github.com/v2fly/v2ray-core/v4"
	"github.com/v2fly/v2ray-core/v4/common/cmdarg"
	_ "github.com/v2fly/v2ray-core/v4/main/distro/all"
)

func startV2Ray(filename string) (core.Server, error) {
	configFiles := cmdarg.Arg{filename}
	config, err := core.LoadConfig("json", configFiles[0], configFiles)
	if err != nil {
		return nil, err
	}
	server, err := core.New(config)
	if err != nil {
		return nil, err
	}
	return server, nil
}

func V2ray_start(dataDir string) {
	server, err := startV2Ray(dataDir + "/config.json")
	if err != nil {
		fmt.Println(err)
		// Configuration error. Exit with a special value to prevent systemd from restarting.
		os.Exit(23)
	}
	if err := server.Start(); err != nil {
		fmt.Println("Failed to start", err)
		os.Exit(-1)
	}
	defer server.Close()
	// Explicitly triggering GC to remove garbage from config loading.
	runtime.GC()
	{
		osSignals := make(chan os.Signal, 1)
		signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)
		<-osSignals
	}
}
