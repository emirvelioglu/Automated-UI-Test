package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/sclevine/agouti"
)

func handleError(err error, action string) {
	if err != nil {
		fmt.Printf("Error while %s: %v\n", action, err)
	}
}

func TestModel(t *testing.T) {
	// Start Chrome driver
	cDriver := agouti.ChromeDriver()
	handleError(cDriver.Start(), "starting the driver")
	defer cDriver.Stop()

	// Open a new page
	page, err := cDriver.NewPage()
	handleError(err, "opening a new page")
	defer page.CloseWindow()

	// Navigate to the specified page address
	err = page.Navigate("http://localhost:8080/")
	handleError(err, "navigating to the specified page address")

	// Fill in email and password inputs
	err = page.FindByID("login-email").Fill("***")
	handleError(err, "filling the email field")
	err = page.FindByID("login-password").Fill("***")
	handleError(err, "filling the password field")

	// Click Sign In
	err = page.FindByButton("Sign In").Click()
	handleError(err, "clicking the Sign In button")

	// Wait for page to load
	page.SetImplicitWait(1000000)

	// Go to Compose
	err = page.Find("a[href='/maestro-ui/modelling']").Click()
	handleError(err, "clicking on the Compose link")
	fmt.Println("Clicked on the Compose link.")

	// Wait for page to load
	page.SetImplicitWait(1000000)

	// Click on Add Model button
	err = page.FindByButton("Add Model").Click()
	handleError(err, "clicking on the Add Model button")
	fmt.Println("Clicked on the Add Model button.")

	// Click on Define Model button
	defineModelButton := page.FindByButton("Define Model")
	if defineModelButton == nil {
		t.Fatalf("Define Model button not found")
	}
	err = defineModelButton.Click()
	handleError(err, "clicking on the Define Model button")
	fmt.Println("Clicked on the Define Model button successfully.")

	// Fill inputs with "TEST_MODEL"
	inputIDs := []string{"name_input", "input-3", "input-2", "__BVID__283"}
	for _, inputID := range inputIDs {
		err := page.FindByID(inputID).Fill("TEST_MODEL")
		handleError(err, fmt.Sprintf("filling the %s input", inputID))
	}

	// Select "Any" option in the Type dropdown
	err = page.Find("#__BVID__285 option[value='611d0e75609ae46a09ccab12']").Click()
	handleError(err, "selecting 'Any' option in the Type dropdown")

	// Click on the Save Model button
	err = page.FindByButton("Save Model").Click()
	handleError(err, "clicking on the Save Model button")

	// Wait for page to load
	page.SetImplicitWait(1000000)

	// Navigate back to the Compose page
	err = page.Navigate("http://localhost:8080/maestro-ui/modelling")
	handleError(err, "navigating back to the previous page")

	// Wait for page to load
	page.SetImplicitWait(1000000)

	// Pause for a moment
	time.Sleep(3 * time.Second)

	// Click on the trash icon
	err = page.Find(".feather-trash-2").Click()
	handleError(err, "clicking on the trash icon")
	fmt.Println("Clicked on the trash icon.")

	// Click on the "Yes" button in the confirmation dialog
	buttons := page.All("button.btn.btn-outline-primary")
	yesButton := buttons.At(1)
	err = yesButton.Click()
	handleError(err, "clicking on the 'Yes' button")
	fmt.Println("Deletion successful.")

	// Wait for deletion
	page.SetImplicitWait(1000000)

	// Pause for a moment
	time.Sleep(5 * time.Second)
}
