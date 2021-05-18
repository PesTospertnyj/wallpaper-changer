package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

func main() {
	var (
		sunrise       int    = 4
		sunriseImage  string = "Sunrise.jpg"
		day           int    = 12
		dayImage      string = "Day.jpg"
		evening       int    = 16
		eveningImage  string = "Evening.jpg"
		twilight      int    = 19
		twilightImage string = "Twilight.jpg"
		night         int    = 22
		nightImage    string = "Night.jpg"
	)
	for {
		currentHour := time.Now().Hour()
		switch {
		case currentHour > sunrise && currentHour < day:
			if !checkCurrentWallpaper(sunriseImage) {
				setWallpaper(sunriseImage)
			}

		case currentHour > day && currentHour < evening:
			if !checkCurrentWallpaper(dayImage) {
				setWallpaper(dayImage)
			}

		case currentHour > evening && currentHour < twilight:
			if !checkCurrentWallpaper(eveningImage) {
				setWallpaper(eveningImage)
			}

		case currentHour > twilight && currentHour < night:
			if !checkCurrentWallpaper(twilightImage) {
				setWallpaper(twilightImage)
			}

		case currentHour > night && currentHour < sunrise:
			if !checkCurrentWallpaper(nightImage) {
				setWallpaper(nightImage)
			}
		}

		time.Sleep(20 * time.Minute)
	}
}

func setWallpaper(name string) {
	path := "/home/artur/Изображения/Dynamic_Wallpapers/ZorinMountain"
	command := "./change-wallpaper.sh"
	args := fmt.Sprintf("%s/%s", path, name)
	cmd := exec.Command(command, args)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Command finished with error: %v", err)
	}

	log.Printf("wallpaper changed to %s", name)
}

func checkCurrentWallpaper(name string) bool {
	path := "/home/artur/Изображения/Dynamic_Wallpapers/ZorinMountain"
	possibleOut := fmt.Sprintf("%s/%s", path, name)
	command := "./check-current-wallpaper.sh"
	cmd := exec.Command(command)
	out, err := cmd.Output()
	if err != nil {
		log.Fatalf("get command output finished with error: %v", err)
	}

	return strings.Contains(string(out), possibleOut)
}
