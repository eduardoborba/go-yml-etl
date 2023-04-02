package main

import (
	"fmt"

	"github.com/eduardoborba/go-yml-etl/logic"
)

var exampleRulesFile = "etl-rule-example.yml"

func main() {
	etlRules := logic.NewEtlRules(exampleRulesFile)

	for _, rule := range etlRules.Rules {
		fmt.Printf("source query: %s\n", *rule.Source.Data.CustomQuery)
		fmt.Printf("destination query: %s\n", *rule.Destination.Data.CustomQuery)
	}
}
