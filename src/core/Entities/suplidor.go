package entities

import (
	"time"

	rightPort "github.com/blackshark537/mantenimiento/src/core/Ports/Right"
)

type Suplidor struct {
	Id          int         `json:"id" gorm:"primary_key"`
	Nombre      string      `json:"nombre" gorm:"index; not null"`
	Direccion   string      `json:"direccion" gorm:"not null"`
	Contactos   []Contacto  `json:"contactos" gorm:"foreignKey:EmpresaId;"`
	Ingrediente Ingrediente `json:"ingrediente" gorm:"foreignKey:SuplidorId; not null"`
	CreatedAt   time.Time   `json:"created_at" gorm:"autoCreateTime"`
	Owner       string      `json:"owner" gorm:"not null"`
}

func (spl *Suplidor) Migrate() {
	db := rightPort.GetDB()
	m := db.Migrator()
	if m.HasTable(spl) == true {
		return
	}
	m.CreateTable(spl)
}
