package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

	//"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"

	"go_learning/srvDiscovy"
	pb "go_learning/srvDiscovy/protobuf/mail"
)

func main() {
	r := srvDiscovy.NewResolver([]string{
		"127.0.0.1:2379",
		"127.0.0.1:22379",
		"127.0.0.1:32379",
	}, "g.srv.mail")
	resolver.Register(r)

	ctx, _ := context.WithTimeout(context.Background(), 3000*time.Second)

	// https://github.com/grpc/grpc/blob/master/doc/naming.md
	/*g.srv.mail经测试，这个可以随便写，底层只是取scheme对应的Build对象*/
	addr := fmt.Sprintf("%s:///%s", r.Scheme(), "g.srv.mail")

	// grpc.WithBalancerName(roundrobin.Name),
	//指定初始化round_robin => balancer (后续可以自行定制balancer和 register、resolver 同样的方式)
	conn, err := grpc.DialContext(ctx, addr, grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`),
		grpc.WithBlock())

	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	c := pb.NewMailServiceClient(conn)

	resp, err := c.SendMail(context.TODO(), &pb.MailRequest{
		Mail: "qq@mail.com",
		Text: "test,test",
	})
	log.Print(resp)
}
