package topsecret

import (
	"errors"
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
	var satelliteDB Records

	if err := r.db.Where("name = ?", satellite.Name).First(&satelliteDB).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			logrus.Errorf("Error getting satellite: %v err:%v", satellite.Name, err)
			return err
		}
	}

	if satelliteDB.Name == "" {
		if err := r.db.Save(satellite.ToRecord()).Error; err != nil {
			logrus.Errorf("Error saving satellite: %v err:%v", satellite.Name, err)
			return err
		}
		return nil
	}

	record := satellite.ToRecord()
	if err := r.db.Where("name = ?", satellite.Name).Model(&satelliteDB).Updates(record).Error; err != nil {
		logrus.Errorf("Error updating satellite: %v err:%v", satellite.Name, err)
		return err
	}

	return nil
}

// Get gets a satellite from the database
func (r *SQLiteRepository) Get(name string) (*Satellite, error) {
	var record Records
	if err := r.db.Where("name = ?", name).First(&record).Error; err != nil {
		logrus.Errorf("Error getting satellite: %v err:%v", name, err)
		return nil, err
	}
	return record.ToSatellite(), nil
}

// GetInfo gets the information of satellites from the database
func (r *SQLiteRepository) GetInfo() ([]Satellite, error) {
	var records []Records
	if err := r.db.Find(&records).Error; err != nil {
		logrus.Errorf("Error getting satellites: %v", err)
		return nil, err
	}

	satellites := make([]Satellite, len(records))
	for i, _ := range records {
		satellites[i] = *records[i].ToSatellite()
	}

	return satellites, nil
}

// Find gets all satellites from the database
func (r *SQLiteRepository) Find() ([]Satellite, error) {
	var records []Records
	if err := r.db.Find(&records).Error; err != nil {
		logrus.Errorf("Error getting satellites: %v", err)
		return nil, err
	}

	satellites := make([]Satellite, len(records))
	for i, _ := range records {
		satellites[i] = *records[i].ToSatellite()
	}

	return satellites, nil
}

// Delete deletes all satellites from the database
func (r *SQLiteRepository) Delete() error {
	if err := r.db.Where("name != ?", "").Delete(&Records{}).Error; err != nil {
		logrus.Errorf("Error deleting satellites: %v", err)
		return err
	}
	return nil
}
