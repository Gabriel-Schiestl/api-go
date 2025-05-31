package pkg

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type UseCaseDecorator struct {
	useCase UseCase
}

func NewUseCaseDecorator(useCase UseCase) UseCaseDecorator {
	return UseCaseDecorator{useCase: useCase}
}

func (d UseCaseDecorator) Execute(props any) (any, error) {
	if d.useCase == nil {
		return nil, nil
	}

	useCaseType := reflect.TypeOf(d.useCase)
    useCaseName := useCaseType.String()
    
    if useCaseType.Kind() == reflect.Ptr {
        useCaseName = useCaseType.Elem().String()
    }

	propsType := "nil"
    if props != nil {
        propsType = reflect.TypeOf(props).String()
    }

	Logger.Debug().Str("useCase", useCaseName).Str("props", propsType).Msg("Executing use case")

	result, err := d.useCase.Execute()
	if err != nil {
        Logger.Error().Err(err).Str("useCase", useCaseName).Msg("Error executing use case")
        return nil, err
    }

	resultValue := "nil"
    if result != nil {
        jsonData, jsonErr := json.Marshal(result)
        if jsonErr == nil {
            resultValue = string(jsonData)
        } else {
            resultValue = fmt.Sprintf("%+v", result)
        }
    }

    Logger.Info().
        Str("useCase", useCaseName).
        Str("result", resultValue).
        Msg("Successfully executed use case")
        
    return result, nil
}