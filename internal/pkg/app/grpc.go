package app

import (
  "fmt"
  "net"
  "google.golang.org/grpc"
  pb "github.com/botscommunity/proto/go"
  "github.com/botscommunity/backend/internal/adapters/handlers"
)

type GRPCServer struct {
  server *grpc.Server
}

func NewGRPCServer() *GRPCServer {
  var (
    server = grpc.NewServer()
    handlers = handlers.NewGRPCHandlers()
  )

  pb.RegisterBotsServiceServer(server, handlers)

  return &GRPCServer{
    server: server,
  }
}

func (s *GRPCServer) Run(port string) error {
  listen, err := net.Listen("tcp", "localhost:" + port)
  if err != nil {
    return err
  }
  fmt.Println("Running gRPC server")
  return s.server.Serve(listen)
}
