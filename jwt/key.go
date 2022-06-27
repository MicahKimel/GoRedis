package jwt

func (j *Jwt) GenKey(name string) []byte {
	// TODO: Pull From ENV
	signingKey := []byte("keymaker")
	// TODO: WRITE BASH TO MANIPULATE KEY

	// TODO: STORE IN DB
	return signingKey
}

func (j *Jwt) GetKey(token string) []byte {
	// TODO: Hash & Pull From DB
	signingKey := []byte("keymaker")
	return signingKey
}
