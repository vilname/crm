package service

import (
	"api/src/model"
	"api/src/repository"

	"github.com/gin-gonic/gin"
)

type AnswerService struct {
	ctx        *gin.Context
	repository *repository.AnswerRepository
}

func NewAnswerService(ctx *gin.Context) *AnswerService {
	return &AnswerService{
		ctx:        ctx,
		repository: repository.NewAnswerRepository(ctx),
	}
}

func (service *AnswerService) List(page uint16, limit uint16) ([]model.Answer, uint16, error) {
	offset := page*limit - limit

	total, err := service.repository.Count()

	if err != nil {
		return nil, total, err
	}

	answers, err := service.repository.GetList(limit, offset)

	return answers, total, err
}

func (service *AnswerService) Get(id int) (model.Answer, error) {

	return service.repository.Get(id)
}

func (service *AnswerService) Create(answer model.Answer) error {
	err := service.repository.Save(answer)

	if err != nil {
		return err
	}

	return nil
}
