package adapters

import "golang.org/x/crypto/bcrypt"

type BCryptAdapter struct{}

func (a *BCryptAdapter) MakeHash(input string) (string, error) {
	inBytes := []byte(input)
	hash, err := bcrypt.GenerateFromPassword(inBytes, 10)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (a *BCryptAdapter) CompareHash(source, target string) error {
	srcBytes := []byte(source)
	tgtBytes := []byte(target)
	return bcrypt.CompareHashAndPassword(srcBytes, tgtBytes)
}
