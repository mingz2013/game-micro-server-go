package tcp

type TCPPacket struct {
	size int
	data []byte
}

func (p *TCPPacket) ToBytes() *[]byte {
	return nil
}
