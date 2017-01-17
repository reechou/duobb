package main

import (
	"bytes"
	"fmt"
	"github.com/gorilla/rpc/json"
	"github.com/reechou/duobb_proto"
	"math/rand"
	"net/http"
	"strconv"
)

type Account struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

func main() {
	for i := 0; i < 100; i++ {
		a := &Account{
			User:     "neice" + strconv.Itoa(300+i),
			Password: strconv.Itoa(100000 + rand.Intn(799999)),
			Phone:    strconv.Itoa(20500 + i),
		}
		_, err := JsonRpcClient(a)
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		fmt.Println(a.User, " ", a.Password)
	}
}

func JsonRpcClient(request interface{}) (interface{}, error) {
	client := &http.Client{}
	url := "http://121.40.85.37:7878/rpc"
	message, err := json.EncodeClientRequest("DuobbAccountService.CreateDuobbAccount", request)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(message))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var result duobb_proto.Response
	err = json.DecodeClientResponse(resp.Body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
