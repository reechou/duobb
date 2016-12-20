package main

import (
	"fmt"
	"github.com/absolute8511/gorpc"
	"github.com/reechou/duobb/service"
	"github.com/reechou/duobb_proto"
)

func main() {
	c := gorpc.NewTCPClient(":7879")
	c.Start()
	defer c.Stop()

	d := gorpc.NewDispatcher()
	d.AddService(duobb_proto.AccountService, new(service.TAccountService))

	dc := d.NewServiceClient(duobb_proto.AccountService, c)
	res, err := dc.Call("GetDuobbAccount", &duobb_proto.GetDuobbAccountReq{User: "reezhou"})
	fmt.Printf("Get=%+v, %+v\n", res, err)
}
