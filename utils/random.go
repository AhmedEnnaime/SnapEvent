package utils

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/AhmedEnnaime/SnapEvent/pb"
	"github.com/google/uuid"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func randomID() uint {
	return uint(uuid.New().ID())
}

func randomUserGender() pb.GENDER {
	switch rand.Intn(2) {
	case 1:
		return pb.GENDER_MALE
	case 2:
		return pb.GENDER_FEMALE
	default:
		return pb.GENDER_MALE
	}

}

func randomUserName() string {
	return randomStringFromSet("Ahmed", "Aya", "MOhammed", "Aymen", "Rajae", "Adnane", "Redoine", "Kaoutar", "Amal")

}

func randomPastDate() time.Time {
	today := time.Now()

	// Generate a random number of days in the past
	randomDays := rand.Intn(today.YearDay())

	// Calculate the random past date
	randomDate := today.AddDate(0, 0, -randomDays)

	return randomDate
}

func randomFutureDate() time.Time {
	today := time.Now()

	// Generate a random number of days in the future
	randomDays := rand.Intn(365) // You can adjust the range as needed

	// Calculate the random future date
	randomDate := today.AddDate(0, 0, randomDays)

	return randomDate
}

func randomUserEmail() string {
	randomNumber := rand.Intn(9000) + 1000

	email := fmt.Sprintf("user%d@gmail.com", randomNumber)
	return email
}

func randomUserPassword() string {
	return randomStringFromSet("admin1234", "root12345")
}

func randomEventTime() string {
	return randomStringFromSet("19:00", "4:00", "20:15", "2:00", "22:30", "1:10", "21:45", "23:00", "00:15")
}

func randomStatus() pb.STATUS {
	switch rand.Intn(3) {
	case 1:
		return pb.STATUS_OPEN
	case 2:
		return pb.STATUS_CLOSED
	case 3:
		return pb.STATUS_Invitation
	default:
		return pb.STATUS_OPEN

	}
}

func randomEntityId() uint32 {
	switch rand.Intn(3) {
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 3
	default:
		return 1
	}
}

func randomDescription() string {
	return randomStringFromSet(
		"Join us for a day of fun and excitement!",
		"Discover new opportunities and insights.",
		"Experience the thrill of live music and entertainment.",
		"Learn from industry experts and leaders.",
		"Connect with like-minded individuals.",
		"Explore the latest trends in technology.",
		"Enjoy delicious food and drinks with friends.")
}

func randomCity() string {
	return randomStringFromSet(
		"Casablanca",
		"Rabat",
		"Marrakech",
		"Fes",
		"Agadir",
		"Tangier",
		"Meknes",
		"Kenitra",
		"Oujda",
		"Tetouan",
		"Salé",
		"El Jadida",
		"Nador",
		"Beni Mellal",
		"Settat")
}

func randomLocation() string {
	return randomStringFromSet("Maarif", "Boulevard C", "Sidi bouzid", "Agdal", "Medina Qdima", "Rue Mohammed V", "Place Moulay Youssef", "Derb Soltan")
}

func randomPosterLink() string {
	return randomStringFromSet("https://www.google.com/", "https://www.youtube.com/", "https://www.linkedin.com/feed/", "https://github.com/")
}

func randomInviteType() pb.TYPE {
	switch rand.Intn(2) {
	case 1:
		return pb.TYPE_ATTENDEE
	case 2:
		return pb.TYPE_VIP
	default:
		return pb.TYPE_ATTENDEE
	}
}

func randomStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}
