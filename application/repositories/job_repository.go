package repositories

import (
	"domain/domain"
	"fmt"
	"github.com/jinzhu/gorm"
)

type JobRepository interface {
	Insert(job *domain.Job) (*domain.Job, error)
	Find(id string) (*domain.Job, error)
	Update(job *domain.Job) (*domain.Job, error)
}

type JobRepositoryDB struct {
	Db *gorm.DB
}

func (repo JobRepositoryDB) Insert(job *domain.Job) (*domain.Job, error) {
	err := repo.Db.Create(job).Error

	if err != nil {
		return nil, err
	}

	return job, nil
}

func (repo JobRepositoryDB) Find(id string) (*domain.Job, error) {
	var job domain.Job
	repo.Db.Preload("Video").First(&job, "id = ?", id)

	if job.ID == "" {
		return nil, fmt.Errorf("job %s not found", id)
	}

	return &job, nil
}

func (repo JobRepositoryDB) Update(job *domain.Job) (*domain.Job, error) {
	err := repo.Db.Save(job).Error

	if err != nil {
		return nil, err
	}

	return job, nil
}
