#!/bin/sh
cat /etc/issue
go version
git version
pwd
ls
service mysql start
echo get code
go get -u -v github.com/jchprj/GeoOrderTest
cd $GOPATH/src/github.com/jchprj/GeoOrderTest/docs
sh gen.sh
cd ../
echo start test
go test ./...

echo start server
$GOPATH/bin/GeoOrderTest --config /root/config.yml &
swagger serve --no-open -p 8081 $GOPATH/src/github.com/jchprj/GeoOrderTest/docs/swagger.json &
bash