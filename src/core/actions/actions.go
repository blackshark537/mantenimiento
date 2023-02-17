package actions

const (
	ServerStart = iota

	GetUser = iota
	ListUser
	CreateUser
	UpdateUser
	DeleteUser
	ClearUsers
	ChangePswd

	GetAllEmpresas = iota
	ListEmpresa
	CreateEmpresa
	UpdateEmpresa
	DeleteEmpresa
	ClearEmpresas

	GetAllContactos = iota
	GetContacto
	ListContacto
	CreateContacto
	UpdateContacto
	DeleteContacto
	ClearContactos

	GetAllGeo = iota
	GetGeo
	ListGeo
	CreateGeo
	UpdateGeo
	DeleteGeo
	ClearGeo

	GetAllAreaType = iota
	GetAreaType
	ListAreaType
	CreateAreaType
	UpdateAreaType
	DeleteAreaType
	ClearAreaTypes

	GetAllAreas = iota
	GetArea
	ListArea
	CreateArea
	UpdateArea
	DeleteArea
	ClearAreas

	GetAllEquipoType = iota
	GetEquipoType
	ListEquipoType
	CreateEquipoType
	UpdateEquipoType
	DeleteEquipoType
	ClearEquipoType

	GetAllEquipos = iota
	GetEquipo
	ListEquipo
	CreateEquipo
	UpdateEquipo
	DeleteEquipo
	ClearEquipos

	GetAllComponentesType = iota
	GetComponenteType
	ListComponenteType
	CreateComponenteType
	UpdateComponenteType
	DeleteComponenteType
	ClearCompoenteTypes

	GetAllComponentes = iota
	GetComponente
	ListComponente
	CreateComponente
	UpdateComponente
	DeleteComponente
	ClearComponentes
)
