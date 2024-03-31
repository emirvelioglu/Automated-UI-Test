package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/sclevine/agouti"
)

const (
	testURL       = "http://localhost:8080/"
	loginEmail    = "***"
	loginPassword = "***"
)

func TestModelFail1(t *testing.T) {
	page := setupPage(t)
	defer page.Destroy()

	navigateAndLogin(t, page)
	goToComposePage(t, page)

	addModelButton := page.FindByButton("Add Model")
	if err := addModelButton.Click(); err != nil {
		t.Fatalf("Failed to click on the Add Model button: %v", err)
	}
	fmt.Println("Clicked on the Add Model button.")

	defineModelButton := page.FindByButton("Define Model")
	if defineModelButton == nil {
		t.Fatalf("Define Model button not found")
	}
	if err := defineModelButton.Click(); err != nil {
		t.Fatalf("Failed to click on the Define Model button: %v", err)
	}
	fmt.Println("Clicked on the Define Model button successfully.")

	nameInput := page.FindByID("name_input")
	if err := nameInput.Fill(" "); err != nil {
		t.Fatalf("Failed to fill the Name input: %v", err)
	}

	page.SetImplicitWait(3000000)

	descriptionInput := page.FindByID("input-3")
	if err := descriptionInput.Fill(" "); err != nil {
		t.Fatalf("Failed to fill the Description input: %v", err)
	}

	groupInput := page.FindByID("input-2")
	if err := groupInput.Fill(" "); err != nil {
		t.Fatalf("Failed to fill the Group input: %v", err)
	}

	nameInput1 := page.FindByID("__BVID__283")
	if err := nameInput1.Fill(" "); err != nil {
		t.Fatalf("Failed to fill the Name input again: %v", err)
	}

	// Click on the 'Any' option in the Type dropdown
	err := page.Find("#__BVID__285 option[value='611d0e75609ae46a09ccab12']").Click()
	handleError1(err, "selecting 'Any' option in the Type dropdown")

	saveModelButton := page.FindByButton("Save Model")
	if err := saveModelButton.Click(); err != nil {
		t.Fatalf("Failed to click on the Save Model button: %v", err)
	}

	previousPageURL := "http://localhost:8080/maestro-ui/modelling"
	if err := page.Navigate(previousPageURL); err != nil {
		t.Fatalf("Failed to navigate back to the previous page: %v", err)
	}

	time.Sleep(5 * time.Second)
}

func setupPage(t *testing.T) *agouti.Page {
	cDriver := agouti.ChromeDriver()
	if err := cDriver.Start(); err != nil {
		t.Fatalf("Failed to start the driver: %s\n", err.Error())
	}
	page, err := cDriver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		t.Fatalf("Failed to open the page: %s\n", err.Error())
	}
	return page
}

func navigateAndLogin(t *testing.T, page *agouti.Page) {
	if err := page.Navigate(testURL); err != nil {
		t.Fatalf("Failed to navigate to the specified page address: %s\n", err.Error())
	}
	emailInput := page.FindByID("login-email")
	if err := emailInput.Fill(loginEmail); err != nil {
		t.Fatalf("Failed to fill the email field: %s\n", err.Error())
	}
	passwordInput := page.FindByID("login-password")
	if err := passwordInput.Fill(loginPassword); err != nil {
		t.Fatalf("Failed to fill the password field: %s\n", err.Error())
	}
	signInButton := page.FindByButton("Sign In")
	if err := signInButton.Click(); err != nil {
		t.Fatalf("Failed to click the Sign In button: %s\n", err.Error())
	}
	page.SetImplicitWait(3000000)
}

func goToComposePage(t *testing.T, page *agouti.Page) {
	composeLink := page.Find("a[href='/maestro-ui/modelling']")
	if err := composeLink.Click(); err != nil {
		t.Fatalf("Failed to click on the Compose link: %v", err)
	}
	fmt.Println("Clicked on the Compose link.")
	page.SetImplicitWait(3000000)
}

func handleError1(err error, msg string) {
	if err != nil {
		fmt.Printf("Error: %s - %s\n", msg, err)
	}
}
