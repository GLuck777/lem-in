package functions

import "html/template"

type (
	Cells interface {
		IsConnectedTo(roomId int) bool
		ConnectTo(roomName string, Cells []Cells)
	}
	Colony interface{}
)

type Ants struct {
	Name string
	Cell string
	Path []string
}
type Room struct {
	N     int
	Start string
	End   string
}
type BasicRoom struct {
	Id       int
	Name     string
	Present  bool
	Adjacent []bool
	Status   string
	PosX     string
	PosY     string
}

func NewAnt(name string, cell string, path []string) Colony {
	Colony := &Ants{
		Name: name,
		Cell: cell,
		Path: path,
	}
	return Colony
}

func NewRoom(id int, name string, numRooms int, status, posX, posY string) Cells {
	Cell := &BasicRoom{
		Id:       id,
		Name:     name,
		Present:  false,
		Adjacent: make([]bool, numRooms),
		Status:   status,
		PosX:     posX,
		PosY:     posY,
	}
	return Cell
}

func (r *BasicRoom) IsConnectedTo(roomId int) bool {
	return r.Adjacent[roomId]
}

//	func (r *BasicRoom) ConnectTo(roomName string) {
//		r.adjacent = append(r.adjacent, true)
//	}
func (r *BasicRoom) ConnectTo(roomName string, Cells []Cells) {
	for i, Cell := range Cells {
		if Cell.(*BasicRoom).Name == roomName {
			r.Adjacent[i] = true
		}
	}
}

type LeminGraph struct {
	ParcoursHTML []string //`json:"parcours"`
	Soluce       []string //`json:"soluce"`
	// start        string
	// end          string
	// fourmis      int
	NBANTS    int
	StartRoom string
	EndRoom   string
	Connexion []string
}

type ImprimeGraph struct {
	ParcoursHTML []template.HTML //`json:"parcours"`
	Soluce       []string        //`json:"soluce"`
	Connexion    []string
	NBANTS       int
	StartRoom    string
	Title string
}
