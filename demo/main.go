package main

import (
	"encoding/json"
	"fmt"
	"github.com/lab-go/go-plugin-std/runtime"
)

func Builder() runtime.PluginFactory {
	return &DemoFactory{}
}

type DemoFactory struct {
}

func (d *DemoFactory) Create(conf []byte) (runtime.Filter, interface{}) {
	c := &Config{}
	json.Unmarshal(conf, c)
	return &DemoFilter{conf: c}, &Config{Name: "name", Type: "type"}
}

type Config struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type DemoFilter struct {
	conf *Config
}

func (d *DemoFilter) Execute(ctx runtime.Context) error {
	fmt.Println("DemoFilter start")
	fmt.Println(ctx.Name)
	fmt.Println(d.conf.Name)
	fmt.Println("DemoFilter end")
	return nil
}
