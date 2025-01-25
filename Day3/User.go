package Day3

import (
	"encoding/json"
	"log"
	"os"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Profile  string `json:"profile"` // Path to profile image
}

const dataFile = "users.json"

func LoadUsers() (map[string]User, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var users map[string]User

	json.NewDecoder(file).Decode(&users)
	return users, nil
}

func saveUsers(users map[string]User) {
	file, err := os.Create(dataFile)
	if err != nil {
		log.Fatalf("Failed to create data file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(users)
	if err != nil {
		log.Fatalf("Failed to save data file: %v", err)
	}
}
