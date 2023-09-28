package config

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"log"

	vault "github.com/hashicorp/vault/api"
	"github.com/ilyakaznacheev/cleanenv"
)

var (
	//go:embed viper.yaml
	rootConfigEmbed []byte
)

const (
	mountPath      = "secret"
	secretPathBase = "gitlab.secret.kz/clearing-center/core/wallets/"
)

type VaultConfig struct {
	Address string `yaml:"address"`
	Token   string `yaml:"token"`
	Env     string `yaml:"env"`
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
	var vaultConfGlob VaultConfig
	var cfg *Config
	err := cleanenv.ParseYAML(bytes.NewReader(rootConfigEmbed), &vaultConfGlob)
	if err != nil {
		panic(err)
	}

	// prepare a client with the given base address
	vConf := vault.DefaultConfig()
	vConf.Address = vaultConfGlob.Address

	client, err := vault.NewClient(vConf)
	if err != nil {
		log.Fatal(err)
	}

	client.SetToken(vaultConfGlob.Token)

	secret, err := client.KVv2(mountPath).Get(context.Background(), fmt.Sprint(secretPathBase, vaultConfGlob.Env))
	if err != nil {
		log.Fatalf("unable to read secret: %v", err)
	}
	jsonb, err := json.Marshal(secret.Data)
	if err != nil {
		log.Fatalf("unable to marshal secret map: %v", err)
	}

	err = json.Unmarshal(jsonb, &cfg)
	if err != nil {
		log.Fatalf("unable to unmarshal secret map: %v", err)
	}
	return cfg
}
