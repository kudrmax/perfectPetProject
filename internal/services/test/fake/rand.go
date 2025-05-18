package fake

import "crypto/rand"

func RandString() string {
	randString := rand.Text()

	return randString
}
