package glc

import (
	"glc-infinite/strutil"
)

// GLC represents a free context grammar
type GLC struct {
	Variables       []string     `json:"variables"`
	Terminals       []string     `json:"terminals"`
	InitialVariable string       `json:"initialVariable"`
	Productions     []Production `json:"productions"`
}

// IsFinite Indica si la gramatica libre de contexto es finita o infinita
func (g GLC) IsFinite() bool {
	var expandedVariables []string
	return g.isFinite(g.InitialVariable, expandedVariables)
}

// IsFinite Indica si la gramatica libre de contexto es finita o infinita
func (g GLC) isFinite(variable string, expandedVariables []string) bool {
	// Validamos que variable exista en variables
	_, variableExist := strutil.Find(g.Variables, variable)
	if !variableExist {
		panic(variable + " not found in variables")
	}

	// Validamos las condiciones de parada
	_, alreadyExpanded := strutil.Find(expandedVariables, variable)
	if alreadyExpanded {
		return false
	}

	// Agregamos variable a expandedVariables
	expandedVariables = append(expandedVariables, variable)
	// fmt.Println(expandedVariables)

	// Expandimos cada production correspondiente a variable.
	finite := true
	for _, production := range g.Productions {
		// Validamos que variable de production exista en variables
		_, productionVariableExist := strutil.Find(g.Variables, production.Variable)
		if !productionVariableExist {
			panic(production.Variable + " not found in variables")
		}

		if production.Variable != variable {
			continue
		}

		// Buscamos alguna variable en la produccion.
		containsProduction := production.Contains(g.Variables)
		// Si contiene alguna produccion expandimos cada produccion.
		if containsProduction {
			extractedVariables := production.ExtractVariables(g.Variables)
			// fmt.Println(extractedVariables)
			for _, extractedVariable := range extractedVariables {
				finite = finite && g.isFinite(extractedVariable, expandedVariables)
			}
		} else {
			// Si no contiene una produccion, corresponde a un estado terminal.
			finite = finite && true
		}
	}

	return finite
}
