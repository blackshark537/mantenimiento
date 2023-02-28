package entities

import (
	"log"
	"time"

	rightPort "github.com/blackshark537/mantenimiento/src/core/Ports/Right"
)

const (
	MACRO = "MACRO"
	MICRO = "MICRO"

	GR       = "GR"
	KG       = "KG"
	LB       = "LB"
	FUNDA    = "FUNDA"
	SACO     = "SACO"
	GRANEL   = "GRANEL"
	BAZUCA   = "BAZUCA"
	TONELADA = "TONELADA"
)

type Ingrediente struct {
	Id           int       `json:"id" gorm:"primary_key"`
	AlimentoId   uint      `json:"alimento_id" gorm:"not null"`
	Nombre       string    `json:"nombre" gorm:"index; not null"`
	Descripcion  string    `json:"descripcion" gorm:"not null"`
	FichaTecnica string    `json:"ficha_tecnica" gorm:"not null"`
	Cantidad     int       `json:"cantidad" gorm:"not null"`
	Unidad       string    `json:"unidad" gorm:"type:tipo_unidad; not null; default:GR"`
	Precio       int       `json:"precio" gorm:"index; not null"`
	Tipo         string    `json:"tipo" gorm:"type:tipo_ingrediente; not null; default:MICRO"`
	SuplidorId   uint      `json:"suplidor_id" gorm:"not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	Uid          string    `json:"uid,omitempty" gorm:"index; not null"`
}

func (ing *Ingrediente) Migrate() {
	db := rightPort.GetDB()
	m := db.Migrator()
	if m.HasTable(ing) == true {
		return
	}
	ing.CreateTypes()
	m.CreateTable(ing)
}

func (ing *Ingrediente) CreateTypes() {
	db := rightPort.GetDB()
	err := db.Exec("CREATE TYPE tipo_ingrediente AS ENUM ('MACRO','MICRO');").Error
	if err != nil {
		log.Fatal(err)
	}
	err = db.Exec("CREATE TYPE tipo_unidad AS ENUM ('GR','KG', 'LB', 'FUNDA', 'SACO', 'GRANEL', 'BAZUCA', 'TONELADA');").Error
	if err != nil {
		log.Fatal(err)
	}
}
