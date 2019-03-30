package minio

import (
	"errors"
	"log"
	"my-archive/backend/internal/utility"
	"my-archive/backend/models"

	minio "github.com/minio/minio-go"
)

var DataSetContentTypes = []string{"text/csv", "test/tsv", "application/json"}

func ListDatasetsInBucket(userID, secret, bucket string) ([]models.DataSet, error) {
	mc, err := ConnectToMinio(userID, secret)
	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return nil, err
	}
	ok, err := BucketExist(mc, bucket)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("bucket does not exist")
	}

	// Create a done channel to control 'ListObjects' go routine.
	doneCh := make(chan struct{})

	// Indicate to our routine to exit cleanly upon return.
	defer close(doneCh)

	isRecursive := true
	var objects []minio.ObjectInfo

	objectCh := mc.ListObjectsV2(bucket, "", isRecursive, doneCh)
	for object := range objectCh {
		if object.Err != nil {
			log.Printf("[ERROR] %s", object.Err.Error())
			return nil, err
		}
		oi, err := mc.StatObject(bucket, object.Key, minio.StatObjectOptions{})
		if err != nil {
			log.Printf("[ERROR] %s", err.Error())
			return nil, err
		}
		objects = append(objects, oi)
	}

	var ds []models.DataSet
	for _, v := range objects {
		cat := "document"
		if utility.Contains(v.ContentType, DataSetContentTypes) {
			cat = "dataset"
		}

		ds = append(ds, models.DataSet{
			Name:         v.Key,
			ContentType:  v.ContentType,
			Size:         v.Size,
			Category:     cat,
			LastModified: v.LastModified,
		})
	}
	return ds, nil
}
