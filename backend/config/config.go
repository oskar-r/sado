package config

import (
	"log"
	"os"
)

type Conf struct {
	port                string
	wsPort              string
	environment         string
	rsaPrivate          string
	rsaPublic           string
	sqlConnectionMain   string
	sqlConnectionPolicy string
	migrationDate       string
	origin              string
	listenAndServe      string
	wsServer            string
	awsID               string
	awsSecret           string
	s3server            string
	policyModel         string
	cipherKey           string
	minioHTTPS          string
	natsServer          string
	natsUser            string
	natsPassword        string
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

	if os.Getenv("WS_PORT") == "" {
		config.wsPort = "8092"
	} else {
		config.wsPort = os.Getenv("WS_PORT")
	}

	if os.Getenv("MINIO_CLIENT_ID") == "" || os.Getenv("MINIO_CLIENT_SECRET") == "" || os.Getenv("MINIO_SERVER") == "" {
		log.Fatal("No minio credentials set\n")
	} else {
		config.awsID = os.Getenv("MINIO_CLIENT_ID")
		config.awsSecret = os.Getenv("MINIO_CLIENT_SECRET")
		config.s3server = os.Getenv("MINIO_SERVER")
	}

	switch config.environment {
	case "localhost":
		config.listenAndServe = "localhost:" + config.port
		config.wsServer = "localhost:" + config.wsPort
	default:
		config.listenAndServe = ":" + config.port
		config.wsServer = ":" + config.wsPort
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
	if os.Getenv("MINIO_HTTPS") == "false" {
		config.minioHTTPS = "false"
	} else {
		config.minioHTTPS = "true"
	}
	config.cipherKey = os.Getenv("CIPHER_KEY")

	config.origin = "*"
	if os.Getenv("POLICY_MODEL_PATH") == "" {
		log.Fatal("No policy model provided \n")
	} else {
		config.policyModel = os.Getenv("POLICY_MODEL_PATH")
	}

	if os.Getenv("NATS_SERVER") == "" || os.Getenv("NATS_USER") == "" || os.Getenv("NATS_PWD") == "" {
		log.Fatal("No nats server and/or credenatials provided \n")
	}
	config.natsServer = os.Getenv("NATS_SERVER")
	config.natsUser = os.Getenv("NATS_USER")
	config.natsPassword = os.Getenv("NATS_PWD")

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
	case "minio-https":
		return c.minioHTTPS
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
	case "minio-server":
		return c.s3server
	case "minio-key":
		return c.awsID
	case "minio-secret":
		return c.awsSecret
	case "ws-server":
		return c.wsServer
	case "nats-server":
		return c.natsServer
	case "nats-user":
		return c.natsUser
	case "nats-pwd":
		return c.natsPassword
	}
	return ""
}

func (c *Conf) GetFromMatrix(sm string, key1 string, key2 string) (string, bool) {
	panic("not implemented")
}
func (c *Conf) SetFromDB() error {
	panic("not implemented")
}
