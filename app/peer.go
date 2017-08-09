package app

type Peer struct {
	Ip string `json:"ip"`
	Port int `json:"port"`
}

func (peer *Peer) getUrl() string{
	return "http" + peer.Ip + ":" + string(peer.Port)
}