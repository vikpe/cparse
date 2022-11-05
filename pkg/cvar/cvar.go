package cvar

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/vikpe/cparse/pkg/comment"
)

const declarationPrefix = "cvar_t"

type Cvar struct {
	Name         string `json:"name"`
	DefaultValue string `json:"default_value"`
	SaveOnExit   string `json:"save_on_exit"`
	OnChange     string `json:"on_change"`
	Description  string `json:"description"`
}

func FromFile(path string) ([]Cvar, error) {
	cvars := make([]Cvar, 0)

	readFile, err := os.Open(path)
	defer readFile.Close()
	if err != nil {
		return cvars, err
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		cvariable, err := FromLine(string(fileScanner.Bytes()))

		if err == nil {
			cvars = append(cvars, cvariable)
		}
	}

	return cvars, nil
}

func FromLine(line string) (Cvar, error) {
	definitionBody, err := declarationCsvFromLine(line)
	if err != nil {
		return Cvar{}, err
	}

	result, err := fromCsv(definitionBody)
	if err != nil {
		return Cvar{}, err
	}

	result.Description = comment.FromSingleLine(line)
	return result, nil
}

func declarationCsvFromLine(line string) (string, error) {
	if !strings.HasPrefix(line, declarationPrefix) {
		errorMsg := fmt.Sprintf("missing declaration prefix (%s)", declarationPrefix)
		return "", errors.New(errorMsg)
	}

	indexOpen := strings.Index(line, "{")
	if -1 == indexOpen {
		return "", errors.New("missing open curly brace")
	}

	indexClose := strings.Index(line, "}")
	if -1 == indexClose {
		return "", errors.New("missing close curly brace")
	}

	csvLen := indexClose - indexOpen - 1
	if csvLen < 1 {
		return "", errors.New("empty definition")
	}

	return strings.TrimSpace(line[indexOpen+1 : indexClose]), nil
}

func fromCsv(csvStr string) (Cvar, error) {
	csvReader := csv.NewReader(strings.NewReader(csvStr))
	csvReader.TrimLeadingSpace = true
	csvReader.FieldsPerRecord = 1 + strings.Count(csvStr, ",")

	record, err := csvReader.Read()
	if err != nil {
		return Cvar{}, err
	}

	return fromRecord(record)
}

func fromRecord(record []string) (Cvar, error) {
	const IndexName = 0
	const IndexDefaultValue = 1
	const IndexSomething = 2
	const IndexOnChange = 3

	result := Cvar{}
	fieldCount := len(record)

	if fieldCount > IndexName {
		result.Name = record[IndexName]
	}

	if fieldCount > IndexDefaultValue {
		result.DefaultValue = record[IndexDefaultValue]
	}

	if fieldCount > IndexSomething {
		result.SaveOnExit = record[IndexSomething]
	}

	if fieldCount > IndexOnChange {
		result.OnChange = record[IndexOnChange]
	}

	return result, nil
}
