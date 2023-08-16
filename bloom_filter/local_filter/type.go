package local_filter

type Encryptor interface {
	Encrypt(string) int32
}
