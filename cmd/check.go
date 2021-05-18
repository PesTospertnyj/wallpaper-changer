package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func CheckCurrentWallpaper(name string) bool {
	path := "/home/artur/Изображения/Dynamic_Wallpapers/ZorinMountain"
	possibleOut := fmt.Sprintf("%s/%s", path, name)
	command := "gsettings"
	cmd := exec.Command(command, "get", "org.gnome.desktop.background", "picture-uri")
	out, err := cmd.Output()
	if err != nil {
		log.Fatalf("get command output finished with error: %v", err)
	}

	return strings.Contains(string(out), possibleOut)
}
