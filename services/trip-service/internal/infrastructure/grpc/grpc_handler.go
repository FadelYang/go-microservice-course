package grpc

import (
	"context"
	"log"
	"ride-sharing/services/trip-service/internal/domain"
	pb "ride-sharing/shared/proto/trip"
	"ride-sharing/shared/types"

	"google.golang.org/grpc"
)

type gRPCHandler struct {
	pb.UnimplementedTripServiceServer

	service domain.TripService
}

func NewGRPCHandler(server *grpc.Server, service domain.TripService) *gRPCHandler {
	handler := &gRPCHandler{
		service: service,
	}

	pb.RegisterTripServiceServer(server, handler)
	return handler
}

func (h *gRPCHandler) PreviewTrip(ctx context.Context, req *pb.PreviewTripRequest) (*pb.PreviewTripResponse, error) {
	pickup := &types.Coordinate{
		Latitude:  req.StartLocation.GetLatitude(),
		Longitude: req.StartLocation.GetLongitude(),
	}

	destination := &types.Coordinate{
		Latitude:  req.EndLocation.GetLatitude(),
		Longitude: req.EndLocation.GetLongitude(),
	}

	r, err := h.service.GetRoute(ctx, pickup, destination)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.PreviewTripResponse{
		Route: r.ToProto(),
	}, nil
}
