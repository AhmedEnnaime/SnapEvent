package store

import (
	"github.com/AhmedEnnaime/SnapEvent/internal/models"
	"github.com/jinzhu/gorm"
)

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (s *UserStore) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := s.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserStore) GetByEmail(email string) (*models.User, error) {
	var m models.User
	if err := s.db.Where("email = ?", email).First(&m).Error; err != nil {
		return nil, err
	}

	return &m, nil
}

func (s *UserStore) GetByID(id uint) (*models.User, error) {
	var m models.User
	if err := s.db.Find(&m, id).Error; err != nil {
		return nil, err
	}

	return &m, nil
}

func (s *UserStore) Create(m *models.User) error {
	return s.db.Create(m).Error
}

func (s *UserStore) Update(m *models.User) error {
	return s.db.Model(m).Update(m).Error
}

func (s *UserStore) GetEventsCreatedByUser(userID uint) ([]models.Event, error) {
	var events []models.Event
	if err := s.db.Model(&models.Event{}).Where("user_id = ?", userID).Find(&events).Error; err != nil {
		return nil, err
	}

	return events, nil
}

func (s *UserStore) GetEventsParticipatedByUser(userID uint) ([]models.Event, error) {
	var user models.User
	if err := s.db.Preload("Events").Find(&user, userID).Error; err != nil {
		return nil, err
	}

	return user.Events, nil
}
