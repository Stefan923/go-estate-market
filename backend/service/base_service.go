package service

import (
	"backend/repository"
	"backend/util"
	"context"
)

type BaseService[T any, Tc any, Tu any, Tr any] struct {
	Repository *repository.BaseRepository[T]
}

func (service BaseService[T, Tc, Tu, Tr]) GetById(id uint) (*Tr, error) {
	object, err := service.Repository.FindById(id)
	if err != nil {
		return nil, err
	}

	convertedObject, _ := util.ConvertTo[Tr](object)
	return convertedObject, nil
}

func (service BaseService[T, Tc, Tu, Tr]) Save(context context.Context, object *Tc) (*Tr, error) {
	convertedObject, _ := util.ConvertTo[T](object)
	createdObject, err := service.Repository.Save(context, convertedObject)
	if err != nil {
		return nil, err
	}

	convertedCreatedObject, _ := util.ConvertTo[Tr](createdObject)
	return convertedCreatedObject, nil
}

func (service BaseService[T, Tc, Tu, Tr]) Update(context context.Context, id uint, object *Tu) (*Tr, error) {
	convertedObject, _ := util.ConvertTo[T](object)
	updatedObject, err := service.Repository.Update(context, id, convertedObject)
	if err != nil {
		return nil, err
	}

	convertedUpdatedObject, _ := util.ConvertTo[Tr](updatedObject)
	return convertedUpdatedObject, nil
}

func (service BaseService[T, Tc, Tu, Tr]) Delete(context context.Context, id uint) error {
	return service.Repository.Delete(context, id)
}
