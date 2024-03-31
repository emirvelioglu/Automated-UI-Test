package testing

import (
	"testing"
	"time"

	"github.com/sclevine/agouti"
)

const (
	loginURL    = "http://localhost:8080/"
	expectedURL = "http://localhost:8080/dashboard"
)

func TestMongo(t *testing.T) {
	driver := agouti.ChromeDriver()
	if err := driver.Start(); err != nil {
		t.Fatalf("Failed to start the driver: %s\n", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage()
	if err != nil {
		t.Fatalf("Failed to open the page: %s\n", err)
	}
	defer page.CloseWindow()

	if err := page.Navigate(loginURL); err != nil {
		t.Fatalf("Failed to navigate to the specified page address: %s\n", err)
	}

	if err := login(page); err != nil {
		t.Fatalf("Login failed: %s\n", err)
	}

	if err := createMongoDBConnection(page); err != nil {
		t.Fatalf("Failed to create MongoDB connection: %s\n", err)
	}

	if err := deleteMongoDBConnection(page); err != nil {
		t.Fatalf("Failed to delete MongoDB connection: %s\n", err)
	}
}

func login(page *agouti.Page) error {
	page.SetImplicitWait(1000000)
	if err := page.FindByID("login-email").Fill("***"); err != nil {
		return err
	}
	if err := page.FindByID("login-password").Fill("***"); err != nil {
		return err
	}
	signInButton := page.FindByButton("Sign In")
	return signInButton.Click()
}

func createMongoDBConnection(page *agouti.Page) error {
	page.SetImplicitWait(1000000)
	if err := navigateToConnections(page); err != nil {
		return err
	}

	if err := navigateToNewConnection(page); err != nil {
		return err
	}

	if err := page.Find("img[src*='mongodb.4f927bb4.png']").Click(); err != nil {
		return err
	}

	if err := fillConnectionDetails(page); err != nil {
		return err
	}

	button := page.FindByButton("Create MongoDb Connection")
	return button.Click()
}

func navigateToConnections(page *agouti.Page) error {
	page.SetImplicitWait(1000000)
	connectLink := page.Find("a[href='/maestro-ui/connections']")
	return connectLink.Click()
}

func navigateToNewConnection(page *agouti.Page) error {
	page.SetImplicitWait(1000000)
	connectLink2 := page.Find("a[href='/maestro-ui/connections/new']")
	return connectLink2.Click()
}

func fillConnectionDetails(page *agouti.Page) error {
	page.SetImplicitWait(1000000)
	if err := page.FindByID("h-connection-name").Fill("TEST1"); err != nil {
		return err
	}
	if err := page.FindByID("h-connection-desc").Fill("TEST1"); err != nil {
		return err
	}
	if err := page.Find("#__BVID__284").Fill("TEST1"); err != nil {
		return err
	}
	if err := page.FindByID("h-connection-host").Fill("192.168.1.1"); err != nil {
		return err
	}
	if err := page.FindByID("h-connection-port").Fill("27017"); err != nil {
		return err
	}
	if err := page.FindByID("h-connection-database").Fill("MaestroAutomationUITEST1"); err != nil {
		return err
	}
	if err := page.FindByID("h-connection-userName").Fill("Maestro"); err != nil {
		return err
	}
	return page.Find("#h-connection-auth").Select("SCRAM-SHA-1")
}

func deleteMongoDBConnection(page *agouti.Page) error {
	page.SetImplicitWait(1000000)
	if err := page.Find(".feather-trash-2").Click(); err != nil {
		return err
	}

	time.Sleep(1 * time.Second)

	buttons := page.All("button.btn.btn-outline-primary")
	yesButton := buttons.At(1)
	return yesButton.Click()
}
