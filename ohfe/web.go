package main

import "log"

func main() {
	ohweb, err := web.NewWebApp(gw_addr, svc_addr, clients)
	if err != nil {
		log.Fatal(err)
	}
	ohweb.Start()
}
