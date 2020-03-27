package server

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type TokenStruct struct {
	Status int
	Gid    string
	Name   string
	Uid    string
}

func ParseToken(token string) (*Member, error) {

	if TokenUrl == "" {
		return defaultParseToken(token)
	}

	client := &http.Client{}

	url := TokenUrl + "?token=" + token
	reqest, err := http.NewRequest("GET", url, nil)

	if err != nil {
		// panic(err)
		return nil, err
	}

	response, _ := client.Do(reqest)

	//返回的状态码
	status := response.StatusCode

	defer response.Body.Close()

	if status == 200 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err

		}
		data := &TokenStruct{}
		if err := json.Unmarshal(body, &data); err != nil {
			return nil, err
		}
		log.Printf("%+v", data)

		if data.Status != 0 {
			return nil, errors.New("token 不合法\\没权限")
		}

		return &Member{
			Uid:       data.Uid,
			Name:      data.Name,
			Gid:       data.Gid,
			LoginTime: time.Now().Format("Jan 02, 2006 15:04:05 UTC"),
		}, nil

	} else {
		return nil, errors.New("the StatusCode is " + strconv.Itoa(status) + " url is: " + url)
	}

	//return nil, errors.New("token error")
}

func defaultParseToken(token string) (*Member, error) {

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

	return nil, errors.New("token can not check pass")
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
