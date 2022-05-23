package tests

import (
	"context"
	"flag"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
	"google.golang.org/grpc"
	"learning-go/cmd/rpcx/controller"
	"learning-go/pkg/pb"
	"testing"
)

var (
	addr = flag.String("addr", "localhost:5000", "server address")
)

func TestSayHello(t *testing.T) {
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("connect failed: %v", err)
	}

	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	request := pb.HelloRequest{
		Name: "grpc",
	}
	reply, err := client.SayHello(context.Background(), &request)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	t.Logf("reply: %s", reply.Message)
}

func TestRPCXSayHello(t *testing.T) {
	flag.Parse()

	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	opt := client.DefaultOption
	opt.SerializeType = protocol.JSON

	xclient := client.NewXClient("GreeterServer", client.Failtry, client.RandomSelect, d, opt)
	defer xclient.Close()

	args := controller.HelloRequest{
		Name: "rpcx",
	}
	reply := &controller.HelloReply{}
	err := xclient.Call(context.Background(), "SayHello", args, reply)
	if err != nil {
		t.Fatalf("failed to call: %v", err)
	}

	t.Logf("reply: %s", reply.Message)
}
