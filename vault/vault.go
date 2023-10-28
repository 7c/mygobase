package vault

import (
	"context"
	"fmt"

	vault "github.com/hashicorp/vault/api"
)

type ModelVault struct {
	client *vault.Client
}

func (v *ModelVault) Init(VaultUrl, token string) (*vault.Client, error) {
	V := vault.DefaultConfig()
	client, err := vault.NewClient(V)
	client.SetAddress(VaultUrl)
	client.SetToken(token)
	if err != nil {
		return nil, err
	}
	v.client = client
	return client, nil
}

func (v *ModelVault) KVv1(subPath string) (*vault.KVSecret, error) {
	secret, err := v.client.KVv1("secret").Get(context.Background(), subPath)
	if err != nil {
		return nil, err
	}
	return secret, nil
}

func (v *ModelVault) AsSqlUri(subPath string) (string, error) {
	secret, err := v.KVv1(subPath)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", secret.Data["login"], secret.Data["password"], secret.Data["host"], secret.Data["port"], secret.Data["db"]), nil
}
