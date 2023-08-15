// Package gosling builds html files with redirects commands for static sites
//
// Example:
//
//      package main
//
//      import (
//              "fmt"
//              "os"
//
//              "github.com/leonidboykov/gosling"
//      )
//
//      func main() {
//              if err := gosling.BuildRedirect("https://example.com", os.Stdout); err != nil {
//                      fmt.Println(err)
//              }
//      }
package gosling
