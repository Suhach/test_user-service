package grpc

import (
	"context"
	pb "github.com/Suhach/test_protoc-cont/proto/user"
	"github.com/Suhach/test_user-service/internal/user"
)

type Handler struct {
	pb.UnimplementedUserServiceServer
	svc *user.Service
}

func NewHandler(svc *user.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user, err := h.svc.Create(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{Id: uint32(user.ID), Email: user.Email}, nil
}

func (h *Handler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := h.svc.Get(uint(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.GetUserResponse{Id: uint32(user.ID), Email: user.Email}, nil
}

func (h *Handler) GetAllUsers(ctx context.Context, req *pb.GetAllUsersRequest) (*pb.GetAllUsersResponse, error) {
	users, err := h.svc.GetAll()
	if err != nil {
		return nil, err
	}
	var pbUsers []*pb.User
	for _, u := range users {
		pbUsers = append(pbUsers, &pb.User{Id: uint32(u.ID), Email: u.Email})
	}
	return &pb.GetAllUsersResponse{Users: pbUsers}, nil
}

func (h *Handler) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	user, err := h.svc.Update(uint(req.Id), req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateUserResponse{Id: uint32(user.ID), Email: user.Email}, nil
}

func (h *Handler) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	if err := h.svc.Delete(uint(req.Id)); err != nil {
		return nil, err
	}
	return &pb.DeleteUserResponse{}, nil
}
