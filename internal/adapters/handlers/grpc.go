package handlers

import (
  "context"
  pb "github.com/botscommunity/proto/go"
)

type GRPCHandlers struct {
  pb.UnimplementedBotsServiceServer
}

func NewGRPCHandlers() *GRPCHandlers {
  return &GRPCHandlers{}
}

func (h *GRPCHandlers) GetBots(ctx context.Context, in *pb.GetBotsRequest) (*pb.GetBotsReply, error) {
  return &pb.GetBotsReply{}, nil
}

func (h *GRPCHandlers) GetBot(ctx context.Context, in *pb.GetBotRequest) (*pb.GetBotReply, error) {
  return &pb.GetBotReply{}, nil
}

func (h *GRPCHandlers) UpdateBot(ctx context.Context, in *pb.UpdateBotRequest) (*pb.UpdateBotReply, error) {
  return &pb.UpdateBotReply{}, nil
}

func (h *GRPCHandlers) CreateBot(ctx context.Context, in *pb.CreateBotRequest) (*pb.CreateBotReply, error) {
  return &pb.CreateBotReply{}, nil
}
