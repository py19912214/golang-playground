1.先安装consul 然后启动
consul agent -dev
访问：http://localhost:8500/
2.安装protoc
（1.github上下载一个cpp包：https://github.com/google/protobuf/releases make make install安装即可
（2.protoc-gen-go
	go get -u github.com/golang/protobuf/protoc-gen-go
（3.安装protoc-gen-micro
	go get github.com/micro/protoc-gen-micro

3.进入protoc目录
protoc --proto_path=./ --micro_out=. --go_out=. pubsub.proto
