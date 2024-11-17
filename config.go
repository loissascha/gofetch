package main

import (
	_ "embed"
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

//go:embed defaults/art/pikachu.txt
var pikachu string

//go:embed defaults/info/minimal.txt
var minimalConfig string

func createDefaultConfig() {
	os.Mkdir(configDirPath, 0755)
	//os.WriteFile(configDirPath+"config.cfg", []byte(""), 0755)

	artDirPath := configDirPath + "/art"
	infoDirPath := configDirPath + "/info"
	os.Mkdir(artDirPath, 0755)
	os.Mkdir(infoDirPath, 0755)

	os.WriteFile(artDirPath+"/pikachu.txt", []byte(pikachu), 0755)
	os.WriteFile(infoDirPath+"/minimal.txt", []byte(minimalConfig), 0755)
}
