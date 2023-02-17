package query

import (
	"encoding/json"
	"errors"

	entities "github.com/blackshark537/mantenimiento/src/core/Entities"
	"github.com/blackshark537/mantenimiento/src/core/actions"
)

func Exec(ofType int, payload []byte) (interface{}, error) {
	return getAll(ofType, payload)
}

func getAll(ofType int, payload []byte) (interface{}, error) {
	switch ofType {
	case actions.GetAllEmpresas:
		e := entities.Entity[entities.Empresa]{}
		json.Unmarshal(payload, &e.Entity)
		err := e.GetAll().Where(e.Entity).Preload("Contactos").Preload("Areas").Preload("Geoposicion").Limit(25).Find(&e.Entities).Error
		return e.Entities, err

	case actions.GetAllAreas:
		e := entities.Entity[entities.Area]{}
		json.Unmarshal(payload, &e.Entity)
		err := e.GetAll().Where(e.Entity).Preload("Equipos").Limit(25).Find(&e.Entities).Error
		return e.Entities, err

	case actions.GetAllAreaType:
		e := entities.Entity[entities.AreaType]{}
		json.Unmarshal(payload, &e.Entity)
		err := e.GetAll().Where(e.Entity).Preload("Areas").Limit(25).Find(&e.Entities).Error
		return e.Entities, err

	case actions.GetAllEquipos:
		e := entities.Entity[entities.Equipo]{}
		json.Unmarshal(payload, &e.Entity)
		err := e.GetAll().Where(e.Entity).Limit(25).Find(&e.Entities).Error
		return e.Entities, err

	case actions.GetAllEquipoType:
		e := entities.Entity[entities.EquipoType]{}
		json.Unmarshal(payload, &e.Entity)
		err := e.GetAll().Where(e.Entity).Limit(25).Find(&e.Entities).Error
		return e.Entities, err

	case actions.GetAllComponentes:
		e := entities.Entity[entities.Componente]{}
		json.Unmarshal(payload, &e.Entity)
		err := e.GetAll().Where(e.Entity).Limit(25).Find(&e.Entities).Error
		return e.Entities, err

	case actions.GetAllComponentesType:
		e := entities.Entity[entities.ComponenteType]{}
		json.Unmarshal(payload, &e.Entity)
		err := e.GetAll().Where(e.Entity).Limit(25).Find(&e.Entities).Error
		return e.Entities, err

	case actions.GetAllContactos:
		e := entities.Entity[entities.Contacto]{}
		json.Unmarshal(payload, &e.Entity)
		err := e.GetAll().Where(e.Entity).Limit(25).Find(&e.Entities).Error
		return e.Entities, err

	case actions.GetAllGeo:
		e := entities.Entity[entities.Geoposicion]{}
		json.Unmarshal(payload, &e.Entity)
		err := e.GetAll().Where(e.Entity).Limit(25).Find(&e.Entities).Error
		return e.Entities, err

	default:
		return nil, errors.New("Error: Query no match!")
	}
}

func getOne(ofType int, payload []byte) (interface{}, error) {
	return nil, nil
}
