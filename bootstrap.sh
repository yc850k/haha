#!/bin/sh
make clean

echo "downloading dependcies, it may take a few minutes..."

godep path > /dev/null 2>&1
if [ "$?" = 0 ]; then
    GOPATH=`godep path`:$GOPATH
    godep restore
else
    go get -u github.com/astaxie/beego/orm
    go get -u github.com/codegangsta/inject
    go get -u github.com/go-martini/martini
    go get -u github.com/go-sql-driver/mysql
    go get -u github.com/martini-contrib/encoder
fi

#generate model
go run generate.go --json=generate.json

make || exit $?
make gotest
