package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

// packageJson will contain the output of go list -json
type packageJson struct {
	Name    string   `json:"Name"`
	Imports []string `json:"Imports"`
}

// ACDep contatins transitive dependencies of package A
// if package A depends on package B and package B depends on package C
// then we can see a dependency graph like this: A -> B -> C
// according to this graph, package A transitively depends on package C
// this struct will contain the name of the package A and list of package Cs
type ACDep struct {
	Name      string
	TransDeps []string
}

func contains(arr []string, keyword string) bool {

	for _, v := range arr {
		if v == keyword {
			return true
		}
	}
	return false
}

// containsAll checks if all members of smallArr are present in the bigArr
// order of the occurence doesn't matter
func containsAll(bigArr []string, smallArr []string) bool {

	if len(smallArr) > len(bigArr) {
		return false
	}
	for _, v := range smallArr {
		if !contains(bigArr, v) {
			return false
		}
	}
	return true
}
func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage : %s <package name(s)>\n", os.Args[0])
		os.Exit(1)
	}

	// listJson contains initial imports of the go packages in given directory
	listJson, err := exec.Command("go", "list", "-json").Output()

	if err != nil {
		panic(err)
	}

	// get first level imports
	var imports packageJson
	err = json.Unmarshal(listJson, &imports)

	if err != nil {
		panic(err)
	}
	var acdep ACDep
	acdep.Name = imports.Name
	var holderImports packageJson

	// iterate the same command for imports.Name for every imports.Imports
	// and add the result to the acdep.TransDeps
	for _, v := range imports.Imports {
		temp, _ := exec.Command("go", "list", "-json", v).Output()
		_ = json.Unmarshal(temp, &holderImports)
		acdep.TransDeps = append(acdep.TransDeps, holderImports.Imports...)
	}
	// compare acdep.TransDeps with given arguments and if all arguments exists
	// in the acdep.TransDeps then pront out the acdep.Name
	if containsAll(acdep.TransDeps, os.Args[1:]) {
		fmt.Printf("%s\n", acdep.Name)
	}
}
