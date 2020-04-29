package rpc

import (
	"context"
	"fmt"
	"github.com/fatezhou/go-lang/pb/rpcTest"
	"google.golang.org/grpc"
)

func ClientTest(){
	conn, err := grpc.Dial("127.0.0.1:5090", grpc.WithInsecure())
	rpcClient := rpcTest.NewRPCFunctionClient(conn)
	req := &rpcTest.ReqParam{A:1, B:2}
	resp, err := rpcClient.Call(context.TODO(), req)
	fmt.Printf("err:%+v, resp:%+v", err, resp)
}