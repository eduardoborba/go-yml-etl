package logic

import "fmt"

var exampleRulesFile = "etl-rule-example.yml"

type EtlProcessor struct {
}

func NewEtlProcessor() *EtlProcessor {
	return &EtlProcessor{}
}

func (p *EtlProcessor) Process() {
	etlRules := NewEtlRules(exampleRulesFile)

	for _, rule := range etlRules.Rules {
		fmt.Printf("rule name: %s\n", rule.Name)
		fmt.Printf("source query: %s\n", *rule.Source.Data.CustomQuery)
		fmt.Printf("destination query: %s\n", *rule.Destination.Data.CustomQuery)
	}
}
