package repository

import (
	"gorm.io/gorm"
	"reflect"
)

type Repository interface {
	FindAll() (interface{}, error)
	FindById(interface{}) (interface{}, error)
	Save(interface{}) (interface{}, error)
	DeleteById(interface{}) error
	Status() error
}

type CrudRepository struct {
	db *gorm.DB
	key interface{}
	domain interface{}
}

func NewCrudRepository(db *gorm.DB, domain interface{}, key interface{}) *CrudRepository {
	return &CrudRepository{db: db, domain: domain, key: key}
}

func (repo *CrudRepository) FindAll() (interface{}, error){
	domainType := reflect.TypeOf(repo.domain)
	domains := reflect.New(reflect.SliceOf(domainType)).Interface()

	if error := repo.db.Model(repo.domain).Find(domains).Error; error != nil {
		return nil, error
	}

	return domains,nil
}

func (repo *CrudRepository) FindById(uuid interface{}) (interface{}, error){
	domain := reflect.New(reflect.TypeOf(repo.domain)).Interface()
	if error := repo.db.Model(repo.domain).First(domain, uuid).Error; error != nil {
		return nil, error
	}
	return domain,nil
}

func (repo *CrudRepository) Save(domainInstance interface{}) (interface{}, error){
	if error := repo.db.Create(domainInstance).Error; error != nil {
		return nil, error
	}
	return domainInstance,nil
}

func (repo *CrudRepository) DeleteById(domainInstance interface{}) error{
	if error := repo.db.Delete(repo.domain,domainInstance).Error; error != nil {
		return error
	}
	return nil
}

func (repo *CrudRepository) Status() error{
	var row int
	err := repo.db.Raw("select 1;").Scan(&row).Error
	if err != nil {
		return err
	}
	return nil
}
