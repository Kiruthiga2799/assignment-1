package main

import (
	"fmt"
	"sync"
)

func main() {
	employee:= []emp{
		emp{Name: "samuel"},
		emp{Name: "kamal"},
        emp{Name: "mani"},
        emp{Name: "sharan"},
        emp{Name: "brathesh"},
        emp{Name: "ranjith"},
}
	result := Registeremployee(employee, stream{Name: "IT"})
	fmt.Print(result)
}

type RegisteremployeeResults struct {
	Results []empRegistrationResult
}

type empRegistrationResult struct {
	Registration empRegistration
	Error        error
}

type empRegistration struct {
	emp emp
	stream stream
}

type stream struct {
	Name string
}

type emp struct {
	Name string
}

func Registeremployee(employee []emp, stream stream) RegisteremployeeResults {
	output := make(chan RegisteremployeeResults)
	input := make(chan empRegistrationResult)
	var wg sync.WaitGroup
	go handleResults(input, output, &wg)
	defer close(output)
	for _, emp := range employee {
		wg.Add(1)
		go ConcurrentRegisteremp(emp, stream, input)
	}

	wg.Wait()
	close(input)
	return <-output
}

func handleResults(input chan empRegistrationResult, output chan RegisteremployeeResults, wg *sync.WaitGroup) {
	var results RegisteremployeeResults
	for result := range input {
		results.Results = append(results.Results, result)
		wg.Done()
	}
	output <- results
}

func ConcurrentRegisteremp(emp emp, stream stream, output chan empRegistrationResult) {
	result := Registeremp(emp, stream)
	output <- result
}

func Registeremp(emp emp, stream stream) empRegistrationResult {
	return empRegistrationResult{
		Registration: empRegistration{
			emp: emp,
			stream:  stream,
		},
	}
}
