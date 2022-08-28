package handler

import (
	"context"
	"gRPC_Users/internal/dto"
	"gRPC_Users/internal/queue"
	"gRPC_Users/pkg"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"strconv"
)

func (h *Handler) GetUsers (_ context.Context, _ *empty.Empty) (*pkg.GetUserResponse, error) {
	users, err := h.services.User.GetUsers()
	if err != nil {
		log.Fatalln(err)
	}

	var response pkg.GetUserResponse

	for _, user := range users {
		u := pkg.User{
			Id: int32(user.Id),
			Name: user.Name,
			Email: user.Email,
		}
		response.Users = append(response.Users, &u)
	}

	return &response, nil
}

func (h *Handler) CreateUser(ctx context.Context, req *pkg.CreateUserRequest) (*pkg.CreateUserResponse, error) {
	command := dto.UserProfile{
		Name: req.Name
		Email: req.Email,
	}

	id, _ := h.services.User.CreateUser(command)

	w := kafka.Writer{Topic: queue.Topic, Addr: kafka.TCP(queue.BrokerAddress)}

	err := w.WriteMessages(ctx, kafka.Message{
		Key: []byte(strconv.Itoa(id)),
		Value: []byte("this is message: " + req.Name + " " + req.Email),
	})
	if err != nil {
		return nil, err
	}

	return &pkg.CreateUserResponse{Id: int32(id)}, nil
}

func (h *Handler) DeleteUser(_ context.Context, req *pkg.DeleteUserRequest) (*empty.Empty, error) {
	command := dto.UserId{Id: int(req.Id)}

	err := h.services.User.DeleteUser(command)
	if err != nil {
		log.Fatalln(err)
	}

	return new(emptypb.Empty), nil
}