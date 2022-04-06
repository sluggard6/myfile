#!/bin/bash

cd ./assets
rm -rf ./dist
npm run build:prod
cd ..
statik -src=./assets/dist
