package playwright

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/playwright-community/playwright-go"
)

var tests = []func(t *testing.T){
	TestPlaywright,
	TestPlaywright2,
}

func TestPlaywright(t *testing.T) {
	var pw Playwright
	headless := true
	opts := playwright.BrowserTypeLaunchOptions {
		Headless: &headless,
	}
	var opts2 playwright.PageGotoOptions
	var opts3 playwright.PageWaitForSelectorOptions

	pw.Launch(opts)
	pw.NewPage()
	pw.Goto("https://www.google.com", opts2)
	pw.WaitForSelector("input[title='Search']", opts3)
	pw.Kill()
}

func TestPlaywright2(t *testing.T) {
	var pw Playwright
	headless := true
	opts := playwright.BrowserTypeLaunchOptions {
		Headless: &headless,
	}
	var opts2 playwright.PageGotoOptions
	var opts3 playwright.PageWaitForSelectorOptions
	var opts4 playwright.PageTypeOptions

	pw.Launch(opts)
	pw.NewPage()
	pw.Goto("https://www.google.com", opts2)
	pw.WaitForSelector("input[title='Search']", opts3)
	pw.Type("input[title='Search']", "how to measure real user metrics with the xk6-playwright extension for k6?", opts4)
	fp := pw.FirstPaint()
	fcp := pw.FirstContentfulPaint()
	ttmi := pw.TimeToMinimallyInteractive()
	fid := pw.FirstInputDelay()
	fmt.Printf("First Paint: %v \n", fp)
	fmt.Printf("First Contentful Paint: %v \n", fcp)
	fmt.Printf("Time to Minimally Interactive: %v \n", ttmi)
	fmt.Printf("First Input Delay: %v \n", fid)
	pw.Kill()
}

func TestPlaywright3(t *testing.T) {
	var pw Playwright
	opts := playwright.BrowserTypeConnectOverCDPOptions {
	}
	var opts2 playwright.PageGotoOptions

	pw.Connect("http://localhost:9222", opts)
	pw.Goto("https://www.unqork.com", opts2)
}

func TestAll(t *testing.T) {
	for i, test := range tests {
		t.Run(strconv.Itoa(i), test)
	}
}
