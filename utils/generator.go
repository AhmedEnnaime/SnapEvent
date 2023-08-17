package utils

import (
	"github.com/AhmedEnnaime/SnapEvent/pb"
)

func NewUser() *pb.User {

	user := &pb.User{
		Name:     randomUserName(),
		Birthday: randomPastDate(),
		Email:    randomUserEmail(),
		Password: randomUserPassword(),
		Gender:   randomUserGender(),
	}

	return user
}

func NewEvent() *pb.Event {
	event := &pb.Event{
		Id:          uint32(randomID()),
		EventDate:   randomFutureDate(),
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
