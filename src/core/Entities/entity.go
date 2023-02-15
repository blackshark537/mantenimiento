package entities

import (
	rightPort "github.com/blackshark537/mantenimiento/src/core/Ports/Right"
	"gorm.io/gorm"
)

type Entity[t interface{}] struct {
	Entity   t
	Entities []t
}

func (e *Entity[t]) GetAll() *gorm.DB {
	return rightPort.GetDB().Model(e.Entity)
}

func (e *Entity[t]) GetOne(key string, value interface{}) *gorm.DB {
	return rightPort.GetDB().Model(e.Entity)
}

func (e *Entity[t]) Create() error {
	return rightPort.GetDB().Model(e.Entity).Create(&e.Entity).Error
}

func (e *Entity[t]) Update() error {
	return rightPort.GetDB().Model(e.Entity).Updates(&e.Entity).Error
}

func (e *Entity[t]) Delete() error {
	return rightPort.GetDB().Model(e.Entity).Delete(&e.Entity).Error
}
