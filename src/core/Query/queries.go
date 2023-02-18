package query

import (
	"errors"

	entities "github.com/blackshark537/mantenimiento/src/core/Entities"
	"github.com/blackshark537/mantenimiento/src/core/actions"
)

func Exec(ofType int, payload []byte) (interface{}, error) {
	return getAll(ofType, payload)
}

func getAll(ofType int, payload []byte) (interface{}, error) {
	switch ofType {
	case actions.GetAllAreas:
		e := entities.Entity[entities.Area]{}
		e.Data = payload
		err := e.GetAll().Preload("Equipos").Limit(25).Find(&e.Entities).Error
		return e.Entities, err

	case actions.GetAllAreaType:
		e := entities.Entity[entities.AreaType]{}
		e.Data = payload
		err := e.GetAll().Preload("Areas").Limit(25).Find(&e.Entities).Error
		return e.Entities, err

	case actions.GetAllEquipos:
		e := entities.Entity[entities.Equipo]{}
		e.Data = payload
		err := e.GetAll().Preload("Componentes").Limit(25).Find(&e.Entities).Error
		return e.Entities, err

	case actions.GetAllEquipoType:
		e := entities.Entity[entities.EquipoType]{}
		e.Data = payload
		err := e.GetAll().Preload("Equipos").Limit(25).Find(&e.Entities).Error
		return e.Entities, err

	case actions.GetAllComponentes:
		e := entities.Entity[entities.Componente]{}
		e.Data = payload
		err := e.GetAll().Limit(25).Find(&e.Entities).Error
		return e.Entities, err

	case actions.GetAllComponentesType:
		e := entities.Entity[entities.ComponenteType]{}
		e.Data = payload
		err := e.GetAll().Limit(25).Find(&e.Entities).Error
		return e.Entities, err

	default:
		return nil, errors.New("Error: Query no match!")
	}
}

func getOne(ofType int, payload []byte) (interface{}, error) {
	return nil, nil
}
