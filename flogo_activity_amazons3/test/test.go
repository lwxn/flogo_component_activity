package test

import (
	"fmt"
	"os"
	"path/filepath"
)

func test() {
	fmt.Println("-------begin")
	dir,_ := os.Getwd()
	s3Location := "/img/1.jpg"
	dir = filepath.Join(dir,s3Location)
	_,err := os.Create(dir)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(dir)
}
