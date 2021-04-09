package glc

import (
	"strings"

	"glc-infinite/strutil"
)

// Production represents a production
type Production struct {
	Variable   string `json:"variable"`
	Production string `json:"production"`
}

// Contains true if the production contains a variable, false otherwise.
func (p Production) Contains(variables []string) bool {
	for _, variable := range variables {
		if strings.Contains(p.Production, variable) {
			return true
		}
	}
	return false
}

// ExtractVariables extract variables from a production.
func (p Production) ExtractVariables(variables []string) []string {
	var extractedVariables []string
	for _, variable := range variables {
		i := strings.IndexAny(p.Production, variable)
		if i != -1 {
			extractedVariables = append(extractedVariables, variable)
		}
	}
	return strutil.RemoveDups(extractedVariables)
}
