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
\033[1m[*osName*]
[*kernelVersion*]
[*desktopSession*] ([*desktopSessionType*])
[*cpuModel*]
[*gpuModel*]
[*memUsed*] GB / [*memTotal*] GB ([*memUsedPercentColored*])
[*packages*] [*flatpakPackages*] [*snaps*]
[*uptime*]`), 0755)

}
