package main

import(
  "fmt"
  "os"
  "strings"

  "playground_askii-art/ascii"
)
func main(){
  if len(os.Args) != 2{
    fmt.Println("ATTENTION: expecting>> go run . [string]")
    os.Exit(0)
  }
  input:= os.Args[1]
  if input == ""{
    os.Exit(0)
  }
  input = strings.ReplaceAll(input, `\n`, "\n")
  if input == "\n"{
    fmt.Print("\n")
  }

  charMap, err:= ascii.LoadBanner("banner/standard.txt")
  if err != nil{
    fmt.Fprintf(os.Stderr, "ATTENTION: failed to load banner file %v\n", err)
    os.Exit(1)
  }
  output := ascii.LoadRender(input, charMap)
  fmt.Println(output)
  os.Exit(0)
}