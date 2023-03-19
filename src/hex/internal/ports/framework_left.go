package ports

import (
	"context"
	"hexArch/internal/adapters/framework/left/grpc/pb"
)

type GRPCPort interface {
	Run()
	GetAdd(ctx context.Context, req *pb.OperationParams) (*pb.Answer, error)
	GetSub(ctx context.Context, req *pb.OperationParams) (*pb.Answer, error)
	GetMul(ctx context.Context, req *pb.OperationParams) (*pb.Answer, error)
	GetDiv(ctx context.Context, req *pb.OperationParams) (*pb.Answer, error)
}
