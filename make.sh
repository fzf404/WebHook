#!/bin/bash

mkdir -p webhooks/config webhooks/log webhooks/script

cp ./config/config.yaml webhooks/config
cp ./script/*.sh webhooks/script

go build -o webhooks/webhooks

tar -zcvf webhooks.tar.gz webhooks

rm -rf webhooks