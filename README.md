# go-plugin-std
go  官方插件使用 demo

## start

```shell
first:
cd ./demo
go build --buildmode=plugin -o demo.so

then:
go test engine_test.go

finally will print:
DemoFilter start
test
xx
DemoFilter end
```
