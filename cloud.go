package main

import (
	"errors"
	"io"
	"net/http"
	"time"
)

var client = &http.Client{
	Timeout: 10 * time.Second,
	Transport: &http.Transport{
		MaxIdleConns:        2,
		MaxIdleConnsPerHost: 2,
		IdleConnTimeout:     30 * time.Second,
	},
}

var (
	Header = []byte("# IP")
	RealIP = []byte("set_real_ip_from ")
	Next   = []byte(";\n")

	Footer = []byte("real_ip_header CF-Connecting-IP;\nreal_ip_recursive on;\n")
)

func LoadIPs(header, url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "go/flare")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var (
		b     byte
		start int
		conf  = make([]byte, 0, len(body)*2+20)
	)

	conf = append(conf, Header...)
	conf = append(conf, header...)
	conf = append(conf, '\n')

	for i := 0; i < len(body); i++ {
		b = body[i]

		if b == '\n' || b == '\r' || b == ' ' {
			if i-start >= 7 {
				conf = append(conf, RealIP...)
				conf = append(conf, body[start:i]...)
				conf = append(conf, Next...)
			}

			for i < len(body) && (body[i] == '\n' || body[i] == '\r' || body[i] == ' ') {
				i++
			}

			start = i

			i += 6
		}
	}

	if start < len(body) && len(body)-start >= 7 {
		conf = append(conf, RealIP...)
		conf = append(conf, body[start:]...)
		conf = append(conf, Next...)
	}

	conf = append(conf, '\n')

	return conf, nil
}
