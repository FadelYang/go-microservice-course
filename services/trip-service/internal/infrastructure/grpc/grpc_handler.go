package grpc

import (
	"context"
	"log"
	"ride-sharing/services/trip-service/internal/domain"
	pb "ride-sharing/shared/proto/trip"
	"ride-sharing/shared/types"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	userID := req.GetUserID()

	r, err := h.service.GetRoute(ctx, pickup, destination)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal, "failed to get route: %v", err)
	}

	// 1. Extimate the ride fares prices based on the route (ex: distance).
	estimatedFares := h.service.EstimatePackagesPriceWithRoute(r)
	// 2. Store the ride fares for the create trip (next lesson) to fetch and validate.
	fares, err := h.service.GenerateTripFares(ctx, estimatedFares, userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate the ride fares: %v", err)
	}

	return &pb.PreviewTripResponse{
		Route:     r.ToProto(),
		RideFares: domain.ToRideFaresProto(fares),
	}, nil
}

func (h *gRPCHandler) CreateTrip(ctx context.Context, req *pb.CreateTripRequest) (*pb.CreateTripResponse, error) {
	// 1. Fetch and validate the fare
	// 2. Call create trip.
	// 3. We also need to initialize an empty driver to the trip.
	// 4. Add acomment at the end of the function to publish an event the Async Comms

	return nil, status.Errorf(codes.Unimplemented, "method CreateTrip not implemented")
}
