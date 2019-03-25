package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "io/ioutil"
)
func help() {
  fmt.Println("\tend - exit from program")
  fmt.Println("\tread - read from file")
  fmt.Println("\twrite - write to file")
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func readFromFile() {
  dat, err := ioutil.ReadFile("./assets/someData")
  check(err)
  fmt.Print(string(dat))
}

func writeToFile() {
  f, err := os.Create("./assets/someData")
  check(err)

  defer f.Close()
  w := bufio.NewWriter(f)
  fmt.Print("Enter text to write: ")
  reader := bufio.NewReader(os.Stdin)
  text, _ := reader.ReadString('\n')
  w.WriteString(text)
  w.Flush()
}

func readCommand() {
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Enter command: ")
  text, _ := reader.ReadString('\n')
  switch strings.TrimRight(text, "\n") {
  case "read":
        readFromFile()
    case "help":
        help()
    case "write":
        writeToFile()
    case "end":
        os.Exit(1)
    }
    readCommand()
}

func main() {
  readCommand()
}
