package main

import (
	"encoding/json"
	"os"
	"os/exec"
	"strings"
)

type InstallConfig struct {
	FileLocation string
	Username     string
}

func runInstaller(d string) {

	elog.Info(1, "Polling for new installation trigger file in: "+d)

	entries, err := os.ReadDir(d)
	if err != nil {
		elog.Error(1, "Cannot read polling dir")
		return
	}
	elog.Info(1, "Read Polling Dir")

	for _, e := range entries {

		if strings.HasSuffix(e.Name(), "install.json") {

			b, err := os.ReadFile(d + "/install.json")
			if err != nil {
				elog.Error(1, "Error reading install.json : "+err.Error())
				return
			}
			elog.Info(1, "Read install.json")

			var cfg InstallConfig
			err = json.Unmarshal(b, &cfg)
			if err != nil {
				elog.Error(1, "Error unmarshalling install.json "+err.Error())
				return
			}
			elog.Info(1, "Unmarshalled install.json")

			process := exec.Command(cfg.FileLocation)
			elog.Info(1, "Running installer")
			err = process.Run()
			if err != nil {
				elog.Error(1, "Error running installer : "+err.Error())
				return
			}
			elog.Info(1, "Installer completed")

			err = os.Remove(d + "/install.json")
			if err != nil {
				elog.Error(1, "Error removing installation trigger file: "+err.Error())
				return
			}
			elog.Info(1, "Removed install.json")

			return
		}
	}
}
