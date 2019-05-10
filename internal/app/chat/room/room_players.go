package room

type RoomPlayers struct {
	players []Player // 玩家队列
}

func (r *RoomPlayers) Init() {
	r.players = []Player{}
}

func (r *RoomPlayers) FindPlayerByUserId(userId int) (player Player, ok bool) {
	ok = false

	for i := 0; i < len(r.players); i++ {
		p := r.players[i]

		if p.UserId == userId {
			return p, true
		}

	}

	return
}

func (r *RoomPlayers) FindIndexByUserId(userId int) (index int, ok bool) {

	for i := 0; i < len(r.players); i++ {
		p := r.players[i]

		if p.UserId == userId {
			return i, true
		}

	}

	return -1, false
}

func (r *RoomPlayers) AddPlayer(userId int) (player Player, ok bool) {
	p, ok := r.FindPlayerByUserId(userId)
	if ok {
		ok = false
		return
	}
	r.players = append(r.players, NewPlayer(userId))
	ok = true
	player = p
	return
}

func (r *RoomPlayers) RmPlayer(userId int) (player Player, ok bool) {
	index, ok := r.FindIndexByUserId(userId)
	if !ok {
		return
	}
	player = r.players[index]

	players := r.players[:index]
	copy(players[len(players):], r.players[index:])
	r.players = players

	return
}
