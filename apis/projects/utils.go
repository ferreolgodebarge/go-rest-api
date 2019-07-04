package projects

import (
	"crypto/rand"
	"fmt"
	"log"
)

// To be deleted
var initialDB = ListProjects{
	Projects: []Project{
		Project{
			ID:          "9438167c-9493-4993-fd48-950b27aad7c9",
			Name:        "Project1",
			Description: "Project1 description",
		},
		Project{
			ID:          "1234167c-9493-4993-fd48-950b27a12341",
			Name:        "Project2",
			Description: "Project2 description",
		},
	},
}

// GenerateUUID : UUID generation
func GenerateUUID() (uuid string) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return
}
