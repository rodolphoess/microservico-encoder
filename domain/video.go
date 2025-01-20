package domain

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type Video struct {
	ID         string    `valid:"uuid"`
	ResourceID string    `valid:"notnull"`
	FilePath   string    `valid:"notnull"`
	CreatedAt  time.Time `valid:"-"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true) //Valida e garante que todos os campos requeridos serão seguidos.
}

func NewVideo() *Video {
	return &Video{} //Retorna a referência ao struct Video
}

func (video *Video) Validate() error {
	_, err := govalidator.ValidateStruct(video)

	if err != nil {
		return err
	}

	return nil
}
