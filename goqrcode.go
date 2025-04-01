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

	c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()

	if re.MatchString(link) {
		fmt.Printf("The URL: " + link + " matches the pattern of a link!")
		fmt.Println()
		error := qrcode.WriteFile(link, qrcode.Medium, 256, "Files/Link.png")
		if error != nil {
			fmt.Println("Error creating QR Code.")
		} else {
			fmt.Println("QR Code created successfully on Files/Link.png!")
		}
	} else {
		fmt.Println("The URL does not match the pattern of a link!")
		fmt.Println("Follow the format of https://www.example.com")
	}
}

func generateVCard(name, mobile, company, job_title, email, link, address string) string {
	return fmt.Sprintf("BEGIN:VCARD\nVERSION:3.0\nFN:%s\nTEL:%s\nORG:%s\nTITLE:%s\nEMAIL:%s\nURL:%s\nADR:%s\nEND:VCARD", name, mobile, company, job_title, email, link, address)
}

func generatePhone(name, phone string) string {
	return fmt.Sprintf("BEGIN:VCARD\nVERSION:3.0\nFN:%s\nTEL:%s\nEND:VCARD", name, phone)
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
	fmt.Println(vCardData)

	error := qrcode.WriteFile(vCardData, qrcode.Medium, 256, "Files/VCard.png")
		if error != nil {
			fmt.Println("Error creating QR Code.")
		} else {
			fmt.Println("QR Code created successfully on Files/VCard.png!")
		}
}

func phone() {
	var name string
	var phone string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter the following details.")

	fmt.Print("Name: ")
	if scanner.Scan() {
		name = scanner.Text()
	}
	fmt.Print("Phone: ")
	if scanner.Scan() {
		phone = scanner.Text()
	}

	c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()

	phoneData := generatePhone(name, phone)
	fmt.Println(phoneData)

	error := qrcode.WriteFile(phoneData, qrcode.Medium, 256, "Files/Phone.png")
	if error != nil {
		fmt.Println("Error creating QR Code.")
	} else {
		fmt.Println("QR Code created successfully on Files/Phone.png!")
	}
}

func text() {
	var text string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter the text: ")
	if scanner.Scan() {
		text = scanner.Text()
	}

	c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()

	error := qrcode.WriteFile(text, qrcode.Medium, 512, "Files/Text.png")
	if error != nil {
		fmt.Println("Error creating QR Code.")
	} else {
		fmt.Println("QR Code created successfully on Files/Text.png!")
	}
}

func email() {
	var email string
	var subject string
	var body string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter the email: ")
	if scanner.Scan() {
		email = scanner.Text()
	}
	fmt.Print("Enter the subject: ")
	if scanner.Scan() {
		subject = scanner.Text()
	}
	fmt.Print("Enter the body: ")
	if scanner.Scan() {
		body = scanner.Text()
	}

	c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()

	emailData := fmt.Sprintf("mailto:%s?subject=%s&body=%s", email, subject, body)
	error := qrcode.WriteFile(emailData, qrcode.Medium, 512, "Files/Email.png")
	if error != nil {
		fmt.Println("Error creating QR Code.")
	} else {
		fmt.Println("QR Code created successfully on Files/Email.png!")
	}
}

func sms() {
	var number string
	var message string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter the phone number: ")
	if scanner.Scan() {
		number = scanner.Text()
	}
	fmt.Print("Enter the message: ")
	if scanner.Scan() {
		message = scanner.Text()
	}

	c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()

	smsData := fmt.Sprintf("SMSTO:%s:%s", number, message)
	error := qrcode.WriteFile(smsData, qrcode.Medium, 512, "Files/SMS.png")
	if error != nil {
		fmt.Println("Error creating QR Code.")
	} else {
		fmt.Println("QR Code created successfully on Files/SMS.png!")
	}

}

func wifi() {
	var ssid string
	var password string
	var encryption string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter the SSID: ")
	if scanner.Scan() {
		ssid = scanner.Text()
	}
	fmt.Print("Enter the password: ")
	if scanner.Scan() {
		password = scanner.Text()
	}
	fmt.Print("Enter the encryption type (WPA/WEP): ")
	if scanner.Scan() {
		encryption = scanner.Text()
	}

	c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()

	wifiData := fmt.Sprintf("WIFI:S:%s;T:%s;P:%s;;", ssid, encryption, password)
	error := qrcode.WriteFile(wifiData, qrcode.Medium, 512, "Files/WiFi.png")
	if error != nil {
		fmt.Println("Error creating QR Code.")
	} else {
		fmt.Println("QR Code created successfully on Files/WiFi.png!")
	}
}

func location() {
	var latitude string
	var longitude string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter the latitude: ")
	if scanner.Scan() {
		latitude = scanner.Text()
	}
	fmt.Print("Enter the longitude: ")
	if scanner.Scan() {
		longitude = scanner.Text()
	}

	c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()

	locationData := fmt.Sprintf("geo:%s,%s", latitude, longitude)
	error := qrcode.WriteFile(locationData, qrcode.Medium, 512, "Files/Location.png")
	if error != nil {
		fmt.Println("Error creating QR Code.")
	} else {
		fmt.Println("QR Code created successfully on Files/Location.png!")
	}
}

func event() {
	var title string
	var startDate string
	var endDate string
	var location string
	var description string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter the event title: ")
	if scanner.Scan() {
		title = scanner.Text()
	}
	fmt.Print("Enter the start date (YYYYMMDD): ")
	if scanner.Scan() {
		startDate = scanner.Text()
	}
	fmt.Print("Enter the end date (YYYYMMDD): ")
	if scanner.Scan() {
		endDate = scanner.Text()
	}
	fmt.Print("Enter the location: ")
	if scanner.Scan() {
		location = scanner.Text()
	}
	fmt.Print("Enter the description: ")
	if scanner.Scan() {
		description = scanner.Text()
	}

	c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()

	eventData := fmt.Sprintf("BEGIN:VEVENT\nSUMMARY:%s\nDTSTART:%s\nDTEND:%s\nLOCATION:%s\nDESCRIPTION:%s\nEND:VEVENT", title, startDate, endDate, location, description)
	error := qrcode.WriteFile(eventData, qrcode.Medium, 512, "Files/Event.png")
	if error != nil {
		fmt.Println("Error creating QR Code.")
	} else {
		fmt.Println("QR Code created successfully on Files/Event.png!")
	}
}

func main() {

	var choice int
	
	fmt.Println(`
▗▄▄▄▖ ▗▄▄▖      ▗▄▄▖ ▗▄▖ ▗▄▄▄ ▗▄▄▄▖     ▗▄▄▖▗▄▄▄▖▗▖  ▗▖▗▄▄▄▖▗▄▄▖  ▗▄▖▗▄▄▄▖▗▄▖ ▗▄▄▖ 
▐▌ ▐▌ ▐▌ ▐▌    ▐▌   ▐▌ ▐▌▐▌  █▐▌       ▐▌   ▐▌   ▐▛▚▖▐▌▐▌   ▐▌ ▐▌▐▌ ▐▌ █ ▐▌ ▐▌▐▌ ▐▌
▐▌ ▐▌ ▐▛▀▚▖    ▐▌   ▐▌ ▐▌▐▌  █▐▛▀▀▘    ▐▌▝▜▌▐▛▀▀▘▐▌ ▝▜▌▐▛▀▀▘▐▛▀▚▖▐▛▀▜▌ █ ▐▌ ▐▌▐▛▀▚▖
▐▙▄▟▙▖▐▌ ▐▌    ▝▚▄▄▖▝▚▄▞▘▐▙▄▄▀▐▙▄▄▖    ▝▚▄▞▘▐▙▄▄▖▐▌  ▐▌▐▙▄▄▖▐▌ ▐▌▐▌ ▐▌ █ ▝▚▄▞▘▐▌ ▐▌   
    `)
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
	} else if choice == 2 {
		phone()
	} else if choice == 3 {
		text()
	} else if choice == 4 {
		email()
	} else if choice == 5 {
		sms()
	} else if choice == 6 {
		wifi()
	} else if choice == 7 {
		location()
	} else if choice == 8 {
		event()
	} else {
		fmt.Println("Invalid choice. Please select a valid option.")
		main()
	}

}