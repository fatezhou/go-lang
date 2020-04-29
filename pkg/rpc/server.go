package rpc

import (
	"context"
	"github.com/fatezhou/go-lang/pb/rpcTest"
	"google.golang.org/grpc"
	"net"
)

type RpcServer struct{}

func (rpc *RpcServer)Call(ctx context.Context, req *rpcTest.ReqParam)(resp *rpcTest.RespResult, err error){
	resp = &rpcTest.RespResult{}
	resp.Sum = req.GetA() + req.GetB()
	resp.Sub = req.GetA() - req.GetB()
	return resp, nil
}

func InitRPC(strIP string){
	lis, _ := net.Listen("tcp", strIP)
	s := grpc.NewServer()
	rpcTest.RegisterRPCFunctionServer(s, &RpcServer{})
	s.Serve(lis)
}