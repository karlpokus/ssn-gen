#!/bin/bash

VERSION=`cat version`

docker run -d -p 9012:9012 --restart always --name ssns pokus2000/ssns:$VERSION
