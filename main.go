package main

import (
	"log"
	"time"

	"wallpaper-changer/cmd"
)

func main() {
	/* TODO: придумать как реализовать для остальных сред рабочего стола
	+ добавить рассчеты по алгоритмам https://edwilliams.org/sunrise_sunset_algorithm.htm
	*/
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
		case currentHour >= sunrise && currentHour < day:
			if !cmd.CheckCurrentWallpaper(sunriseImage) {
				cmd.SetWallpaper(sunriseImage)
				break
			}

			log.Println("wallpaper not changed")

		case currentHour >= day && currentHour < evening:
			if !cmd.CheckCurrentWallpaper(dayImage) {
				cmd.SetWallpaper(dayImage)
				break
			}

			log.Println("wallpaper not changed")

		case currentHour >= evening && currentHour < twilight:
			if !cmd.CheckCurrentWallpaper(eveningImage) {
				cmd.SetWallpaper(eveningImage)
				break
			}

			log.Println("wallpaper not changed")

		case currentHour >= twilight && currentHour < night:
			if !cmd.CheckCurrentWallpaper(twilightImage) {
				cmd.SetWallpaper(twilightImage)
				break
			}

			log.Println("wallpaper not changed")

		case currentHour >= night:
			if !cmd.CheckCurrentWallpaper(nightImage) {
				cmd.SetWallpaper(nightImage)
				break
			}

			log.Println("wallpaper not changed")

		}

		time.Sleep(1 * time.Minute)
	}
}
