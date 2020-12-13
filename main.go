package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

// GLC Representa una gramatica libre de contexto (GLC)
type GLC struct {
	Variables       []string     `json:"variables"`
	Terminals       []string     `json:"terminals"`
	InitialVariable string       `json:"initialVariable"`
	Productions     []Production `json:"productions"`
}

// Production Representa una produccion
type Production struct {
	Variable   string `json:"variable"`
	Production string `json:"production"`
}

// handleError Maneja errores
func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

// findString Toma un arreglo de string y busca un string en el. Si lo encuentra retorna su posicion y true, sino retorna -1 y false.
func findString(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func removeDuplicatedString(slice []string) []string {
	keys := make(map[string]bool)
	newSlice := []string{}
	for _, key := range slice {
		if _, value := keys[key]; !value {
			keys[key] = true
			newSlice = append(newSlice, key)
		}
	}
	return newSlice
}

// containsProduction true si la produccion contiene alguna variable, false sino.
func containsProduction(production string, variables []string) bool {
	for _, variable := range variables {
		if strings.Contains(production, variable) {
			return true
		}
	}
	return false
}

// extractVariables Extrae las variables que contiene una produccion
func extractVariables(production string, variables []string) []string {
	var extractedVariables []string
	for _, variable := range variables {
		i := strings.IndexAny(production, variable)
		if i != -1 {
			extractedVariables = append(extractedVariables, variable)
		}
	}
	return removeDuplicatedString(extractedVariables)
}

// isFinite Indica si la gramatica libre de contexto es finita o infinita
func isFinite(variable string, glc GLC, expandedVariables []string) bool {
	// Validamos que variable exista en variables
	_, variableExist := findString(glc.Variables, variable)
	if !variableExist {
		panic(variable + " not found in variables")
	}

	// Validamos las condiciones de parada
	_, alreadyExpanded := findString(expandedVariables, variable)
	if alreadyExpanded {
		return false
	}

	// Agregamos variable a expandedVariables
	expandedVariables = append(expandedVariables, variable)
	// fmt.Println(expandedVariables)

	// Expandimos cada production correspondiente a variable.
	finite := true
	for _, production := range glc.Productions {
		// Validamos que variable de production exista en variables
		_, productionVariableExist := findString(glc.Variables, production.Variable)
		if !productionVariableExist {
			panic(production.Variable + " not found in variables")
		}

		if production.Variable != variable {
			continue
		}

		// Buscamos alguna variable en la produccion.
		containsProduction := containsProduction(production.Production, glc.Variables)
		// Si contiene alguna produccion expandimos cada produccion.
		if containsProduction {
			extractedVariables := extractVariables(production.Production, glc.Variables)
			// fmt.Println(extractedVariables)
			for _, extractedVariable := range extractedVariables {
				finite = finite && isFinite(extractedVariable, glc, expandedVariables)
			}
		} else {
			// Si no contiene una produccion, corresponde a un estado terminal.
			finite = finite && true
		}
	}

	return finite
}

// main Funcion principal del programa.
func main() {
	// Ingrese el nombre del archivo.
	fmt.Println("Enter the file name:")
	var fileName string
	fmt.Scanf("%s", &fileName)

	// Leemos el archivo.
	body, err := ioutil.ReadFile(fileName)
	handleError(err)

	// Solo con proposito de pruebas...
	// fmt.Println(string(body))

	// Mapeamos el archivo en una estructura.
	glc := GLC{}
	err = json.Unmarshal([]byte(body), &glc)
	handleError(err)

	// Solo con proposito de pruebas...
	// fmt.Println(glc)

	// Validamos si el AFD es finito o infinito.
	var expandedVariables []string
	finite := isFinite(glc.InitialVariable, glc, expandedVariables)
	fmt.Println(finite)
}
