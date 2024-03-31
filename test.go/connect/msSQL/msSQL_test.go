package main

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/sclevine/agouti"
)

const (
	loginURL    = "http://localhost:8080/"
	expectedURL = "http://localhost:8080/dashboard"
)

func TestMsSQL(t *testing.T) {
	cDriver := agouti.ChromeDriver()
	if err := cDriver.Start(); err != nil {
		t.Fatalf("Failed to start the driver: %s\n", err.Error())
	}
	defer cDriver.Stop()

	page, err := cDriver.NewPage()
	if err != nil {
		t.Fatalf("Failed to open the page: %s\n", err.Error())
	}
	defer page.CloseWindow()

	if err := page.Navigate(loginURL); err != nil {
		t.Fatalf("Failed to navigate to the specified page address: %s\n", err.Error())
	}

	if err := login(page); err != nil {
		t.Fatalf("Login failed: %s\n", err.Error())
	}

	if err := navigateToConnections(page); err != nil {
		t.Fatalf("Failed to navigate to connections: %s\n", err.Error())
	}

	if err := navigateToNewConnection(page); err != nil {
		t.Fatalf("Failed to navigate to new connection: %s\n", err.Error())
	}

	if err := selectMsSQLConnection(page); err != nil {
		t.Fatalf("Failed to select MsSQL connection: %s\n", err.Error())
	}

	if err := fillConnectionDetails(page); err != nil {
		t.Fatalf("Failed to fill connection details: %s\n", err.Error())
	}

	if err := createConnection(page); err != nil {
		t.Fatalf("Failed to create connection: %s\n", err.Error())
	}

	if err := waitAndDeleteConnection(page); err != nil {
		t.Fatalf("Failed to delete connection: %s\n", err.Error())
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

func selectMsSQLConnection(page *agouti.Page) error {
	page.SetImplicitWait(1000000)
	mysqlIcon := page.Find("img[src*='mysql.1c88db6f.png']")

	if err := mysqlIcon.Click(); err != nil {
		log.Fatalf("MySQL simgesine tıklanamadı: %v\n", err)
	}
	fmt.Println("MySQL simgesine başarıyla tıklandı.")
	connectionTypeDropdown := page.FindByID("h-connection-type")
	page.SetImplicitWait(1000000)
	sqlOptgroup := connectionTypeDropdown.Find("optgroup[label='SQL']")
	page.SetImplicitWait(1000000)
	msSqlOption := sqlOptgroup.Find("option[value='msSql']")
	return msSqlOption.Click()
}

func fillConnectionDetails(page *agouti.Page) error {
	page.SetImplicitWait(1000000)
	if err := page.FindByID("h-connection-name").Fill("TEST_3"); err != nil {
		return err
	}
	if err := page.FindByID("h-connection-desc").Fill("TEST 3"); err != nil {
		return err
	}
	if err := page.FindByID("h-connection-host").Fill("1433"); err != nil {
		return err
	}
	if err := page.FindByID("h-connection-port").Fill("1433"); err != nil {
		return err
	}
	if err := page.FindByID("h-connection-database").Fill("TEST 3"); err != nil {
		return err
	}
	if err := page.FindByID("h-connection-userName").Fill("TEST 3"); err != nil {
		return err
	}
	return nil
}

func createConnection(page *agouti.Page) error {
	page.SetImplicitWait(1000000)
	createConnectionButton := page.FindByButton("Create Connection")
	return createConnectionButton.Click()
}

func waitAndDeleteConnection(page *agouti.Page) error {
	page.SetImplicitWait(1000000)
	trashIcon := page.Find(".feather-trash-2")
	if err := trashIcon.Click(); err != nil {
		return err
	}

	time.Sleep(1 * time.Second)

	buttons := page.All("button.btn.btn-outline-primary")
	yesButton := buttons.At(1)
	return yesButton.Click()

}
