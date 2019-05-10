package room

type Player struct {
	UserId int
	Name   string
}

func (p *Player) Init() {

}

func NewPlayer(userId int) Player {
	p := Player{UserId: userId}
	p.Init()
	return p
}
