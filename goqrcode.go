package main

import ("fmt"
	"github.com/skip2/go-qrcode"
	"regexp"
	"bufio"
	"os"
	"os/exec"
)

func url() {
	var link string
	regex := `^(https?:\/\/[a-zA-Z0-9-]+\.[a-zA-Z0-9-]+(\.[a-zA-Z]{2,})?)$`

	fmt.Printf("Enter the URL: ")
	fmt.Scan(&link)
	re := regexp.MustCompile(regex)

	if re.MatchString(link) {
		fmt.Println("The URL matches the pattern!")
		error := qrcode.WriteFile(link, qrcode.Medium, 256, "QR_Code_Link.png")
		if error != nil {
			fmt.Println("Error creating QR Code.")
		} else {
			fmt.Println("QR Code created successfully!")
		}
	} else {
		fmt.Println("The URL does not match the pattern!")
		fmt.Println("Follow the format of https://www.example.com")
	}
}

func generateVCard(name, mobile, company, job_title, email, link, address string) string {
	return fmt.Sprintf("BEGIN:VCARD\nVERSION:3.0\nFN:%s\nTEL;TYPE=CELL:%s\nORG:%s\nTITLE:%s\nEMAIL:%s\nURL:%s\nADR:%s\nEND:VCARD", name, mobile, company, job_title, email, link, address)

}

func vcard() {
	var name string
	var mobile string
	var company string
	var job_title string
	var email string
	var link string
	var address string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter the following VCard details.")

	fmt.Print("Name: ")
	if scanner.Scan() {
		name = scanner.Text()
	}
	fmt.Print("Mobile: ")
	if scanner.Scan() {
		mobile = scanner.Text()
	}
	fmt.Print("Company: ")
	if scanner.Scan() {
		company = scanner.Text()
	}
	fmt.Print("Job Title: ")
	if scanner.Scan() {
		job_title = scanner.Text()
	}
	fmt.Print("Email: ")
	if scanner.Scan() {
		email = scanner.Text()
	}
	fmt.Print("Website URL: ")
	if scanner.Scan() {
		link = scanner.Text()
	}
	fmt.Print("Address: ")
	if scanner.Scan() {
		address = scanner.Text()
	}

	c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()

	vCardData := generateVCard(name, mobile, company, job_title, email, link, address)
	fmt.Println("VCard Data: ", vCardData)

	error := qrcode.WriteFile(vCardData, qrcode.Medium, 256, "QR_Code_VCard.png")
		if error != nil {
			fmt.Println("Error creating QR Code.")
		} else {
			fmt.Println("QR Code created successfully!")
		}
}

func main() {

	var choice int

	fmt.Println("[0] : URL")
	fmt.Println("[1] : VCard")
	fmt.Println("[2] : Phone")
	fmt.Println("[3] : Text")
	fmt.Println("[4] : Email")
	fmt.Println("[5] : SMS")
	fmt.Println("[6] : WIFI")
	fmt.Println("[7] : Location")
	fmt.Println("[8] : Event")
	fmt.Println()
	fmt.Printf("Select your choice: ")

	fmt.Scan(&choice)

	if choice == 0 {
		url()
	} else if choice == 1 {
		vcard()
	}

}