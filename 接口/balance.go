package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Balancer interface {
	DoBalance([]*Instance) (*Instance, error)
}

type Instance struct {
	host string
	port int
}

func NewInstance(host string, port int) *Instance {
	return &Instance{
		host: host,
		port: port,
	}
}
func (p *Instance) GetHost() string {
	return p.host
}
func (p *Instance) GetPort() int {
	return p.port
}

func (p *Instance) String() string {
	return p.host + ":" + strconv.Itoa(p.port)
}

type RandomBalance struct {
}

func (p *RandomBalance) DoBalance(insts []*Instance) (inst *Instance, err error) {
	if len(insts) == 0 {
		err = errors.New("No instance")
		return
	}
	lens := len(insts)
	index := rand.Intn(lens)
	inst = insts[index]
	return
}

type RoundRobinBalance struct {
	curIndex int
}

func (p *RoundRobinBalance) DoBalance(insts []*Instance) (inst *Instance, err error) {
	if len(insts) == 0 {
		err = errors.New("No instance")
		return
	}
	lens := len(insts)
	if p.curIndex >= lens {
		p.curIndex = 0
	}
	inst = insts[p.curIndex]
	p.curIndex = (p.curIndex + 1) % lens
	return
}

func main() {
	var insts []*Instance
	for i := 0; i < 16; i++ {
		host := fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255))
		one := NewInstance(host, 8080)
		insts = append(insts, one)
	}
	var balancer Balancer
	var conf = "random"
	if len(os.Args) > 1 {
		conf = os.Args[1]
	}
	if conf == "random" {
		balancer = &RandomBalance{}
		fmt.Println("use random balancer")
	} else if conf == "roundranbin" {
		balancer = &RoundRobinBalance{}
		fmt.Println("use roundrobin balancer")
	}
	for {
		inst, err := balancer.DoBalance(insts)
		if err != nil {
			fmt.Println("do balanc err:", err)
			continue
		}
		fmt.Println(inst)
		time.Sleep(time.Second)
	}
}
