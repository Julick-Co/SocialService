package social

import (
	pb "SocialService/proto/pkg/api/v1/social"
	"context"
	_ "fmt"
	"log"
)

type SocialServiceImpl struct {
	pb.UnimplementedSocialServiceServer
}

func (s *SocialServiceImpl) SendFriendRequest(ctx context.Context, req *pb.SendFriendRequestRequest,
) (*pb.SendFriendRequestResponse, error) {

	log.Println("Достучался до метода")

	return &pb.SendFriendRequestResponse{
		FriendRequest: &pb.FriendRequest{},
	}, nil
}

func (s *SocialServiceImpl) ListRequests(context.Context, *pb.ListRequestsRequest,
) (*pb.ListRequestsResponse, error) {

	log.Println("Достучался до метода")

	return &pb.ListRequestsResponse{
		Requests: []*pb.FriendRequest{},
	}, nil

}
