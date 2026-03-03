package intermediate

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {

	password := "password123"

	hash := sha256.Sum256([]byte(password))
	hash512 := sha512.Sum512([]byte(password))

	fmt.Println(password)
	fmt.Println(hash)
	fmt.Println(hash512)
	fmt.Printf("SHA-256 Hash hex val: %x\n", hash)
	fmt.Printf("SHA-512 Hash hex val: %x\n", hash512)

	// Generate salt
	salt, err := generateSalt()
	if err != nil {
		fmt.Println("Error generating salt:", err)
		return
	}

	// Hash the password with salt
	signUpHash := hashPassword(password, salt)

	// Store the salt and password in database (printing now)
	saltStr := base64.StdEncoding.EncodeToString(salt)
	fmt.Println("Salt:", salt)
	fmt.Printf("Salt: %x\n", salt)
	fmt.Println("Salt base64 encoded:", saltStr)
	fmt.Println("Sign up hash:", signUpHash)
	hashOriginalPassword := sha256.Sum256([]byte(password))
	fmt.Println("Hash of just the password w/o salt:", base64.StdEncoding.EncodeToString(hashOriginalPassword[:]))

	// Verify password
	// Retrieve the salt and decode
	decodedSalt, err := base64.StdEncoding.DecodeString(saltStr)
	if err != nil {
		fmt.Println("Error decoding salt:", err)
		return
	}
	loginHash := hashPassword("password124", decodedSalt)

	// Compare the stored sign up hash with the login hash
	if signUpHash == loginHash {
		fmt.Println("Password is correct. You are logged in.")
	} else {
		fmt.Println("Login failed. Please check user credentials.")
	}
}

func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

// Function to hash password
func hashPassword(password string, salt []byte) string {
	saltedPassword := append(salt, []byte(password)...)
	hash := sha256.Sum256(saltedPassword)
	return base64.StdEncoding.EncodeToString(hash[:])
}
