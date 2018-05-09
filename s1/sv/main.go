package main

import (
	"fmt"

	pb "github.com/gaodihu/study/s1/proto"
	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/consul"
	"golang.org/x/net/context"
)

type G struct {
}

func (g *G) Hello(c context.Context, rqt *pb.TestRequest, rsp *pb.TestResponse) error {
	rsp.Msg = rqt.Name + " product"
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("test-service"),
		micro.Version("last"),
	)

	service.Init()
	pb.RegisterTestServiceHandler(service.Server(), new(G))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
