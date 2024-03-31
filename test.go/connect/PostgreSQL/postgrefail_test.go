package main

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/sclevine/agouti"
)

func TestPostgreSQLFail(t *testing.T) {
	driver := agouti.ChromeDriver()
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start the driver: %v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage()
	if err != nil {
		log.Fatalf("Failed to open a new page: %v", err)
	}
	defer page.CloseWindow()

	if err := page.Navigate("http://localhost:8080/"); err != nil {
		log.Fatalf("Failed to navigate to the page: %v", err)
	}

	// Fill in the login credentials
	if err := fillLoginForm(page); err != nil {
		log.Fatalf("Failed to fill login form: %v", err)
	}

	// Create a new connection
	if err := createNewConnection(page); err != nil {
		log.Fatalf("Failed to create a new connection: %v", err)
	}

	// Wait for 3 seconds before finishing the test
	time.Sleep(3 * time.Second)
}

func fillLoginForm(page *agouti.Page) error {
	// Fill in the email input
	if err := page.FindByID("login-email").Fill("***"); err != nil {
		return fmt.Errorf("failed to fill the email field: %v", err)
	}

	// Fill in the password input
	if err := page.FindByID("login-password").Fill("***"); err != nil {
		return fmt.Errorf("failed to fill the password field: %v", err)
	}

	// Sign In
	if err := page.FindByButton("Sign In").Click(); err != nil {
		return fmt.Errorf("failed to click the Sign In button: %v", err)
	}

	return nil
}

func createNewConnection(page *agouti.Page) error {
	page.SetImplicitWait(1000000)

	if err := page.Find("a[href='/maestro-ui/connections']").Click(); err != nil {
		return fmt.Errorf("failed to click on the Connect link: %v", err)
	}

	if err := page.Find("a[href='/maestro-ui/connections/new']").Click(); err != nil {
		return fmt.Errorf("failed to click on the New Connection link: %v", err)
	}

	if err := page.Find("img[src*='postgresql.4ae7e94d.png']").Click(); err != nil {
		return fmt.Errorf("failed to click on the PostgreSQL logo: %v", err)
	}

	// Fill in the connection details
	connectionDetails := map[string]string{
		"#h-connection-name":     "0000",
		"#h-connection-desc":     "1111",
		"#h-connection-host":     "AAAA",
		"#h-connection-port":     "BBBB",
		"#h-connection-database": "100101",
		"#h-connection-userName": "111001",
	}

	for selector, value := range connectionDetails {
		if err := page.Find(selector).Fill(value); err != nil {
			return fmt.Errorf("failed to fill %s: %v", selector, err)
		}
	}

	// Click the Create Connection button
	if err := page.FindByButton("Create Connection").Click(); err != nil {
		return fmt.Errorf("failed to click the Create Connection button: %v", err)
	}

	// Go back to the previous page
	if err := page.Back(); err != nil {
		log.Printf("Failed to go back to the previous page: %v", err)
	}

	return nil
}

func TestPostgreSQLFail2(t *testing.T) {
	driver := agouti.ChromeDriver()
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start the driver: %v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage()
	if err != nil {
		log.Fatalf("Failed to open a new page: %v", err)
	}
	defer page.CloseWindow()

	if err := page.Navigate("http://localhost:8080/"); err != nil {
		log.Fatalf("Failed to navigate to the page: %v", err)
	}

	// Fill in the login credentials
	if err := fillLoginForm1(page); err != nil {
		log.Fatalf("Failed to fill login form: %v", err)
	}

	// Create a new connection
	if err := createNewConnection1(page); err != nil {
		log.Fatalf("Failed to create a new connection: %v", err)
	}

	// Wait for 3 seconds before finishing the test
	time.Sleep(3 * time.Second)
}

func fillLoginForm1(page *agouti.Page) error {
	// Fill in the email input
	if err := page.FindByID("login-email").Fill("admin@maestrohub.com"); err != nil {
		return fmt.Errorf("failed to fill the email field: %v", err)
	}

	// Fill in the password input
	if err := page.FindByID("login-password").Fill("ygiOhWVoi!4WiP"); err != nil {
		return fmt.Errorf("failed to fill the password field: %v", err)
	}

	// Sign In
	if err := page.FindByButton("Sign In").Click(); err != nil {
		return fmt.Errorf("failed to click the Sign In button: %v", err)
	}

	return nil
}

func createNewConnection1(page *agouti.Page) error {
	page.SetImplicitWait(1000000)

	if err := page.Find("a[href='/maestro-ui/connections']").Click(); err != nil {
		return fmt.Errorf("failed to click on the Connect link: %v", err)
	}

	if err := page.Find("a[href='/maestro-ui/connections/new']").Click(); err != nil {
		return fmt.Errorf("failed to click on the New Connection link: %v", err)
	}

	if err := page.Find("img[src*='postgresql.4ae7e94d.png']").Click(); err != nil {
		return fmt.Errorf("failed to click on the PostgreSQL logo: %v", err)
	}

	// Fill in the connection details
	connectionDetails := map[string]string{
		"#h-connection-name":     "AAAA",
		"#h-connection-desc":     "BBBB",
		"#h-connection-host":     "0000",
		"#h-connection-port":     "1111",
		"#h-connection-database": "101001",
		"#h-connection-userName": "010010",
	}

	for selector, value := range connectionDetails {
		if err := page.Find(selector).Fill(value); err != nil {
			return fmt.Errorf("failed to fill %s: %v", selector, err)
		}
	}

	// Click the Create Connection button
	if err := page.FindByButton("Create Connection").Click(); err != nil {
		return fmt.Errorf("failed to click the Create Connection button: %v", err)
	}

	// Go back to the previous page
	if err := page.Back(); err != nil {
		log.Printf("Failed to go back to the previous page: %v", err)
	}

	return nil
}
