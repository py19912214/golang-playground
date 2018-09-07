package main
 
import (
	"context"
	"fmt"
	micro "github.com/micro/go-micro"
	proto "github.com/py19912214/golang-playground/microserver/proto"
	version "github.com/py19912214/golang-playground/go-micro-restful-demo/proto/go"
)
 
 
func main() {
	microserverDemo();
	goMicroRestfulDemo();
}
//连接自己写的微服务
func microserverDemo(){
	// Create a new service. Optionally include some options here.
	service := micro.NewService(micro.Name("greeter.client"))
	service.Init()
 
	// Create new greeter client
	greeter := proto.NewGreeterService("greeter", service.Client())
 
	// Call the greeter
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "panpan"})
	if err != nil {
		fmt.Println(err)
	}
 
	// Print response
	fmt.Println(rsp.Greeting)
}
//连接micro的微服务
func goMicroRestfulDemo(){
	// Create a new service. Optionally include some options here.
	service := micro.NewService(micro.Name("go.micro.restful.demo.client"))
	service.Init()
	// Create new greeter client
	versionReq := version.NewVersionService("go.micro.restful.demo", service.Client())
 
	// Call the greeter
	rsp, err := versionReq.Version(context.TODO(), &version.Request{Name: "yueyue"})
	if err != nil {
		fmt.Println(err)
	}
 
	// Print response
	fmt.Println(rsp.Msg)
}