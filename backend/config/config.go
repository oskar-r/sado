package config

import (
	"log"
	"os"
)

type Conf struct {
	port                string
	environment         string
	rsaPrivate          string
	rsaPublic           string
	sqlConnectionMain   string
	sqlConnectionPolicy string
	migrationDate       string
	origin              string
	listenAndServe      string
	awsID               string
	awsSecret           string
	policyModel         string
	cipherKey           string
}

func SetConfType() *Conf {
	config := &Conf{}

	if os.Getenv("ENVIRONMENT") == "" {
		config.environment = "localhost"
	} else {
		config.environment = os.Getenv("ENVIRONMENT")
	}
	if os.Getenv("PORT") == "" {
		config.port = "8091"
	} else {
		config.port = os.Getenv("PORT")
	}

	if os.Getenv("MINIO_CLIENT_ID") == "" || os.Getenv("MINIO_CLIENT_SECRET") == "" {
		log.Fatal("No minio credentials set\n")
	} else {
		config.awsID = os.Getenv("MINIO_CLIENT_ID")
		config.awsSecret = os.Getenv("MINIO_CLIENT_SECRET")
	}

	switch config.environment {
	case "localhost":
		config.listenAndServe = "localhost:" + config.port
	default:
		config.listenAndServe = ":" + config.port
	}

	config.rsaPrivate = os.Getenv("PORTAL_RSA_PRIV_KEY")
	config.rsaPublic = os.Getenv("PORTAL_RSA_PUB_KEY")

	if os.Getenv("PSQL_CONNECTION_POLICY") == "" {
		log.Fatal("No policy postgress configuration")
	}
	config.sqlConnectionPolicy = os.Getenv("PSQL_CONNECTION_POLICY")

	if os.Getenv("PSQL_CONNECTION_MAIN") == "" {
		log.Fatal("No main postgress configuration")
	}
	config.sqlConnectionMain = os.Getenv("PSQL_CONNECTION_MAIN")

	if os.Getenv("CIPHER_KEY") == "" {
		log.Fatal("no cipher key set")
	}

	config.cipherKey = os.Getenv("CIPHER_KEY")

	config.origin = "*"
	if os.Getenv("POLICY_MODEL_PATH") == "" {
		log.Fatal("No policy model provided \n")
	} else {
		config.policyModel = os.Getenv("POLICY_MODEL_PATH")
	}
	return config
}

func (c *Conf) Get(option string) string {
	switch option {
	case "port":
		return c.port
	case "env":
		return c.environment
	case "rsapub":
		return c.rsaPublic
	case "rsapriv":
		return c.rsaPrivate
	case "sqlcon-main":
		return c.sqlConnectionMain
	case "sqlcon-policy":
		return c.sqlConnectionPolicy
	case "origin":
		return c.origin
	case "server":
		return c.listenAndServe
	case "migrationdate":
		return c.migrationDate
	case "aws-id":
		return c.awsID
	case "aws-secret":
		return c.awsSecret
	case "policy-model":
		return c.policyModel
	case "cipher-key":
		return c.cipherKey
	}
	return ""
}

func (c *Conf) GetFromMatrix(sm string, key1 string, key2 string) (string, bool) {
	panic("not implemented")
}
func (c *Conf) SetFromDB() error {
	panic("not implemented")
}
