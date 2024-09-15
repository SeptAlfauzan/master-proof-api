package service

import "master-proof-api/dto"

type ActivityService interface {
	CreateActivity(request *dto.CreateActivityRequest) error
}