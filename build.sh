#!/usr/bin/env zsh
stop () {
    cd ../
    exit
}

rungo() {
    cd ../dist/api
    export TARGET="192.168.0.1/22"
    export FREQUENCY="1m"
    export PORT="6136"
    ./facemasq
    cd ../../
}

buildgo() {
    echo "Building API"
    cd ./api
    go build --tags "linux sqlite_foreign_keys=1" -o ../dist/api/facemasq . || stop
}

mkdir -p ./dist/api ./dist/ui
case $1 in
  api)
    buildgo
    rungo
    ;;
  container)
    docker build . -t=$2
    ;;
  ui)
    cd ./ui
    npm run build
    cd ../../
    ;;
  *)
    echo "Missing command"
    ;;
esac

