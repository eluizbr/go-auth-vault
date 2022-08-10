package configs

import (
	"log"

	vault "github.com/hashicorp/vault/api"
)

var VaultConn *vault.Client

func ConnectVault() {
	config := vault.DefaultConfig()
	config.Address = "http://localhost:8200"
	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("unable to initialize Vault client: %v", err)
	}

	client.SetToken("123456")
	VaultConn = client

}
