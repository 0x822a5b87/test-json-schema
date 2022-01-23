package main

import (
	"github.com/go-playground/validator/v10"
	"json/valid"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {
	file, err := os.ReadFile("./valid/data/example.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
		return
	}
	log.Printf("%s\n", string(file))
	flinkJob := valid.FlinkJob{}
	err = yaml.Unmarshal([]byte(string(file)), &flinkJob)
	if err != nil {
		log.Fatalf("error: %v", err)
		return
	}
	validate := validator.New()
	err = validate.Struct(flinkJob)
	if err != nil {
		log.Fatalf("error: %v", err)
		return
	}
}
