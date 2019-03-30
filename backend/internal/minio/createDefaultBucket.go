package minio

import (
	"errors"
	"log"
	"my-archive/backend/internal/config"

	minio "github.com/minio/minio-go"
)

func CreateDefaultBucket(bucket string) error {

	mc, err := ConnectToMinio(config.Get("minio-key"), config.Get("minio-secret"))
	if err != nil {
		return err
	}

	lb, err := BucketExist(mc, bucket)
	if err != nil {
		return err
	}
	if lb {
		return errors.New("bucket already exits")
	}

	err = mc.MakeBucket(bucket, "")
	if err != nil {
		log.Printf("[ERROR] could not create bucket: %+v", err)
	}
	return err
}

func DeleteBucket(bucket string) error {
	mc, err := ConnectToMinio(config.Get("minio-key"), config.Get("minio-secret"))
	if err != nil {
		return err
	}
	lb, err := BucketExist(mc, bucket)
	if err != nil {
		return err
	}
	if !lb {
		return errors.New("bucket does not exist")
	}

	err = mc.RemoveBucket(bucket)
	if err != nil {
		log.Printf("[ERROR] could not delete bucket: %+v", err)
	}
	return err
}

func BucketExist(mc *minio.Client, bucket string) (bool, error) {
	lb, err := mc.BucketExists(bucket)
	if err != nil {
		log.Printf("[ERROR] Could not check for bucket existance: %+v", err)
		return false, err
	}
	return lb, err
}
