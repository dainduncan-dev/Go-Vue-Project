package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/yaml.v3"
)

type objects struct {
  Name string `yaml:"name"`
  Description string `yaml:"description"`
  Attributes []string `yaml:"attributes"`
}

type objectsJson struct {
  Name string `json:"name"`
  Description string `json:"description"`
  Attributes []string `json:"attributes"`
}

func getName(c echo.Context) error {
    fmt.Println("Parsing YAML File")

    var fileName string = `C:\Users\ddnkn\Desktop\taubyteGoVueAssessment\objects\object1.yaml`

    yamlFile, err := ioutil.ReadFile(fileName)
    if err != nil {
      fmt.Printf("Error reading from yaml file: %s\n", err)
    }

    var obj objects
    yaml.Unmarshal(yamlFile, &obj)

    var objJson = objectsJson{
      Name: obj.Name,
      Description: obj.Description,
      Attributes: obj.Attributes,
    } 

    jsonOutput, err := json.Marshal(objJson)
      c.JSON(http.StatusOK, objJson)
    err = ioutil.WriteFile(`C:\Users\ddnkn\Desktop\taubyteGoVueAssessment\objects\object1.json`, jsonOutput, 0644)

    return err
}

func main() {
    e := echo.New()

    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    e.Use(middleware.CORS())

    // Routes
    e.GET("/get", getName)

    e.Logger.Fatal(e.Start(":1323"))
}

