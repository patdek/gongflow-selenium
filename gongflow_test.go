package main

import (
	"errors"
	"testing"

	"bitbucket.org/tebeka/selenium"
)

var (
	wd  selenium.WebDriver
	err error
)

func TestCreatingWebDriver(t *testing.T) {
	caps := selenium.Capabilities{"browserName": "firefox"}
	wd, err = selenium.NewRemote(caps, "")
	if err != nil {
		t.Fatal("Unable to intialize wd.  ", err)
	}
}

func TestOpeningDemo(t *testing.T) {
	// XXX: This crashes when given a bad domain.
	// TODO: Report bug to author or write patch to fix.
	err = wd.Get("http://localhost:8080")
	if err != nil {
		t.Error("Unable to resolve page.  ", err)
	}
}

func TestButtonUpload(t *testing.T) {
	err = findAndClickButton("span.btn:nth-child(2)", "Button Upload; ")
	if err != nil {
		t.Error("Unable to locate button. ", err)
	}
}

func TestInputsBrowse(t *testing.T) {
	err = findAndClickButton("div.span6:nth-child(1) > input:nth-child(2)", "Inputs Browse; ")
	if err != nil {
		t.Error("Unable to locate button. ", err)
	}
}

func TestQuittingWebDriver(t *testing.T) {
	err := wd.Quit()
	if err != nil {
		t.Error("Expected clean exit, got: ", err)
	}
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Below here are helper functions
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func findAndClickButton(fieldID, errID string) error {
	elem, err := wd.FindElement(selenium.ByCSSSelector, fieldID)
	if err != nil {
		return errors.New("Unable to find button: " + errID + " " + err.Error())
	}
	err = elem.Click()
	if err != nil {
		return errors.New("Unable to click button: " + errID + " " + err.Error())
	}
	return nil
}
