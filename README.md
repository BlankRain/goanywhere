# GoAnywhere
some code in go
### You Can Go AnyWhere. So Do I.

Before you run `go run ./eqrcode.go`

You need a csv file  like this 

`10003,060102001,项目名称,安装位置`

The pattern is: $orgId,$ecode,$orgName,$ePosition

then you need create a folder  ./$org

then run 

`go run ./eqrcode.go pic` to create qrcode file

and 

`go run ./eqrcode.go somethingelse` to create a html file
