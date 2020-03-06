package main

import (
	"fmt"
	"github.com/tebeka/selenium"
	"net"
	"os"
	"os/exec"
	"strconv"
	// "time"
	// "strings"
)

// Used snippet from https://github.com/tebeka/selenium/issues/103

// ChromeTest is an example of a basic selenium headless call through Goolge Chrome to make a headless instance 
func main() {
	// browserPath := GetBrowserPath("chromium")
	port, err := pickUnusedPort()

	var opts []selenium.ServiceOption
	service, err := selenium.NewChromeDriverService("./chromedriver",
		port, opts...)

	if err != nil {
		fmt.Printf("Error starting the ChromeDriver server: %v", err)
	}

	caps := selenium.Capabilities{
		"browserName": "chrome",
	}

	wd, err := selenium.NewRemote(caps, "http://127.0.0.1:"+strconv.Itoa(port)+"/wd/hub")
	if err != nil {
		panic(err)
	}

	wd.Refresh()

	wd.Get("google.com")
		defer service.Stop()


	// From here down in an example of a query to snopes.com
	// wd.Get("snopes.com")

	// // Get a reference to the search bar
	// elem, err := wd.FindElement(selenium.ByCSSSelector, "#header-search")
	// if err != nil {
	// 	panic(err)
	// }

	// // Remove the boilerplate "Search Snopes.com" already in the text box.
	// if err := elem.Clear(); err != nil {
	// 	panic(err)
	// }

	// // Enter the user's query, soon to be from Slack
	// err = elem.SendKeys("Did Trump Fire the US Pandemic Response Team?")
	// if err != nil {
	// 	panic(err)
	// }

	// // Click the run button.
	// btn, err := wd.FindElement(selenium.ByCSSSelector, "#run")
	// if err != nil {
	// 	panic(err)
	// }
	// if err := btn.Click(); err != nil {
	// 	panic(err)
	// }

	// // Wait for the program to finish running and get the output.
	// // This simply returns if the query is true or false. The selector for the explanation itself is:
	// // body > div.theme-content > div > div > main > article > div.content-wrapper.card > div.content
	// outputDiv, err := wd.FindElement(selenium.ByCSSSelector, "body > div.theme-content > div > div > main > article > div.rating-wrapper.card > div > div > div > h5")
	// if err != nil {
	// 	panic(err)
	// }

	// var output string
	// for {
	// 	output, err = outputDiv.Text()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	if output != "Waiting for remote server..." {
	// 		break
	// 	}
	// 	time.Sleep(time.Millisecond * 100)
	// }

	// fmt.Printf("%s", strings.Replace(output, "\n\n", "\n", -1))
}

func pickUnusedPort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	port := l.Addr().(*net.TCPAddr).Port
	if err := l.Close(); err != nil {
		return 0, err
	}
	return port, nil
}

// GetBrowserPath Don't need this, just for testing
func GetBrowserPath(browser string) string {
	if _, err := os.Stat(browser); err != nil {
		path, err := exec.LookPath(browser)
		if err != nil {
			panic("Browser binary path not found")
		}
		return path
	}
	return browser
}