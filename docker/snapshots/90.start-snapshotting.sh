#!/usr/bin/env bash

WD=/snapshots
DB=/lachesis/datadir

chmod 777 ${WD}/files
chmod 777 ${DB}

ln -s ${WD}/files /usr/share/nginx/html/snapshots
rm -f /usr/share/nginx/html/index.html
ln -s ${WD}/files/index.html /usr/share/nginx/html/index.html

sudo -u "#${NODE_UID}" nohup ${WD}/snapshotting.sh ${DB} &
