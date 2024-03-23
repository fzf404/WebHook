#!/bin/bash

mkdir -p webhook/config webhook/log webhook/script

cp ./config/config.yaml webhook/config
cp ./script/*.sh webhook/script

go build -o webhook/webhook

tar -zcvf webhook.tar.gz webhook

rm -rf webhook
