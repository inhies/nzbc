package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/inhies/newznab/client"
)

var Conf Config

func init() {
	conf_str, err := ioutil.ReadFile("config.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(conf_str, &Conf)

	if err != nil {
		panic(err)
	}

}

func output() {
	output, err := json.MarshalIndent(Conf, "", "\t")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(output))
	return
}

type Config struct {
	// A list of newznab indexers
	Indexers []client.Indexer
}
