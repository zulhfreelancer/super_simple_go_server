package main

import (
  "fmt"
  "net/http"
  "os"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
  message := "Hello!"
  w.Write([]byte(message))
}

func main() {
  port_ := os.Getenv("PORT")
  if port_ == "" {
    fmt.Println("Please supply PORT environment variable")
    os.Exit(1)
  }
  port  := fmt.Sprintf(":%s", port_)

  http.HandleFunc("/", sayHello)
  fmt.Println("Server is running on port", port_)
  if err := http.ListenAndServe(port, nil); err != nil {
    panic(err)
  }
}
