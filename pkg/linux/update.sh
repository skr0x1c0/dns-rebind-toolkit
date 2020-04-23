#!/bin/bash

sudo cp pointerpw /usr/bin
sudo systemctl stop pointerpw
sudo cp pointerpw.service /lib/systemd/system
sudo systemctl daemon-reload
sudo systemctl restart pointerpw

sudo cp api.pointer.pw /etc/nginx/sites-available
sudo nginx -s reload