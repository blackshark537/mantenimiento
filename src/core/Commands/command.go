package commands

import (
	"encoding/json"
	"errors"

	entities "github.com/blackshark537/mantenimiento/src/core/Entities"
	left "github.com/blackshark537/mantenimiento/src/core/Ports/Left"
	"github.com/blackshark537/mantenimiento/src/core/actions"
	"github.com/fatih/color"
)

func Exec(ofType int, payload []byte) error {
	switch ofType {
	case actions.ServerStart:
		return left.Serve(string(payload[:]))

		// Users
	case actions.ListUser:
		e := entities.User{}
		return e.List(payload)

	case actions.CreateUser:
		e := entities.Entity[entities.User]{}
		json.Unmarshal(payload, &e.Entity)
		return e.Create()

	case actions.ClearUsers:
		e := entities.Entity[entities.User]{}
		return e.Clear("users")

		// Empresas
	case actions.ListEmpresa:
		e := entities.Empresa{}
		return e.List(payload)

	case actions.CreateEmpresa:
		e := entities.Entity[entities.Empresa]{}
		json.Unmarshal(payload, &e.Entity)
		return e.Create()

	case actions.ClearEmpresas:
		e := entities.Entity[entities.Empresa]{}
		return e.Clear("empresas")

		// Contacto
	case actions.ListContacto:
		e := entities.Contacto{}
		return e.List(payload)

	case actions.CreateContacto:
		e := entities.Entity[entities.Contacto]{}
		json.Unmarshal(payload, &e.Entity)
		return e.Create()

	case actions.ClearContactos:
		e := entities.Entity[entities.Contacto]{}
		return e.Clear("contactos")

		// Geoposicion
	case actions.ListGeo:
		e := entities.Geoposicion{}
		return e.List(payload)

	case actions.CreateGeo:
		e := entities.Entity[entities.Geoposicion]{}
		json.Unmarshal(payload, &e.Entity)
		return e.Create()

	case actions.ClearGeo:
		e := entities.Entity[entities.Geoposicion]{}
		return e.Clear("geoposicions")

	// Area
	case actions.ListAreaType:
		e := entities.AreaType{}
		return e.List(payload)

	case actions.CreateAreaType:
		e := entities.Entity[entities.AreaType]{}
		json.Unmarshal(payload, &e.Entity)
		return e.Create()

	case actions.ClearAreaTypes:
		e := entities.Entity[entities.AreaType]{}
		return e.Clear("area_types")

		// Area
	case actions.ListArea:
		e := entities.Area{}
		return e.List(payload)

	case actions.CreateArea:
		e := entities.Entity[entities.Area]{}
		json.Unmarshal(payload, &e.Entity)
		return e.Create()

	case actions.ClearAreas:
		e := entities.Entity[entities.Area]{}
		return e.Clear("areas")

		// EquipoType
	case actions.ListEquipoType:
		e := entities.EquipoType{}
		return e.List(payload)

	case actions.CreateEquipoType:
		e := entities.Entity[entities.EquipoType]{}
		json.Unmarshal(payload, &e.Entity)
		return e.Create()

	case actions.ClearEquipoType:
		e := entities.Entity[entities.EquipoType]{}
		return e.Clear("equipo_types")

		// Equipo
	case actions.ListEquipo:
		e := entities.Equipo{}
		return e.List(payload)

	case actions.CreateEquipo:
		e := entities.Entity[entities.Equipo]{}
		json.Unmarshal(payload, &e.Entity)
		return e.Create()

	case actions.ClearEquipos:
		e := entities.Entity[entities.Equipo]{}
		return e.Clear("equipos")

		// ComponenteType
	case actions.ListComponenteType:
		e := entities.ComponenteType{}
		return e.List(payload)

	case actions.CreateComponenteType:
		e := entities.Entity[entities.ComponenteType]{}
		json.Unmarshal(payload, &e.Entity)
		return e.Create()

	case actions.ClearCompoenteTypes:
		e := entities.Entity[entities.ComponenteType]{}
		return e.Clear("componente_types")

		// Compoente
	case actions.ListComponente:
		e := entities.Componente{}
		return e.List(payload)

	case actions.CreateComponente:
		e := entities.Entity[entities.Componente]{}
		json.Unmarshal(payload, &e.Entity)
		return e.Create()

	case actions.ClearComponentes:
		e := entities.Entity[entities.Componente]{}
		return e.Clear("componentes")
	default:
		return errors.New(color.RedString("Error: Command not implemented"))
	}
}
