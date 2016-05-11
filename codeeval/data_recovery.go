package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  // "strconv"
  "strings"
)

func main() {
  file, err := os.Open(os.Args[1])
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    //'scanner.Text()' represents the test case, do something with it
    s := strings.Split(scanner.Text(), ";")
    words, ordinals := strings.Fields(s[0]), strings.Fields(s[1])

    sentence := make(map[int]string)

    for _, ordinal := range ordinals {
      fmt.Println(ordinal)
      // num := strconv.Atoi(ordinal)
      // sentence[num] = "hello"
    }

    fmt.Println(words)
    fmt.Println(ordinals)
    fmt.Println(sentence)
  }
}
