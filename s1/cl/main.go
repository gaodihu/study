package main

import (
	"fmt"

	pb "github.com/gaodihu/study/s1/proto"
	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/consul"
	"golang.org/x/net/context"
)

func main() {
	service := micro.NewService(
		micro.Name("test-cl"),
		micro.Version("last"),
		micro.Metadata(map[string]string{"p": "p1", "p2": "p2"}),
	)

	g := pb.NewTestServiceClient("test-service", service.Client())
	r, e := g.Hello(context.Background(), &pb.TestRequest{Name: "pppp"})
	fmt.Println(r)
	fmt.Println(e)
	service.Init()
	service.Run()
}
