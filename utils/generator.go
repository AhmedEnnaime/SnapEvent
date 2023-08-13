package utils

import (
	"github.com/AhmedEnnaime/SnapEvent/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewUser() *pb.User {

	user := &pb.User{
		Name:     randomUserName(),
		Birthday: timestamppb.New(randomDate()),
		Email:    randomUserEmail(),
		Password: randomUserPassword(),
		Gender:   randomUserGender(),
	}

	return user
}

func NewEvent() *pb.Event {
	event := &pb.Event{
		EventDate:   timestamppb.New(randomDate()),
		Time:        randomEventTime(),
		Description: randomDescription(),
		City:        randomCity(),
		Location:    randomLocation(),
		Poster:      randomPosterLink(),
		CreatorId:   randomEntityId(),
		Status:      randomStatus(),
	}
	return event
}

func NewInvite() *pb.Invite {
	invite := &pb.Invite{
		UserId:  randomEntityId(),
		EventId: randomEntityId(),
		Type:    randomInviteType(),
	}
	return invite
}
