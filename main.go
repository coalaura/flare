package main

import (
	"fmt"
	"os"
	"sync"
)

const ConfigPath = "/etc/nginx/flare.conf"

func main() {
	var (
		conf = make([]byte, 0, 1024)

		wg sync.WaitGroup

		v4 []byte
		v6 []byte

		errV4 error
		errV6 error
	)

	wg.Add(2)

	go func() {
		defer wg.Done()

		v4, errV4 = LoadIPs("v4", "https://www.cloudflare.com/ips-v4/")
	}()

	go func() {
		defer wg.Done()

		v6, errV6 = LoadIPs("v6", "https://www.cloudflare.com/ips-v6/")
	}()

	wg.Wait()

	fail(errV4, errV6)

	conf = append(conf, v4...)
	conf = append(conf, v6...)
	conf = append(conf, Footer...)

	if !HasConfigChanged(conf) {
		fmt.Println("config has not changed")

		return
	}

	if len(os.Args) > 1 && len(os.Args[1]) >= 5 && os.Args[1][:5] == "--dry" {
		os.Stdout.Write(conf)

		return
	}

	fail(WriteToConfig(conf))
	fail(ReloadNginx())

	fmt.Println("config updated")
}

func fail(errs ...error) {
	for _, err := range errs {
		if err == nil {
			continue
		}

		fmt.Printf("error: %v\n", err)

		os.Exit(1)
	}
}
