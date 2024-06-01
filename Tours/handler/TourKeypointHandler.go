package handler

import (
	"context"
	"database-example/model"
	pb "database-example/proto/tours"
	"database-example/service"
	"errors"
	"fmt"
	"strconv"
)

type TourKeypointHandler struct {
	pb.UnimplementedTourServiceServer
	TourKeypointService *service.TourKeypointService
}

func (handler *TourKeypointHandler) Get(ctx context.Context, req *pb.TourKeypointIdRequest) (*pb.TourKeypoint, error) {
	tourKeypoint, err := handler.TourKeypointService.Find(strconv.Itoa(int(req.Id)))

	if err != nil {
		return nil, fmt.Errorf("tour keypoint id %d not found", req.Id)
	}
	return modelToProto(tourKeypoint), nil
}

func (handler *TourKeypointHandler) Create(ctx context.Context, req *pb.TourKeypoint) (*pb.TourKeypoint, error) {
	if req == nil {
		return nil, fmt.Errorf("request or tour keypoint data is nil")
	}

	tourKeypoint := protoToModel(req)

	if tourKeypoint == nil {
		return nil, errors.New("Error while parsing tour keypoint! (create)")
	}
	createdTourKeypoint, err := handler.TourKeypointService.Create(tourKeypoint)
	if err != nil {
		if err != nil {
			return nil, fmt.Errorf("failed to create tour keypoint: %v", err)
		}
	}

	return modelToProto(createdTourKeypoint), nil
}

func (handler *TourKeypointHandler) Update(ctx context.Context, req *pb.TourKeypoint) (*pb.TourKeypoint, error) {
	if req == nil {
		return nil, fmt.Errorf("request or tour keypoint data is nil")
	}

	tourKeypoint := protoToModel(req)
	if tourKeypoint == nil {
		return nil, errors.New("Error while parsing tour keypoint! (update)")
	}
	updatedTourKeypoint, err := handler.TourKeypointService.Update(tourKeypoint)
	if err != nil {
		return nil, fmt.Errorf("failed to update tour keypoint: %v", err)
	}
	return modelToProto(updatedTourKeypoint), nil
}

func (handler *TourKeypointHandler) Delete(ctx context.Context, req *pb.TourKeypointIdRequest) (*pb.EmptyResponse, error) {

	err := handler.TourKeypointService.Delete(strconv.Itoa(int(req.Id)))
	if err != nil {
		return nil, err
	}
	return &pb.EmptyResponse{}, nil
}

func modelToProto(model *model.TourKeypoint) *pb.TourKeypoint {

	return &pb.TourKeypoint{
		Id:             int32(model.ID),
		Name:           string(model.Name),
		Description:    string(model.Description),
		Image:          string(model.Image),
		Latitude:       float64(model.Latitude),
		Longitude:      float64(model.Longitude),
		TourId:         int32(model.TourID),
		Secret:         string(model.Secret),
		PositionInTour: int32(model.PositionInTour),
		PublicPointId:  int32(model.PublicPointID),
	}
}

func protoToModel(req *pb.TourKeypoint) *model.TourKeypoint {

	return &model.TourKeypoint{
		ID:             int(req.GetId()),
		Name:           req.GetName(),
		Description:    req.GetDescription(),
		Image:          req.GetImage(),
		Latitude:       float64(req.GetLatitude()),
		Longitude:      float64(req.GetLongitude()),
		TourID:         int(req.GetTourId()),
		Secret:         req.GetSecret(),
		PositionInTour: int(req.GetPositionInTour()),
		PublicPointID:  int(req.GetPublicPointId()),
	}
}
