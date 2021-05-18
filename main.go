package main

import (
	"time"

	"wallpaper-changer/cmd"
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
			if !cmd.CheckCurrentWallpaper(sunriseImage) {
				cmd.SetWallpaper(sunriseImage)
			}

		case currentHour > day && currentHour < evening:
			if !cmd.CheckCurrentWallpaper(dayImage) {
				cmd.SetWallpaper(dayImage)
			}

		case currentHour > evening && currentHour < twilight:
			if !cmd.CheckCurrentWallpaper(eveningImage) {
				cmd.SetWallpaper(eveningImage)
			}

		case currentHour > twilight && currentHour < night:
			if !cmd.CheckCurrentWallpaper(twilightImage) {
				cmd.SetWallpaper(twilightImage)
			}

		case currentHour > night && currentHour < sunrise:
			if !cmd.CheckCurrentWallpaper(nightImage) {
				cmd.SetWallpaper(nightImage)
			}
		}

		time.Sleep(20 * time.Minute)
	}
}
