package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

// func test() int {
// 	x := 5
// 	defer func() {
// 		x += 1
// 	}()
// 	return x
// }

// func test() (x int) {
// 	defer func() {
// 		x += 1
// 	}()
// 	return 5
// }

// func test() (y int) {
// 	x := 5
// 	defer func() {
// 		x += 1
// 	}()
// 	return x
// }

// func test() (x int) {
// 	defer func(x int) {
// 		x += 1
// 	}(x)
// 	return 5
// }

func ListDir(dirPath string, deep int) (err error) {
	dir, err := ioutil.ReadDir(dirPath) //ReadDir内置函数列出当前所有文件夹
	if err != nil {
		return err
	}
	if deep == 1 {
		fmt.Printf("!---%s\n", filepath.Base(dirPath))
	}
	sep := string(os.PathSeparator) //os.PathSeparator是目录的分隔符，win下是\，linux下是/
	for _, fi := range dir {
		if fi.IsDir() {
			fmt.Printf("|")
			for i := 0; i < deep; i++ {
				fmt.Printf("    |")
			}
			fmt.Printf("----%s\n", fi.Name())
			ListDir(dirPath+sep+fi.Name(), deep+1)
			continue
		}
		fmt.Printf("|")
		for i := 0; i < deep; i++ {
			fmt.Printf("    |")
		}
		fmt.Printf("----%s\n", fi.Name())
	}
	return nil
}
func test() {
	app := cli.NewApp()
	app.Name = "tree"
	app.Usage = "list all file"
	app.Action = func(c *cli.Context) error {
		var dir string = "."
		if c.NArg() > 0 {
			dir = c.Args()[0]
		}
		ListDir(dir, 1)
		return nil
	}
	app.Run(os.Args)
}

func main() {
	// fmt.Println(test())
	test()
}
