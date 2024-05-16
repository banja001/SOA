package handler

import (
	"context"
	"database-example/model"
	"database-example/proto/tours"
	"database-example/service"
	"errors"
)

type TourKeypointHandler struct {
	tours.UnimplementedTourServiceServer
	TourKeypointService *service.TourKeypointService
}

func (handler *TourKeypointHandler) Get(ctx context.Context, req *tours.TourKeypointIdRequest) (*tours.TourKeypoint, error) {
	tourKeypoint, err := handler.TourKeypointService.Find(string(req.GetId()))
	if err != nil {
		return nil, err
	}
	return modelToProto(tourKeypoint), nil
}

func (handler *TourKeypointHandler) Create(ctx context.Context, req *tours.TourKeypoint) (*tours.TourKeypoint, error){
	tourKeypoint := protoToModel(req)

	if tourKeypoint == nil {
		return nil, errors.New("Error while parsing tour keypoint! (create)")
	}
	createdTourKeypoint, err := handler.TourKeypointService.Create(tourKeypoint)
	if err != nil {
		return nil, err
	}

	return tours.TourKeypoint{createdTourKeypoint}, nil
}

func (handler *TourKeypointHandler) Update(ctx context.Context, req *tours.TourKeypoint) (*tours.TourKeypoint, error){
	tourKeypoint := protoToModel(req)
	if tourKeypoint == nil {
		return nil, errors.New("Error while parsing tour keypoint! (update)")
	}
	updatedTourKeypoint, err := handler.TourKeypointService.Update(tourKeypoint)
	if err != nil {
		return nil, err
	}
	return &tours.TourKeypoint{TourKeypoint: updatedTourKeypoint}, nil
}

func (handler *TourKeypointHandler) Delete(ctx context.Context, req *tours.TourKeypointIdRequest) (*tours.EmptyResponse, error)  {

	err := handler.TourKeypointService.Delete(string(req.GetId()))
	if err != nil {
		return nil, err
	}
	return &tours.EmptyResponse{}, nil
}

func modelToProto(model *model.TourKeypoint)*tours.TourKeypoint {

    return &tours.TourKeypoint{
        Id:               int32(model.ID),
        Name:             string(model.Name),
        Description:	  string(model.Description),
        Image:            string(model.Image),
		Latitude:         float64(model.Latitude),
		Longitude:        float64(model.Longitude),
		TourId:           int32(model.TourID),
		Secret:           string(model.Secret),
		PositionInTour:   int32(model.PositionInTour),
		PublicPointId:    int32(model.PublicPointID),
    }
}

func protoToModel(req *tours.TourKeypoint)*model.TourKeypoint {

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
