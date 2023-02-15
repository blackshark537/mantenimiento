package entities

import (
	"time"

	rightPort "github.com/blackshark537/mantenimiento/src/core/Ports/Right"
)

type Alimento struct {
	Id        int           `json:"id" gorm:"primary_key"`
	Nombre    string        `json:"nombre" gorm:"index; not null"`
	Macros    []Ingrediente `json:"macros" gorm:"foreignKey:AlimentoId"`
	Micros    []Ingrediente `json:"micros" gorm:"foreignKey:AlimentoId"`
	Precio    int           `json:"precio" gorm:"index; not null"`
	CreatedAt time.Time     `json:"created_at" gorm:"autoCreateTime"`
	Owner     string        `json:"owner" gorm:"not null"`
}

func (alm *Alimento) Migrate() {
	db := rightPort.GetDB()
	m := db.Migrator()
	if m.HasTable(alm) == true {
		return
	}
	m.CreateTable(alm)
}
