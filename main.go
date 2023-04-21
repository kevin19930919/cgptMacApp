package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"os/signal"

	"github.com/zserge/lorca"
	"gopkg.in/yaml.v2"

	"github.com/kevin19930919/cgptMacApp/chatgpt"
)

type Config struct {
	ChatGPTAPIKey string `yaml:"chatGPTAPIKey"`
}

func loadHTMLFile() (*string, error) {
	myFileContents, err := ioutil.ReadFile("./assets/index.html")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	loadableContents := "data:text/html," + url.PathEscape(string(myFileContents))
	return &loadableContents, nil
}

func main() {
	// Create the UI window
	ui, err := lorca.New("", "", 1440, 900, "--remote-allow-origins=*")
	if err != nil {
		fmt.Println("Error creating window:", err)
		return
	}
	defer ui.Close()

	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Parse the YAML data into the struct
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	c := &chatgpt.ChatGPTFunc{
		APIKEY: config.ChatGPTAPIKey,
	}
	// binding method
	ui.Bind("sendRequest", c.SendRequest)

	content, err := loadHTMLFile()
	if err != nil {
		return
	}
	// Set up the UI HTML content
	ui.Load("data:text/html," + *content)

	// Wait until the interrupt signal arrives or browser window is closed
	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
	case <-sigc:
	case <-ui.Done():
	}

	log.Println("exiting...")

}
