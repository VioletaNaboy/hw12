package password

import (
	"fmt"
)

type Manager struct {
	storage *Storage
}

func NewManager(filename string) *Manager {
	storage := NewStorage(filename)
	return &Manager{storage: storage}
}

func (m *Manager) ListPasswords() {
	passwords := m.storage.Load()
	fmt.Println("Stored passwords:")
	for name := range passwords {
		fmt.Println(name)
	}
}

func (m *Manager) SavePassword(name, password string) {
	passwords := m.storage.Load()
	if passwords == nil {
		passwords = make(map[string]string)
	}
	passwords[name] = password
	m.storage.Save(passwords)
}

func (m *Manager) RetrievePassword(name string) {
	passwords := m.storage.Load()
	if pass, found := passwords[name]; found {
		fmt.Printf("Password for %s: %s\n", name, pass)
	} else {
		fmt.Println("No password found for that name.")
	}
}
