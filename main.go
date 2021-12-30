package main

import (
	"fmt"
	"io/ioutil"
	"os"
  "encoding/json"
)

func main() {
	readFile()
  goTypeToJsonString()
  jsonStringToGoType()
}

func readFile() {
  jsonFile, err := os.Open("file.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

  byteValue, _ := ioutil.ReadAll(jsonFile)

  fmt.Printf("%v\n", string(byteValue))
  // [
  //   {
  //     "title": "Title 1",
  //     "url": "Url 1",
  //     "channel": {
  //       "name": "Channel Name 1",
  //       "url": "Channel Url 1"
  //     }
  //   },
  //   {
  //     "title": "Title 2",
  //     "url": "Url 2",
  //     "channel": {
  //     "name": "Channel Name 2",
  //     "url": "Channel Url 2"
  //     }
  //   },
  //   {
  //   "title": "Title 3",
  //   "url": "Url 3",
  //   "channel": {
  //     "name": "Channel Name 3",
  //       "url": "Channel Url 3"
  //     }
  //   }
  // ]
  fmt.Printf("%T\n", string(byteValue))
  // string
}

func jsonStringToGoType() {
  jsonString := []byte(`{
    "title": "Title 1",
    "url": "Url 1",
    "channel": {
      "name": "Channel Name 1",
      "url": "Channel Url 1"
    }
  }`)

  type channel struct {
    Name string
    URL string
  }

  type myType struct {
    Title string
    URL string
    Channel channel
  }

  var goType myType

  err := json.Unmarshal(jsonString, &goType)
  if err != nil {
    fmt.Println(err)
  }

  fmt.Printf("%+v\n", goType)
  // {Title:Title 1 URL:Url 1 Channel:{Name:Channel Name 1 URL:Channel Url 1}}
  fmt.Printf("%T\n", goType)
  // main.myType
}

func goTypeToJsonString() {
  
  type channel struct {
    Name string
    URL string
  }

  type myType struct {
    Title string
    URL string
    Channel channel
  }

  data := myType{"Title 1", "Url 1", channel{"Channel Name 1", "Channel Url 1"}}

  jsonData, err := json.Marshal(data)
  if err != nil {
    fmt.Println(err)
  }

  fmt.Println(string(jsonData))
  // {"Title":"Title 1","URL":"Url 1","Channel":{"Name":"Channel Name 1","URL":"Channel Url 1"}}
  fmt.Printf("%T\n", string(jsonData))
  // string
}
