package server

var EnableNotifyAllWhenLogin bool = true
var EnableNotifyAllWhenLogout bool = true

const (
	P2P int = iota //somone msg to notify somone 0
	P2G            // somone msg to notify group 1
	N2G            // msg to notify all 2
	N2A            // msg to notify all 3
	N2P            // msg to notify someone 4
)
