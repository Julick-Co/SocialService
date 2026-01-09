package social

import (
	pb "SocialService/proto/pkg/api/social"
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

func (s *SocialServiceImpl) AcceptFriendRequest(context.Context, *pb.AcceptFriendRequestRequest,
) (*pb.AcceptFriendRequestResponse, error) {

	log.Println("Достучался до метода AcceptFriendRequest")

	return &pb.AcceptFriendRequestResponse{
		RequestId: "generated Id",
		Status:    pb.FriendRequestStatus_ACCEPTED,
	}, nil
}

func (s *SocialServiceImpl) DeclineFriendRequest(context.Context, *pb.DeclineFriendRequestRequest,
) (*pb.DeclineFriendRequestResponse, error) {

	log.Println("Достучался до метода DeclineFriendRequest")

	return &pb.DeclineFriendRequestResponse{
		RequestId: "generated Id",
		Status:    pb.FriendRequestStatus_DECLINED,
	}, nil
}

func (s *SocialServiceImpl) RemoveFriend(context.Context, *pb.RemoveFriendRequest,
) (*pb.RemoveFriendResponse, error) {

	log.Println("Достучался до метода RemoveFriend")

	return &pb.RemoveFriendResponse{}, nil
}

func (s *SocialServiceImpl) ListFriends(context.Context, *pb.ListFriendsRequest,
) (*pb.ListFriendsResponse, error) {

	log.Println("Достучался до метода ListFriends")

	return &pb.ListFriendsResponse{
		FriendUserIds: []string{},
	}, nil
}
