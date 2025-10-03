package service

import (
	"fabiloco/hotel-trivoli-api/pkg/entities"

	"gorm.io/gorm"
)

type Repository interface {
	Create(data *entities.Service) (*entities.Service, error)
	Read() (*[]entities.Service, error)
	ReadPaginated(params *entities.PaginationParams) (*[]entities.Service, int64, error)
	Update(id uint, data *entities.Service) (*entities.Service, error)
	Delete(id uint) (*entities.Service, error)
	ReadById(id uint) (*entities.Service, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Read() (*[]entities.Service, error) {
	var services []entities.Service

	r.db.Find(&services)

	return &services, nil
}

func (r *repository) ReadPaginated(params *entities.PaginationParams) (*[]entities.Service, int64, error) {
	var services []entities.Service
	var total int64

	params.Normalize()

	if err := r.db.Model(&entities.Service{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := params.GetOffset()
	limit := params.GetLimit()

	if err := r.db.Offset(offset).Limit(limit).Find(&services).Error; err != nil {
		return nil, 0, err
	}

	return &services, total, nil
}

func (r *repository) ReadById(id uint) (*entities.Service, error) {
	var service entities.Service

	result := r.db.First(&service, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &service, nil
}

func (r *repository) Create(data *entities.Service) (*entities.Service, error) {
	var service entities.Service

	service = entities.Service{
		Name:  data.Name,
		Price: data.Price,
	}

	result := r.db.Create(&service)

	if result.Error != nil {
		return nil, result.Error
	}

	return &service, nil
}

func (r *repository) Update(id uint, data *entities.Service) (*entities.Service, error) {
	service, error := r.ReadById(id)

	if error != nil {
		return nil, error
	}

	result := r.db.Model(&service).Updates(
		entities.Service{
			Name:  data.Name,
			Price: data.Price,
		},
	)

	if result.Error != nil {
		return nil, result.Error
	}

	return service, nil
}

func (r *repository) Delete(id uint) (*entities.Service, error) {
	service, error := r.ReadById(id)

	if error != nil {
		return nil, error
	}

	result := r.db.Delete(&service, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return service, nil
}
