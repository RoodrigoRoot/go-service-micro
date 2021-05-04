package main

import
(
	"context"
	"log"
	"fmt"
	pb "github.com/roodrigoroot/go-service-micro/proto"
	"github.com/asim/go-micro/v3"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *pb.Request, rsp *pb.Response)  error {
	rsp.Msg="Hola " + req.Name
	fmt.Print(rsp)
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("helloworld"),
		micro.Address(":8081"),
	
	)

	service.Init()

	pb.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
