package vigenereCipher

import "testing"

func TestEncrypt(t *testing.T) {

	input := "Example"
	key := "key"
	expectedOutput := "OBywtJo"

	output, err := Encrypt(input, key)

	if err != nil {
		t.Fatalf("Encryption failed %v", err)
	}

	if output != expectedOutput {
		t.Fatalf("Encryption: Got '%s' expected '%s'", output, expectedOutput)
	}
}

func TestEncryptFail(t *testing.T) {

	input := "ääääääää"
	key := "key"

	_, err := Encrypt(input, key)

	if err == nil {
		t.Fatalf("Encryption should have failed")
	}
}

func TestDecrypt(t *testing.T) {

	input := "OBywtJo"
	key := "key"
	expectedOutput := "Example"

	output, err := Decrypt(input, key)

	if err != nil {
		t.Fatalf("Decryption failed %v", err)
	}

	if output != expectedOutput {
		t.Fatalf("Decryption: Got '%s' expected '%s'", output, expectedOutput)
	}
}

func TestDecryptFail(t *testing.T) {

	input := "ääääääää"
	key := "key"

	_, err := Decrypt(input, key)

	if err == nil {
		t.Fatalf("Decryption should have failed")
	}

}
