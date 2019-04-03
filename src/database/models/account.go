package models

import (
	"encoding/base64"
	"time"

	"github.com/kubesmith/kubesmith-server/src/config"
	"github.com/kubesmith/kubesmith-server/src/encryption"
	uuid "github.com/satori/go.uuid"
)

type Account struct {
	ID                string `json:"id"`
	UserID            string `json:"userID"`
	Email             string `json:"email"`
	Password          string `json:"-"`
	PasswordResetCode string `json:"-"`
	Type              string `json:"type"`
	Verified          bool   `json:"verified"`
	CreatedAt         uint32 `json:"createdAt"`
	UpdatedAt         uint32 `json:"updatedAt"`
}

func (a *Account) decodeValue(value string) (string, error) {
	tmp, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return "", err
	}

	decrypedValue, err := encryption.Decrypt(tmp, []byte(config.Parsed.DatabaseEncryptionKey))
	if err != nil {
		return "", err
	}

	return string(decrypedValue), nil
}

func (a *Account) encodeValue(value string) (string, error) {
	tmp, err := encryption.Encrypt([]byte(value), []byte(config.Parsed.DatabaseEncryptionKey))
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(tmp), nil
}

func (a *Account) AfterFind() error {
	if a.Password != "" {
		password, err := a.decodeValue(a.Password)
		if err != nil {
			return err
		}

		a.Password = password
	}

	return nil
}

func (a *Account) BeforeCreate() error {
	a.ID = uuid.NewV4().String()
	a.CreatedAt = uint32(time.Now().Unix())

	return nil
}

func (a *Account) BeforeSave() error {
	if a.Password != "" {
		password, err := a.encodeValue(a.Password)
		if err != nil {
			return err
		}

		a.Password = password
	}

	a.UpdatedAt = uint32(time.Now().Unix())

	return nil
}
