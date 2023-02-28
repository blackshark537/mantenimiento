package entities

import (
	"time"

	rightPort "github.com/blackshark537/mantenimiento/src/core/Ports/Right"
)

type Alimento struct {
	Id        int           `json:"id" gorm:"primary_key"`
	Nombre    string        `json:"nombre" gorm:"index; not null"`
	Macros    []Ingrediente `json:"macros" gorm:"foreignKey:AlimentoId; constraint:OnDelete:CASCADE"`
	Micros    []Ingrediente `json:"micros" gorm:"foreignKey:AlimentoId; constraint:OnDelete:CASCADE"`
	Precio    int           `json:"precio" gorm:"index; not null"`
	CreatedAt time.Time     `json:"created_at" gorm:"autoCreateTime"`
	Uid       string        `json:"uid,omitempty" gorm:"index; not null"`
}

func (alm *Alimento) Migrate() {
	db := rightPort.GetDB()
	m := db.Migrator()
	if m.HasTable(alm) == true {
		return
	}
	m.CreateTable(alm)
}
