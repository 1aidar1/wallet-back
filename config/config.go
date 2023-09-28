package config

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	vault "github.com/hashicorp/vault/api"
)

var (
	generalOnce     sync.Once
	generalConfGlob *Config

	vaultConfOnce sync.Once
	vaultConfGlob *VaultConfig

	//go:embed viper.yaml
	rootConfigEmbed []byte
)

const (
	mountPath      = "secret"
	secretPathBase = "gitlab.secret.kz/clearing-center/core/wallets/"
)

type VaultConfig struct {
}

type Config struct {
	App struct {
		HttpPort int `json:"http_port"`
		GrpcPort int `json:"grpc_port"`
	} `json:"app"`
	Postgres struct {
		User     string `json:"user"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Db       string `json:"db"`
		SslMode  string `json:"ssl_mode"`
		Timezone string `json:"timezone"`
	}
	Redis struct {
		Addr     string `json:"addr"`
		Password string `json:"password"`
		DB       int    `json:"db"`
	}
}

func GetGeneralConfig() *Config {
	generalOnce.Do(func() {
		parseVaultData()
	})

	return generalConfGlob
}

// func GetVaultConfig() *VaultConfig {
// 	vaultConfOnce.Do(func() {
// 		conf := VaultConfig{}

// 		err := conf.Init(rootConfigEmbed)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		vaultConfGlob = &conf
// 	})

//		return vaultConfGlob
//	}
func parseVaultData() {

	ctx := context.Background()

	// prepare a client with the given base address
	vConf := vault.DefaultConfig()
	vConf.Address = "http://127.0.0.1:8200"

	client, err := vault.NewClient(vConf)
	if err != nil {
		log.Fatal(err)
	}

	// authenticate with a root token (insecure)
	client.SetToken("hvs.fpmG4vI4xoru4iHm3bGrqlnQ")

	secret, err := client.KVv2(mountPath).Get(ctx, getSecretPath())
	if err != nil {
		log.Fatalf("unable to read secret: %v", err)
	}
	jsonb, err := json.Marshal(secret.Data)
	if err != nil {
		log.Fatalf("unable to marshal secret map: %v", err)
	}

	err = json.Unmarshal(jsonb, &generalConfGlob)
	if err != nil {
		log.Fatalf("unable to unmarshal secret map: %v", err)
	}
	// conf := GetVaultConfig()

	// client, err := kit.VaultClient(conf)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// ctx := context.Background()

	// secret, err := client.KVv2(mountPath).Get(ctx, getSecretPath())
	// if err != nil {
	// 	log.Fatalf("unable to read secret: %v", err)
	// }

	// jsonb, err := json.Marshal(secret.Data)
	// if err != nil {
	// 	log.Fatalf("unable to marshal secret map: %v", err)
	// }

	// err = json.Unmarshal(jsonb, &generalConfGlob)
	// if err != nil {
	// 	log.Fatalf("unable to unmarshal secret map: %v", err)
	// }
}

func getSecretPath() string {
	// vault := GetVaultConfig()
	return fmt.Sprint(secretPathBase, "dev")
}
