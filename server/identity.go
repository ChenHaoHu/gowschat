package server

import (
	"errors"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

type Member struct {
	Uid       string
	Name      string
	Conn      *websocket.Conn
	LoginTime string
}

func ParseToken(token string) (*Member, string, error) {

	strs := strings.Split(token, ",")

	if len(strs) == 2 {

		return &Member{
			Uid:       strs[0],
			Name:      strs[1],
			LoginTime: time.Now().Format("Jan 02, 2006 15:04:05 UTC"),
		}, "ALL", nil
	}

	if len(strs) == 3 {

		return &Member{
			Uid:       strs[0],
			Name:      strs[1],
			LoginTime: time.Now().Format("Jan 02, 2006 15:04:05 UTC"),
		}, strs[2], nil
	}

	return nil, "", errors.New("token error")
}

func CheckInden(request *RequestEntity, member *Member, gid string) error {

	//return errors.New("you do not have power")

	/**

	const (
		P2P int = iota //somone msg to notify somone 0
		P2G            // somone msg to notify group 1

		N2G            // msg to notify all 2
		N2A            // msg to notify all 3
		N2P            // msg to notify someone 4
	)
	**/

	if request.SendType <= P2G {
		return nil
	}

	return errors.New("you do not have power ")
}
