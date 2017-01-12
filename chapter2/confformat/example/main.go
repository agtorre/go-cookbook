package main

import (
	"fmt"

	"github.com/agtorre/go-cookbook/chapter2/confformat"
)

const (
	exampleTOML = `name="Example1"
age=99
    `

	exampleJSON = `{"name":"Example2","age":98}`

	exampleYAML = `name: Example3
age: 97    
    `
)

func marshalAll() error {
	t := confformat.TOMLData{
		Name: "Name1",
		Age:  20,
	}

	j := confformat.JSONData{
		Name: "Name2",
		Age:  30,
	}

	y := confformat.YAMLData{
		Name: "Name3",
		Age:  40,
	}

	tomlRes, err := t.ToTOML()
	if err != nil {
		return err
	}

	fmt.Println("TOML Marshal =", tomlRes.String())

	jsonRes, err := j.ToJSON()
	if err != nil {
		return err
	}

	fmt.Println("JSON Marshal=", jsonRes.String())

	yamlRes, err := y.ToYAML()
	if err != nil {
		return err
	}

	fmt.Println("YAML Marshal =", yamlRes.String())
	return nil
}

func unmarshalAll() error {
	t := confformat.TOMLData{}
	j := confformat.JSONData{}
	y := confformat.YAMLData{}

	if _, err := t.Decode([]byte(exampleTOML)); err != nil {
		return err
	}
	fmt.Println("TOML Unmarshal =", t)

	if err := j.Decode([]byte(exampleJSON)); err != nil {
		return err
	}
	fmt.Println("JSON Unmarshal =", j)

	if err := y.Decode([]byte(exampleYAML)); err != nil {
		return err
	}
	fmt.Println("Yaml Unmarshal =", y)
	return nil
}

func main() {

	if err := marshalAll(); err != nil {
		panic(err)
	}

	if err := unmarshalAll(); err != nil {
		panic(err)
	}

	if err := confformat.OtherJSONExamples(); err != nil {
		panic(err)
	}

}
