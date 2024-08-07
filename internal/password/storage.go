package password

import (
	"encoding/json"
	"fmt"
	"os"
)

type Storage struct {
	filename string
}

func NewStorage(filename string) *Storage {
	return &Storage{filename: filename}
}

func (s *Storage) Load() map[string]string {
	file, err := os.Open(s.filename)
	if err != nil {
		if os.IsNotExist(err) {
			return make(map[string]string)
		}
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	var passwords map[string]string
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&passwords); err != nil {
		fmt.Println("Error decoding file:", err)
		return nil
	}
	return passwords
}

func (s *Storage) Save(passwords map[string]string) {
	file, err := os.Create(s.filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(passwords); err != nil {
		fmt.Println("Error encoding file:", err)
	}
}
