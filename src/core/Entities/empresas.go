package entities

import (
	"fmt"
	"time"

	"github.com/rodaine/table"

	rightPort "github.com/blackshark537/mantenimiento/src/core/Ports/Right"
	"github.com/fatih/color"
)

type Empresa struct {
	Id          uint        `json:"id" gorm:"primary_key"`
	Nombre      string      `json:"nombre" gorm:"index; not null"`
	Contactos   []Contacto  `json:"contactos" gorm:"foreignKey:EmpresaId; constraint:OnDelete:CASCADE"`
	Direccion   string      `json:"direccion" gorm:"index; not null"`
	Provincia   string      `json:"provincia" gorm:"index; not null"`
	Email       string      `json:"email"`
	Geoposicion Geoposicion `json:"geoposicion" gorm:"foreignKey:EmpresaId; not null; constraint:OnDelete:CASCADE"`
	Areas       []AreaType  `json:"areas" gorm:"foreignKey:EmpresaId; constraint:OnDelete:CASCADE"`
	CreatedAt   time.Time   `json:"created_at" gorm:"autoCreateTime"`
	Owner       string      `json:"owner" gorm:"not null"`
}

func (e *Empresa) Migrate() {
	db := rightPort.GetDB()
	m := db.Migrator()
	if m.HasTable(e) == true {
		return
	}
	m.CreateTable(e)
}

func (e *Empresa) List(filter []byte) error {
	t := new(Entity[Empresa])
	t.Data = filter
	err := t.GetAll().Preload("Contactos").Preload("Geoposicion").Limit(25).Find(&t.Entities).Error
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("ID", "Nombre", "Email", "Telefono", "Lugar", "Direccion", "Provincia", "Lat", "Lng")
	fmt.Printf("%s %v Items\n", color.MagentaString("[Results]:"), len(t.Entities))
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	for _, el := range t.Entities {
		var phone string
		var lugar string
		if len(el.Contactos) > 0 {
			phone = el.Contactos[0].Numero
			lugar = el.Contactos[0].Tipo
		}
		tbl.AddRow(el.Id, el.Nombre, el.Email, phone, lugar, el.Direccion, el.Provincia, el.Geoposicion.Lat, el.Geoposicion.Lng)
	}
	tbl.Print()

	return err
}
