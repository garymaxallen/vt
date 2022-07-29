package main

import "vt/v2pn"

func main() {
	// v2pn.V2ray("D:\\tmp\\vt\\v2ray-core-4.45.0\\config.json")
	//v2pn.V2ray_start("/root/vt/config.json")
	//v2pn.V2ray_start("D:\\tmp\\tmp2\\vt\\config.json")
	// v2pn.Tun2socks_start("socks5://127.0.0.1:2080", 3, "D:\\tmp\\tmp2\\t2s.log")
	//V2ray_start("/root/vt/config.json")
	v2pn.Http_server_start("D:\\Downloads", "D:\\")
}
