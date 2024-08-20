package main

import (
	"io/fs"
	"os"
)

var configDirPath string
var configDir []fs.DirEntry

func initConfig() {
	homeDir, _ := os.UserHomeDir()
	configDirPath = homeDir + "/.config/gofetch"
	var err error
	configDir, err = os.ReadDir(configDirPath)
	if err != nil {
		createDefaultConfig()
		configDir, err = os.ReadDir(configDirPath)
		if err != nil {
			panic(err)
		}
	}

}

func createDefaultConfig() {
	os.Mkdir(configDirPath, 0755)
	//os.WriteFile(configDirPath+"config.cfg", []byte(""), 0755)

	artDirPath := configDirPath + "/art"
	infoDirPath := configDirPath + "/info"
	os.Mkdir(artDirPath, 0755)
	os.Mkdir(infoDirPath, 0755)

	os.WriteFile(artDirPath+"/pikachu.txt", []byte(` \033[30m░░░░░░░░\033[33m▀████▀▄▄\033[30m░░░░░░░░░░░░░░\033[33m▄█
\033[30m░░░░░░░░░░\033[33m█▀\033[30m░░░░\033[33m▀▀▄▄▄▄▄\033[30m░░░░\033[33m▄▄▀▀█
\033[30m░░\033[33m▄\033[30m░░░░░░░░\033[33m█\033[30m░░░░░░░░░░\033[33m▀▀▀▀▄\033[30m░░\033[33m▄▀
\033[30m░\033[33m▄▀\033[30m░\033[33m▀▄\033[30m░░░░░░\033[33m▀▄\033[30m░░░░░░░░░░░░░░\033[33m▀▄▀
\033[33m▄▀\033[30m░░░░\033[33m█\033[30m░░░░░\033[33m█▀\033[30m░░░\033[33m▄█▀▄\033[30m░░░░░░\033[33m▄█
\033[33m▀▄\033[30m░░░░░\033[33m▀▄\033[30m░░\033[33m█\033[30m░░░░░\033[33m▀██▀\033[30m░░░░░\033[33m██▄█
\033[30m░\033[33m▀▄\033[30m░░░░\033[33m▄▀\033[30m░\033[33m█\033[30m░░░\033[33m▄██▄\033[30m░░░\033[33m▄\033[30m░░\033[33m▄\033[30m░░\033[33m▀▀\033[30m░\033[33m█
\033[30m░░\033[33m█\033[30m░░\033[33m▄▀\033[30m░░\033[33m█\033[30m░░░░\033[33m▀██▀\033[30m░░░░\033[33m▀▀\033[30m░\033[33m▀▀\033[30m░░\033[33m▄▀
\033[30m░\033[33m█\033[30m░░░\033[33m█\033[30m░░\033[33m█\033[30m░░░░░░\033[33m▄▄\033[30m░░░░░░░░░░░\033[33m▄▀ `), 0755)

	os.WriteFile(infoDirPath+"/minimal.txt", []byte(`\033[1m[*user*]@[*hostname*]
\033[1m\033[32mOS:\033[0m [*osName*]
\033[1m\033[33mKernel:\033[0m [*kernelVersion*]
\033[1m\033[34mWM:\033[0m [*desktopSession*] ([*desktopSessionType*])
\033[1m\033[31mCPU:\033[0m [*cpuModel*]
\033[1m\033[36mGPU:\033[0m [*gpuModel*]
\033[1m\033[32mMemory:\033[0m [*memUsed*] GB / [*memTotal*] GB ([*memUsedPercentColored*])
\033[1m\033[33mPackages:\033[0m [*packages*] [*flatpakPackages*] [*snaps*]
\033[1m\033[34mUptime:\033[0m [*uptime*]
`), 0755)

}
