## protobuf语法

syntax 声明采用proto3语法
```
protoc --go_out =. hello.proto
其中go_out参数告知protoc编译器去加载对应的protoc-gen-go工具，
然后通过该工具生成代码，生成代码放到当前目录。最后是一系列要处理的protobuf文件的列表。
```


```
$ protoc --go_out=plugins=grpc:. hello.proto

```