package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func SetWallpaper(name string) {
	path := "/home/artur/Изображения/Dynamic_Wallpapers/ZorinMountain"
	command := "gsettings"
	file := fmt.Sprintf("\"file://%s/%s\"", path, name)
	cmd := exec.Command(command, "set", "org.gnome.desktop.background", "picture-uri", file)
	cmd.Stderr = os.Stderr
	out, err := cmd.Output()
	if err != nil {
		log.Fatalf("Command finished with error: %v, log : %v", err, out)
	}

	log.Printf("wallpaper changed to %s", name)
}
