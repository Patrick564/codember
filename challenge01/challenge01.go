package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

func verifyFields(user []string) bool {
	// Correct fields
	fields := []string{"usr", "eme", "psw", "age", "loc", "fll"}

	// Set all user fields
	uFields := []string{}
	for _, f := range user {
		uFields = append(uFields, strings.Split(f, ":")[0])
	}

	// Remove correct fields
	for idx, v := range fields {
		if len(fields) == 0 {
			return true
		}

		if slices.Contains(uFields, v) {
			fields = fields[idx:]
		}
	}

	return false
}

func main() {
	users := make([][]string, 0)

	file, err := os.ReadFile("users.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Remove all blank lines and create an array for every possible user
	r := strings.ReplaceAll(string(file), "\n", " ")
	content := strings.Split(r, "  ")

	for _, c := range content {
		user := strings.Split(c, " ")

		// Omit user if not have all fields
		if len(user) < 6 {
			continue
		}

		if verifyFields(user) {
			users = append(users, user)
		}
	}

	fmt.Printf("List of users: %s\nCount of users: %d\nLast valid user: %s\n", users, len(users), users[len(users)-1])
}
