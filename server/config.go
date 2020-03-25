package server

var EnableNotifyAllWhenLogin bool = true
var EnableNotifyAllWhenLogout bool = true

const (
	P2P int = iota //somone msg to notify somone
	P2A            // somone msg to notify all
	N2A            // msg to notify all
	N2P            // msg to notify someone
)
