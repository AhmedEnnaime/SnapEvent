package utils

import (
	"github.com/AhmedEnnaime/SnapEvent/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewUser() *pb.User {

	user := &pb.User{
		Name:     randomUserName(),
		Birthday: timestamppb.New(randomPastDate()),
		Email:    randomUserEmail(),
		Password: randomUserPassword(),
		Gender:   randomUserGender(),
	}

	return user
}

func NewEvent() *pb.Event {
	event := &pb.Event{
		Id:          uint32(randomID()),
		EventDate:   timestamppb.New(randomFutureDate()),
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
		EventId: 3,
		Type:    randomInviteType(),
	}
	return invite
}
