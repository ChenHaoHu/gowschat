package server

import (
	"errors"
	"strings"
	"time"
)

func ParseToken(token string) (*Member, error) {

	strs := strings.Split(token, ",")

	if len(strs) == 2 {

		return &Member{
			Uid:       strs[0],
			Name:      strs[1],
			Gid:       "ALL",
			LoginTime: time.Now().Format("Jan 02, 2006 15:04:05 UTC"),
		}, nil
	}

	if len(strs) == 3 {

		return &Member{
			Uid:       strs[0],
			Name:      strs[1],
			Gid:       strs[2],
			LoginTime: time.Now().Format("Jan 02, 2006 15:04:05 UTC"),
		}, nil
	}

	return nil, errors.New("token error")
}

func CheckInden(request *RequestEntity, member *Member) error {

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
