package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Generator struct {
}

func NewGenerator() *Generator {
	return &Generator{}
}

// Parse takes an input string, extracts placeholders enclosed in `{{` and `}}`, and replaces them with values obtained by calling methods on the provided data structure.
// It returns the updated string and an error if any issues occur during parsing.
func (g *Generator) Parse(input string, data interface{}, args ...map[string]interface{}) (string, error) {
	parsed := input
	structValue := reflect.ValueOf(data)
	methodMap := make(map[string]reflect.Value)

	// Retrieve the type of the provided data structure and create a map of method names to their associated reflect.Value.
	structType := reflect.TypeOf(data)
	for i := 0; i < structType.NumMethod(); i++ {
		method := structType.Method(i)
		methodMap[method.Name] = structValue.MethodByName(method.Name)
	}
	//fmt.Println("structValue: ", structValue)
	//fmt.Println("structType: ", structType)
	//fmt.Println("kind", structValue.Kind())
	//fmt.Println("NMethod: ", structType.NumMethod())
	for {
		start := strings.Index(parsed, "{{")
		end := strings.Index(parsed, "}}")

		if start == -1 || end == -1 {
			break
		}
		// Extract the placeholder name enclosed in `{{` and `}}`.
		fieldName := parsed[start+2 : end]

		field := structValue.MethodByName(fieldName)

		// Check if the method exists and is a function.
		if field.IsValid() && field.Type().Kind() == reflect.Func {
			methodArgs := make([]reflect.Value, 0)
			// Check if an argument exists for the current method, and add it if found.
			if len(args) > 0 {
				if arg, argExists := args[0][fieldName]; argExists {
					methodArgs = append(methodArgs, reflect.ValueOf(arg))
				}
			}
			var result []reflect.Value
			if len(methodArgs) > 0 {
				result = field.Call(methodArgs)
			} else {
				result = field.Call(nil)
			}

			if len(result) > 0 {
				parsed = strings.Replace(parsed, "{{"+fieldName+"}}", result[0].Interface().(string), 1)
			}
		} else {
			return "", fmt.Errorf("method '%s' does not exist on the struct", fieldName)
		}
	}

	return parsed, nil
}

func (g *Generator) CallUserFuncArray(callback interface{}, args []interface{}) []interface{} {
	funcValue := reflect.ValueOf(callback)
	if funcValue.Kind() != reflect.Func {
		panic("Callback must be a function")
	}

	argValues := make([]reflect.Value, len(args))
	for i, arg := range args {
		argValues[i] = reflect.ValueOf(arg)
	}

	resultValues := funcValue.Call(argValues)

	results := make([]interface{}, len(resultValues))
	for i, resultValue := range resultValues {
		results[i] = resultValue.Interface()
	}

	return results
}
