package db

import (
	"fmt"
	"joynet-assignment/models"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PgConfig struct {
	DbName     string
	DbPassword string
}

func (pgc *PgConfig) connectionString() string {
	return fmt.Sprintf("postgres://postgres:%s@postgres:5432/%s", pgc.DbPassword, pgc.DbName)
}

func NewDb(pgConfig PgConfig, log *logrus.Logger) (DbInterface, error) {
	fmt.Println(pgConfig.connectionString())
	dbconn, err := gorm.Open(postgres.Open(pgConfig.connectionString()), &gorm.Config{})
	if err != nil {
		return &db{}, fmt.Errorf("connecting to postgres db: %v", err)
	}
	return &db{
		Conn: dbconn,
		log:  log,
	}, nil
}

type DbInterface interface {
	Migrate()
	GetDevices() ([]models.Device, error)
	GetInterfaces() ([]models.Interface, error)
	SaveDevice(models.Device) error
	SaveInterfaces(models.Interface) error
}

type db struct {
	Conn *gorm.DB
	log  *logrus.Logger
}

func (db *db) Migrate() {
	db.log.Info("Running migrations")
	db.Conn.AutoMigrate(&models.Device{})
	db.Conn.AutoMigrate(&models.Interface{}, &models.IPInfo{}, &models.ARP{})
}

func (db *db) GetDevices() (devices []models.Device, err error) {
	result := db.Conn.Find(&devices)
	err = result.Error
	return devices, err
}
func (db *db) GetInterfaces() (ints []models.Interface, err error) {
	result := db.Conn.Model(&models.Interface{}).Preload("IPInfo").Find(&ints)
	err = result.Error
	return ints, err
}
func (db *db) SaveDevice(device models.Device) error {
	result := db.Conn.Save(&device)
	return result.Error
}
func (db *db) SaveInterfaces(intfc models.Interface) error {
	result := db.Conn.Save(&intfc)
	return result.Error
}
