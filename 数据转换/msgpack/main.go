package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"

	"github.com/vmihailenco/msgpack"
)

type Person struct {
	Name string
	Age  int
	Sex  string
}

func writeMsgpack(filename string) (err error) {
	var person []*Person
	for i := 0; i < 10; i++ {
		p := &Person{
			Name: fmt.Sprintf("name%d", i),
			Age:  rand.Intn(50),
			Sex:  "man",
		}
		person = append(person, p)
	}
	data, err := msgpack.Marshal(person)
	if err != nil {
		fmt.Printf("marshal failed,err:%v\n", err)
		return
	}
	err = ioutil.WriteFile(filename, data, 0755)
	if err != nil {
		fmt.Printf("write file failed,err:%v\n", err)
		return
	}
	return
}

func readMsgpack(filename string) (err error) {
	var persons []*Person
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	err = msgpack.Unmarshal(data, &persons)
	if err != nil {
		return
	}
	return
}

func main() {
	filename := "D:/go视频/person.dat"
	err := writeMsgpack(filename)
	if err != nil {
		fmt.Printf("write msgpack failed,err:%\v\n", err)
		return
	}
	err = readMsgpack(filename)
	if err != nil {
		fmt.Printf("read msgpack failed,err:%v\n", err)
		return
	}
}
