package handler

import (
	"context"
	"encgo/model"
	pb "encgo/proto/user-experience"
	"encgo/service"
	"fmt"
	"log"
	"strconv"
)

type UserExperienceHandler struct {
	UserExperienceService *service.UserExperienceService
	pb.UnimplementedUserExperienceServiceServer
}

func (handler *UserExperienceHandler) FindByUserId(ctx context.Context, req *pb.FindByUserIdRequest) (*pb.UserExperience, error) {
	userExperience, err := handler.UserExperienceService.FindByUserId(int(req.UserId))
	if err != nil {
		return nil, fmt.Errorf("user experience with user id %d not found", req.UserId)
	}

	return &pb.UserExperience{
		Id:     int32(userExperience.ID),
		UserId: int32(userExperience.UserID),
		Xp:     int32(userExperience.XP),
		Level:  int32(userExperience.Level),
	}, nil
}

func (handler *UserExperienceHandler) AddXP(ctx context.Context, req *pb.AddXPRequest) (*pb.UserExperience, error) {
	userExperience, err := handler.UserExperienceService.AddXP(int(req.Id), int(req.Xp))
	if err != nil {
		return nil, fmt.Errorf("user experience with id %d not found: %v", req.Id, err)
	}

	return &pb.UserExperience{
		Id:     int32(userExperience.ID),
		UserId: int32(userExperience.UserID),
		Xp:     int32(userExperience.XP),
		Level:  int32(userExperience.Level),
	}, nil
}
func (handler *UserExperienceHandler) Create(ctx context.Context, req *pb.UserExperience) (*pb.UserExperience, error) {
	// Check for nil on the request and user experience data
	if req == nil {
		return nil, fmt.Errorf("request or user experience data is nil")
	}

	// Assuming the UserExperience struct fields are properly typed and don't need conversion
	userExperience := model.UserExperience{
		ID:     int(req.Id),
		UserID: int(req.UserId),
		XP:     int(req.Xp),
		Level:  int(req.Level),
	}

	if handler.UserExperienceService == nil {
		return nil, fmt.Errorf("user experience service is not initialized")
	}

	createdUserExperience, err := handler.UserExperienceService.Create(&userExperience)
	if err != nil {
		log.Printf("Error creating user experience: %v", err) // Use your preferred logging method
		return nil, fmt.Errorf("failed to create user experience: %v", err)
	}

	return &pb.UserExperience{
		Id:     int32(createdUserExperience.ID),
		UserId: int32(createdUserExperience.UserID),
		Xp:     int32(createdUserExperience.XP),
		Level:  int32(createdUserExperience.Level),
	}, nil
}

func (handler *UserExperienceHandler) Delete(ctx context.Context, req *pb.UserExperience) (*pb.DeleteUserExperienceResponse, error) {
	idStr := strconv.Itoa(int(req.Id))
	err := handler.UserExperienceService.Delete(idStr)
	if err != nil {
		return &pb.DeleteUserExperienceResponse{}, fmt.Errorf("failed to delete user experience with id %d: %v", req.Id, err)
	}
	return &pb.DeleteUserExperienceResponse{}, nil
}

func (handler *UserExperienceHandler) Update(ctx context.Context, req *pb.UserExperience) (*pb.UserExperience, error) {
	userExperience := model.UserExperience{
		ID:     int(req.Id),
		UserID: int(req.UserId),
		XP:     int(req.Xp),
		Level:  int(req.Level),
	}
	updatedUserExperience, err := handler.UserExperienceService.Update(&userExperience)
	if err != nil {
		return nil, err
	}

	return &pb.UserExperience{
		Id:     int32(updatedUserExperience.ID),
		UserId: int32(updatedUserExperience.UserID),
		Xp:     int32(updatedUserExperience.XP),
		Level:  int32(updatedUserExperience.Level),
	}, nil
}
