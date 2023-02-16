package core

import (
	Commands "github.com/blackshark537/mantenimiento/src/core/Commands"
	entities "github.com/blackshark537/mantenimiento/src/core/Entities"
	left_port "github.com/blackshark537/mantenimiento/src/core/Ports/Left"
	right_port "github.com/blackshark537/mantenimiento/src/core/Ports/Right"
	Query "github.com/blackshark537/mantenimiento/src/core/Query"
	"gorm.io/gorm"
)

type API interface {
	ForRoot(port string) error
}

type Action struct {
	OfType  int
	Payload []byte
}

func InjectDatabase(db *gorm.DB) {
	right_port.InjectDB(db)
}

func InjectApi(api API) {
	left_port.InjectApi(api)
}

func (a *Action) Exec() error {
	return Commands.Exec(a.OfType, a.Payload)
}

func (a *Action) Query() ([]interface{}, error) {
	return Query.Exec(a.OfType, a.Payload)
}

func Migrate() {
	usr := entities.User{}
	usr.Migrate()

	emp := entities.Empresa{}
	emp.Migrate()

	geo := entities.Geoposicion{}
	geo.Migrate()

	cto := entities.Contacto{}
	cto.Migrate()

	area := entities.Area{}
	area.Migrate()

	areaType := entities.AreaType{}
	areaType.Migrate()

	eqp := entities.Equipo{}
	eqp.Migrate()

	eqpType := entities.EquipoType{}
	eqpType.Migrate()

	cmp := entities.Componente{}
	cmp.Migrate()

	cmpType := entities.ComponenteType{}
	cmpType.Migrate()

	alm := entities.Alimento{}
	alm.Migrate()

	ing := entities.Ingrediente{}
	ing.Migrate()

	spl := entities.Suplidor{}
	spl.Migrate()
}
