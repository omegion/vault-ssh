package vault

import "fmt"

// SecretDataNotFound occurs when the requested fixture file does not exist.
type SecretDataNotFound struct {
	Key  string
	Path string
}

func (e SecretDataNotFound) Error() string {
	return fmt.Sprintf("Vault secret not found with path %s, key %s", e.Path, e.Key)
}
