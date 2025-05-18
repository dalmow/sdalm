package utils

import (
	"context"
	"crypto/rand"
	"errors"
	"math/big"
)

type DuplicityChecker interface {
	AlreadyExists(ctx context.Context, id string) (bool, error)
}

var ErrUnableToGenerateAlias = errors.New("could not generate unique alias after retries")

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const maxRetries = 3
const length = 6

func RandomAliasGen(ctx context.Context, checker DuplicityChecker) (string, error) {
	for range maxRetries {
		alias, err := generateAlias()
		if err != nil {
			return "", err
		}

		exists, err := checker.AlreadyExists(ctx, alias)
		if err != nil {
			return "", nil
		}

		if !exists {
			return alias, nil
		}
	}

	return "", ErrUnableToGenerateAlias
}

func generateAlias() (string, error) {
	alias := make([]byte, length)
	for i := range alias {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		alias[i] = charset[n.Int64()]
	}

	return string(alias), nil
}
