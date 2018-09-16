package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// BaseController -- controller type
type BaseController struct {
	name string
}

func (c *BaseController) toString(body io.Reader) string {
	contents, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	return string(contents)
}

func (c *BaseController) toStruct(body io.Reader, obj interface{}) {
	jsonString := c.toString(body)
	err := json.Unmarshal([]byte(jsonString), &obj)
	if err != nil {
		panic(err)
	}
	return
}
