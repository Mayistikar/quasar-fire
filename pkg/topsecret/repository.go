package topsecret

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// SQLiteRepository struct to represent a repository for satellite
type SQLiteRepository struct {
	db *gorm.DB
}

// NewRepository creates a new repository for satellite
func NewRepository(db *gorm.DB) *SQLiteRepository {
	return &SQLiteRepository{db: db}
}

// Save saves a satellite in the database
func (r *SQLiteRepository) Save(satellite *Satellite) error {
	if err := r.db.Save(satellite).Error; err != nil {
		logrus.Errorf("Error saving satellite: %v err:%v", *satellite, err)
		return err
	}
	return nil
}

// Get gets a satellite from the database
func (r *SQLiteRepository) Get(name string) (*Satellite, error) {
	var satellite Satellite
	if err := r.db.Where("name = ?", name).First(&satellite).Error; err != nil {
		logrus.Errorf("Error getting satellite: %v err:%v", name, err)
		return nil, err
	}
	return &satellite, nil
}

// GetInfo gets the information of satellites from the database
func (r *SQLiteRepository) GetInfo() ([]Records, error) {
	var satellites []Records
	if err := r.db.Find(&satellites).Error; err != nil {
		logrus.Errorf("Error getting satellites: %v", err)
		return nil, err
	}
	return satellites, nil
}
