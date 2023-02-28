package entities

import (
	"encoding/json"
	"fmt"

	rightPort "github.com/blackshark537/mantenimiento/src/core/Ports/Right"
	"gorm.io/gorm"
)

type Entity[t interface{}] struct {
	Data     []byte
	Entity   t
	Entities []t
	Id       uint
}

func (e *Entity[t]) GetAll() *gorm.DB {
	json.Unmarshal(e.Data, &e.Entity)
	return rightPort.GetDB().Model(e.Entity).Where(e.Entity)
}

func (e *Entity[t]) GetOne(key string, value interface{}) *gorm.DB {
	return rightPort.GetDB().Model(e.Entity)
}

func (e *Entity[t]) Create() error {
	return rightPort.GetDB().Model(e.Entity).Create(&e.Entity).Error
}

func (e *Entity[t]) Update() *gorm.DB {
	return rightPort.GetDB().Model(e.Entity)
}

func (e *Entity[t]) Delete() *gorm.DB {
	return rightPort.GetDB().Model(e.Entity)
}

func (e *Entity[t]) Clear(table string) error {
	sql := fmt.Sprintf("DELETE FROM %s;", table)
	return rightPort.GetDB().Exec(sql).Error
}
