package utils

import (
	"github.com/AhmedEnnaime/SnapEvent/internal/models"
	"github.com/AhmedEnnaime/SnapEvent/pb"
)

func MapModelGenderToProtoGender(gender models.GENDER) pb.GENDER {
	switch gender {
	case models.MALE:
		return pb.GENDER_MALE
	case models.FEMALE:
		return pb.GENDER_FEMALE
	default:
		// Handle any other cases or errors here
		return pb.GENDER_MALE
	}
}
