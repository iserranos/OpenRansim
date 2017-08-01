//	This file is part of OpenRansim.
//
//	Foobar is free software: you can redistribute it and/or modify
//	it under the terms of the GNU General Public License as published by
//	the Free Software Foundation, either version 3 of the License, or
//	(at your option) any later version.
//
//	Foobar is distributed in the hope that it will be useful,
//	but WITHOUT ANY WARRANTY; without even the implied warranty of
//	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//	GNU General Public License for more details.
//
//	You should have received a copy of the GNU General Public License
//	along with Foobar.  If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"crypto/aes"
	"io"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"time"
	rand2 "math/rand"
	"encoding/hex"
	"crypto/sha512"
)

var CryptCmd = &cobra.Command{
	Use:   "crypt",
	Short: "Realize cryptological operations",
	Long:  `Realize cryptological operations`,
	RunE: func(cmd *cobra.Command, args []string) error {

		key := "openransim000000"

		if len(args) == 2 {
			text, err := read_from_file(args[1])
			check(err)
			switch args[0] {
			case "encrypt":
				ciphertext := encrypt(string(text), key)
				write_to_file(ciphertext, args[1])
			case "decrypt":
				plaintext := decrypt(string(text), key)
				write_to_file(plaintext, args[1])
			default:
				fmt.Println("You can:\nencrypt\ndecrypt\nexit")
			}
		} else {
			fmt.Println("Two args is required.\n\tFirst: (encrypt|decrypt)\n\tSecond: File path")
		}
		return nil
	},
}

func decrypt(cipherstring string, keystring string) string {
	// Byte array of the string
	ciphertext := []byte(cipherstring)

	// Key
	key := []byte(keystring)

	// Create the AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// Before even testing the decryption,
	// if the text is too small, then it is incorrect
	if len(ciphertext) < aes.BlockSize {
		panic("Text is too short")
	}

	// Get the 16 byte IV
	iv := ciphertext[:aes.BlockSize]

	// Remove the IV from the ciphertext
	ciphertext = ciphertext[aes.BlockSize:]

	// Return a decrypted stream
	stream := cipher.NewCFBDecrypter(block, iv)

	// Decrypt bytes from ciphertext
	stream.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext)
}

func encrypt(plainstring, keystring string) string {
	// Byte array of the string
	plaintext := []byte(plainstring)

	// Key
	key := []byte(keystring)

	// Create the AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// Empty array of 16 + plaintext length
	// Include the IV at the beginning
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))

	// Slice of first 16 bytes
	iv := ciphertext[:aes.BlockSize]

	// Write 16 rand bytes to fill iv
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	// Return an encrypted stream
	stream := cipher.NewCFBEncrypter(block, iv)

	// Encrypt bytes from plaintext to ciphertext
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return string(ciphertext)
}

func generate_rsa_key() string {
	fmt.Println("Generating key")
	reader := rand.Reader
	bitSize := 4096

	private, err := rsa.GenerateKey(reader, bitSize)
	check(err)
	public := &private.N
	hash := sha512.Sum384([]byte(public.String()))
	pass := hex.EncodeToString(hash[:])[0:16]
	fmt.Println(pass)
	fmt.Sprintf("Key generated: %s", pass)
	return pass
}

func generate_key() string {
	rand2.Seed(time.Now().UTC().UnixNano())
	random := make([]byte, 16)
	for i := 0; i < len(random); i++ {
		random[i] = byte(65 + rand2.Intn(90-65))
	}
	return string(random)
}
