package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os/exec"
	"strings"
	"log"
	"net/http"
	"io/ioutil"
)

// Language-specific import patterns can be dynamically discovered via parsing
const port = ":8080"

// Entry point for the HTTP server
func main() {
	http.HandleFunc("/autoimport", autoImportHandler)
	log.Println("Starting server on", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

// HTTP handler for /autoimport endpoint
func autoImportHandler(w http.ResponseWriter, r *http.Request) {
	// Read the code from the POST request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Extract code and language from the query parameters
	code := string(body)
	language := r.URL.Query().Get("language")
	if language == "" {
		http.Error(w, "Language is required", http.StatusBadRequest)
		return
	}

	// Dynamically determine required imports based on the code and language
	var imports []string
	if language == "go" {
		imports = getGoImports(code)
	} else {
		imports = append(imports, "Error: Language not supported")
	}

	// Return the modified code with imports
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"imports":%v, "code":"%s"}`, imports, code)
}

// Dynamically detect Go imports based on the code
func getGoImports(code string) []string {
	var imports []string

	// Step 1: Parse the Go code into an AST (Abstract Syntax Tree)
	fs := token.NewFileSet()
	node, err := parser.ParseFile(fs, "main.go", code, parser.AllErrors)
	if err != nil {
		log.Println("Error parsing Go code:", err)
		return nil
	}

	// Step 2: Find all the function calls or methods used in the code
	functions := findFunctionCalls(node)

	// Step 3: For each function, search the Go modules system for the appropriate imports
	for _, function := range functions {
		// We query Go modules to find which package provides the function
		packages := findGoPackage(function)
		for _, pkg := range packages {
			// Suggest the package as an import if it hasn't been added already
			if !contains(imports, pkg) {
				imports = append(imports, pkg)
			}
		}
	}

	return imports
}

// findFunctionCalls analyzes the AST and returns a list of function calls
func findFunctionCalls(node ast.Node) []string {
	var functions []string
	ast.Inspect(node, func(n ast.Node) bool {
		// Look for function calls
		if call, ok := n.(*ast.CallExpr); ok {
			if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
				functions = append(functions, sel.Sel.Name)
			}
		}
		return true
	})
	return functions
}

// findGoPackage queries the Go modules to determine which packages provide the functions
func findGoPackage(function string) []string {
	var result []string

	// Search for packages that provide the function
	cmd := exec.Command("go", "list", "std", "-json")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Println("Error running 'go list':", err)
		return nil
	}

	// Parse the output of 'go list'
	packages := strings.Split(out.String(), "\n")
	for _, pkg := range packages {
		// If the function exists in the package, suggest it as an import
		if strings.Contains(pkg, function) {
			result = append(result, pkg)
		}
	}

	return result
}

// Helper function to check if a string is in a slice
func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
