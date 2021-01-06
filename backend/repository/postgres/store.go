package postgres

import (
	"fmt"

	"holycode-task/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{db: db}
}

func Open(config PostgresConfig) (*Store, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Name)

	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	store := NewStore(db)
	return store, nil
}

func (store *Store) CheckStoreConnection() error {
	return store.db.DB().Ping()
}

func (store *Store) Close() error {
	return store.db.Close()
}

func (store *Store) MigrateTables() error {
	return store.db.AutoMigrate(&model.User{}, &model.Product{}, &model.InstagramAccount{}, &model.FacebookResponse{}).Error
}

func (store *Store) CreateUser(user *model.User) error {
	return store.db.Create(&user).Error
}

func (s *Store) DB() *gorm.DB {
	return s.db
}

func (s *Store) FindUserByID(id uint) (*model.User, error) {
	user := &model.User{}
	if err := s.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Store) FindUserByUsername(username string) (*model.User, error) {
	user := &model.User{}
	if err := s.db.Where("username = ?", username).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Store) FindAll() ([]model.User, error) {
	var users []model.User
	if err := s.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// Products
func (s *Store) FindAllProducts() ([]model.Product, error) {
	var products []model.Product
	if err := s.db.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (s *Store) FindProductByID(id uint) (*model.Product, error) {
	product := &model.Product{}
	if err := s.db.First(&product, id).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (s *Store) CreateProduct(p *model.Product) error {
	return s.db.Create(p).Error
}

func (s *Store) FindInstagramAccountByUsername(username string) (*model.InstagramAccount, error) {
	acc := &model.InstagramAccount{}
	if err := s.db.Where("username = ?", username).First(&acc).Error; err != nil {
		return nil, err
	}

	return acc, nil
}

func (s *Store) SaveInstagramAccount(acc *model.InstagramAccount) error {
	return s.db.Create(acc).Error
}

func (s *Store) FindFacebookAccountByUsername(username string) (*model.FacebookResponse, error) {
	acc := &model.FacebookResponse{}
	if err := s.db.Where("username = ?", username).First(&acc).Error; err != nil {
		return nil, err
	}

	return acc, nil
}

func (s *Store) SaveFacebookAccount(acc *model.FacebookResponse) error {
	return s.db.Create(acc).Error
}
