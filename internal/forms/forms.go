package forms

import (
	"github.com/cmd/internal/database"
	"github.com/cmd/internal/database/datanotes"
	"log"
	"strings"
	"unicode"
)

func IsEmail(email string) bool {

	if strings.HasSuffix(email, "@mail.ru") {
		return true
	} else if strings.HasSuffix(email, "@gmail.com") {
		return true
	} else if strings.HasSuffix(email, "@Outlook.com") {
		return true
	} else {
		return false
	}
}

func IsPassword(password string) bool {
	len := len(password)
	fchar := password[0]

	if len > 8 && unicode.IsUpper(rune(fchar)) {
		return true
	} else {
		return false
	}
}

func IsUsername(username string) bool {
	usern, err := usecase.GetByUserName()

	if err != nil {
		log.Fatal("Failed to get username")
	}

	var IsTrue bool

	for _, names := range usern {
		if names == username {
			IsTrue = false
		} else {
			IsTrue = true
		}
	}

	return IsTrue
}

func IsID(id string) bool {
	notes, err := notesdb.GetNotes()

	if err != nil {
		log.Print(err)
	}

	noteID := []string{}

	for _, nt := range notes {
		noteID = append(noteID, nt.ID)
	}

	var IsTrue bool = false

	for _, i := range noteID {
		if i == id {
			IsTrue = true
		}
	}
	
	return IsTrue
}
