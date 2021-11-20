#! /bin/bash

trap "rm server;kill 0" EXIT

go build   -o server mutitest.go

./server -port=8001 &
./server -port=8002 &
./server -port=8003 -api=1 &

sleep 2

curl "http://localhost:9999/api?key=Tom" &
curl "http://localhost:9999/api?key=Tom" &
curl "http://localhost:9999/api?key=Tom" &

sleep 3