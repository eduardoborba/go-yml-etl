package logic

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type EtlRules struct {
	Rules         []EtlRule `yaml:"rules"`
	rulesFilename string
}

type EtlRule struct {
	Name        string         `yaml:"name"`
	Source      EtlSource      `yaml:"source"`
	Destination EtlDestination `yaml:"destination"`
}

type EtlSource struct {
	DBMS     string  `yaml:"dbms"`
	Host     string  `yaml:"host"`
	Port     int     `yaml:"port"`
	User     string  `yaml:"user"`
	Password *string `yaml:"password"`
	Data     EtlData `yaml:"data"`
}

type EtlDestination struct {
	DBMS     string  `yaml:"dbms"`
	Host     string  `yaml:"host"`
	Port     int     `yaml:"port"`
	User     string  `yaml:"user"`
	Password *string `yaml:"password"`
	Data     EtlData `yaml:"data"`
}

type EtlData struct {
	Database    string  `yaml:"database"`
	CustomQuery *string `yaml:"custom_query"`
}

func NewEtlRules(rulesFilename string) *EtlRules {
	etlRules := &EtlRules{
		rulesFilename: rulesFilename,
	}

	etlRules.Unmarshal()

	return etlRules
}

func (e *EtlRules) Unmarshal() {
	yfile, err := ioutil.ReadFile(e.rulesFilename)
	if err != nil {
		log.Fatal(err)
	}

	err2 := yaml.Unmarshal(yfile, e)

	if err2 != nil {
		log.Fatal(err2)
	}
}
