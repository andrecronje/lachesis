# Snapshot service

allows to download the network events as single file to sync your node faster, without p2p load.


## Docker image "lachesis-snapshots"

is an example how to organize snapshot service.
There are lachesis node which exports events snapshots and nginx which shares snapshots for downloading.
See details at [snapshotting.sh](./snapshots/snapshotting.sh).

Build it by `make snapshots` from [Dockerfile.snapshots](./Dockerfile.snapshots).

Run it by
```sh
docker run --rm -d \
    -p 127.0.0.1:8080:80 \
    lachesis-snapshots:latest
```
and see the download page at [localhost:8080](http://127.0.0.1:8080).


## Deploy

example at [snapshots/deploy_example/](./snapshots/deploy_example).
