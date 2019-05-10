package tcp

type Protocol struct {
}

func (p *Protocol) OnMsg(b []byte) {
	p.onMsg(b)
}

func (p *Protocol) onMsg(b []byte) {

}
