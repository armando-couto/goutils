package goutils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

// Chave e IV devem ter tamanhos específicos para AES-256 (32 bytes para chave e 16 bytes para IV)
type Crypt struct {
	Key []byte
	Iv  []byte
}

func (crypt Crypt) Encrypt(data []byte) (string, error) {
	block, err := aes.NewCipher(crypt.Key)
	if err != nil {
		return "", err
	}

	// Preencher o bloco com dados
	data = pad(data, block.BlockSize())

	// Criar um modo de cifra em bloco com o bloco e o IV
	mode := cipher.NewCFBEncrypter(block, crypt.Iv)

	// Criar um slice para o texto cifrado (tamanho igual ao texto original)
	ciphertext := make([]byte, len(data))

	// Criptografar os dados
	mode.XORKeyStream(ciphertext, data)

	return base64.StdEncoding.EncodeToString(append(crypt.Iv, ciphertext...)), nil
}

func pad(data []byte, blockSize int) []byte {
	padSize := blockSize - len(data)%blockSize
	pad := bytes.Repeat([]byte{byte(padSize)}, padSize)
	return append(data, pad...)
}

func (crypt Crypt) Decrypt(encryptedData string) ([]byte, error) {
	// Decodifica o dado criptografado de Base64
	data, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return nil, err
	}

	// Cria um novo bloco de cifra AES
	block, err := aes.NewCipher(crypt.Key)
	if err != nil {
		return nil, err
	}

	// Verifica se o tamanho dos dados é válido
	if len(data) < aes.BlockSize {
		return nil, fmt.Errorf("texto cifrado muito curto")
	}

	// Criar um modo de cifra em bloco com o bloco e o IV
	mode := cipher.NewCFBDecrypter(block, crypt.Iv)

	// Criar um slice para o texto decifrado (tamanho igual ao texto cifrado)
	decrypted := make([]byte, len(data))

	// Decifrar os dados
	mode.XORKeyStream(decrypted, data)

	// Remover o preenchimento
	decrypted = unpad(decrypted)

	return decrypted, nil
}

func unpad(data []byte) []byte {
	if len(data) == 0 {
		return data
	}
	padSize := int(data[len(data)-1])
	if padSize > len(data) {
		return data
	}
	return data[:len(data)-padSize]
}
