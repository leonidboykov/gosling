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
//              if err := gosling.BuildRedirect(os.Stdout, "https://example.com"); err != nil {
//                      fmt.Println(err)
//              }
//      }
package gosling
