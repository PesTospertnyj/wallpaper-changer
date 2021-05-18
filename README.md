sudo ln -s /home/artur/projects/golang/wallpaper-changer/wallpaper_changer.service /etc/systemd/system/wallpaper_changer.service

sudo systemctl daemon-reload

sudo systemctl enable wallpaper_changer.service

sudo rm /etc/systemd/system/wallpaper_changer.service

sudo systemctl start wallpaper_changer

sudo systemctl status wallpaper_changer

journalctl -f -u wallpaper_changer