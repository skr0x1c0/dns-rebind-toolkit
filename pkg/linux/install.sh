#!/bin/bash

sudo cp pointerpw /usr/bin
sudo cp pointerpw.service /lib/systemd/system
sudo systemctl enable pointerpw
sudo systemctl start pointerpw

sudo apt update
sudo apt install -y nginx
sudo cp api.pointer.pw /etc/nginx/sites-available
sudo ln -s /etc/nginx/sites-available/api.pointer.pw /etc/nginx/sites-enabled/api.pointer.pw
sudo nginx -s reload