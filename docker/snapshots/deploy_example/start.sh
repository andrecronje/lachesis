#!/usr/bin/env bash
cd $(dirname $0)

mkdir -p datadir
mkdir -p snapshots

docker run --rm -d \
	--name lachesis-snapshots \
	-e NODE_UID=$(id -u) \
	-v ${PWD}/datadir:/lachesis/datadir \
	-v ${PWD}/snapshots:/snapshots/files \
	-p 127.0.0.1:8080:80 \
	lachesis-snapshots:latest
