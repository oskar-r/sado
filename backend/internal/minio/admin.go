package minio

import (
	"encoding/json"
	"errors"
	"fmt"
	"my-archive/backend/internal/config"

	"github.com/minio/minio/pkg/madmin"
)

func CreateUser(accessKey, secret, bucket string) error {
	mdmClnt, err := madmin.New(config.Get("minio-server"), config.Get("minio-key"), config.Get("minio-secret"), true)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = mdmClnt.AddUser(accessKey, secret)
	if err != nil {
		fmt.Println(err)
		return err
	}
	dp := DefaultPolicy(bucket)
	policy, err := json.Marshal(&dp)
	if err != nil {
		fmt.Println(err)
		return errors.New("Could not create policy for user")
	}

	err = mdmClnt.AddCannedPolicy(accessKey, string(policy))
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = mdmClnt.SetUserPolicy(accessKey, accessKey)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
