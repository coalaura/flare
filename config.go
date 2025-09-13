package main

import (
	"bytes"
	"os"
	"os/exec"
)

func HasConfigChanged(conf []byte) bool {
	existing, err := os.ReadFile(ConfigPath)
	if err != nil {
		if os.IsNotExist(err) {
			return true
		}

		fail(err)
	}

	return !bytes.Equal(existing, conf)
}

func WriteToConfig(conf []byte) error {
	return os.WriteFile(ConfigPath, conf, 0644)
}

func ReloadNginx() error {
	if err := exec.Command("systemctl", "reload", "nginx").Run(); err == nil {
		return nil
	}

	return exec.Command("nginx", "-s", "reload").Run()
}
