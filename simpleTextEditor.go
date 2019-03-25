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
  fmt.Println("\tclean - clean the file data")
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func readFromFile() string {
  dat, err := ioutil.ReadFile("./assets/someData")
  check(err)
  // fmt.Print(string(dat))
  return string(dat)
}

func cleanFile() {
  f, err := os.Create("./assets/someData")
  check(err)

  defer f.Close()
}

func writeToFile() {
  oldText := readFromFile();
  f, err := os.Create("./assets/someData")
  check(err)

  defer f.Close()
  w := bufio.NewWriter(f)
  fmt.Print("Enter text to write: ")
  reader := bufio.NewReader(os.Stdin)
  text, _ := reader.ReadString('\n')
  w.WriteString(oldText + text)
  w.Flush()
}

func readCommand() {
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Enter command: ")
  text, _ := reader.ReadString('\n')
  switch strings.TrimRight(text, "\n") {
  case "read":
        fmt.Print(readFromFile())
    case "help":
        help()
    case "write":
        writeToFile()
    case "clean":
        cleanFile()
    case "end":
        os.Exit(1)
    }
    readCommand()
}

func main() {
  readCommand()
}
