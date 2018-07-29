package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "os"
  "log"
)

type Tag struct {
  Layer string `json:"layer"`
  Name  string `json:"name"`
}

func main() {

  if len(os.Args) != 2 {
    log.Printf("Error. Please pass docker image name.")
    os.Exit(-1)
  }
  var image_name = os.Args[1]

  var url = "https://registry.hub.docker.com/v1/repositories"
  url = url + "/" + image_name + "/tags"

  log.Printf("Requesting to : %s", url)
  res, _ := http.Get(url)
  defer res.Body.Close()

  bytes, err := ioutil.ReadAll(res.Body)
  if err != nil {
    panic(err.Error())
  }

  var tags []Tag
  json.Unmarshal(bytes, &tags)

  for _, v := range tags {
    fmt.Printf("%s\n", v.Name)
  }

}
