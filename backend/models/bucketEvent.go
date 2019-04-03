package models

import "time"

type BucketEvent struct {
	EventName string `json:"EventName,omitempty"`
	Key       string `json:"Key,omitempty"`
	Records   []struct {
		EventVersion string     `json:"eventVersion,omitempty"`
		EventSource  string     `json:"eventSource,omitempty"`
		AWSRegion    string     `json:"awsRegion,omitempty"`
		EventTime    *time.Time `json:"eventTime,omitempty"`
		EventName    string     `json:"eventName,omitempty"`
		UserIdentity struct {
			PrincipalID string `json:"principalId,omitempty"`
		} `json:"userIdentity,omitempty"`
		RequestParameters struct {
			AccessKey       string `json:"accessKey,omitempty"`
			Region          string `json:"region,omitempty"`
			SourceIPAddress string `json:"sourceIPaddress,omitempty"`
		} `json:"request_parameters,omitempty"`
		S3 struct {
			Bucket struct {
				Name          string `json:"name,omitempty"`
				OwnerIdentity struct {
					PrincipalID string `json:"principalId,omitempty"`
				} `json:"ownerIdentity,omitempty"`
				Arn string `json:"arn,omitempty"`
			} `json:"bucket,omitempty"`
			Object struct {
				Key          string `json:"key,omitempty"`
				Size         int64  `json:"size,omitempty"`
				ContentType  string `json:"contentType,omitempty"`
				UserMetadata struct {
					XAmzMetaContentType string `json:"X-Amz-Meta-Content_type,omitempty"` //"X-Amz-Meta-Content_type":"application/zip",
					XAmzMetaUserID      string `json:"X-Amz-Meta-User_id,omitempty"`      //"X-Amz-Meta-User_id":"72f91914-f3da-4f89-bc3a-b12fb9444cda",
					XAmzMetaUsername    string `json:"X-Amz-Meta-Username,omitempty"`     //"X-Amz-Meta-Username":"oskar",
				} `json:"userMetadata,omitempty"`
				VersionID string `json:"versionId,omitempty"`
				Sequencer string `json:"sequencer,omitempty"`
			} `json:"object,omitempty"`
		} `json:"s3,omitempty"`
	} `json:"records,omitempty"`
}

type EventNotification struct {
	Name        string     `json:"name,omitempty"`
	ContentType string     `json:"content_type,omitempty"`
	Size        int64      `json:"size,omitempty"`
	Time        *time.Time `json:"time,omitempty"`
}

/*
{
	"EventName": "s3:ObjectCreated:Put",
	"Key": "my-bucket/index.html",
	"Records": [{
		"eventVersion": "2.0",
		"eventSource": "minio:s3",
		"awsRegion": "",
		"eventTime": "2019-04-03T20:32:10Z",
		"eventName": "s3:ObjectCreated:Put",
		"userIdentity": {
			"principalId": "72f91914-f3da-4f89-bc3a-b12fb9444cda"
		},
		"requestParameters": {
			"accessKey": "72f91914-f3da-4f89-bc3a-b12fb9444cda",
			"region": "",
			"sourceIPAddress": "172.19.0.5"
		},
		"responseElements": {
			"x-amz-request-id": "1592110406E9D1CC",
			"x-minio-deployment-id": "7d558291-1ad0-4d90-9beb-8cc4434598b6",
			"x-minio-origin-endpoint": "http://172.19.0.4:9000"
		},
		"s3": {
			"s3SchemaVersion": "1.0",
			"configurationId": "Config",
			"bucket": {
				"name": "my-bucket",
				"ownerIdentity": {
					"principalId": "72f91914-f3da-4f89-bc3a-b12fb9444cda"
				},
				"arn": "arn:aws:s3:::my-bucket"
			},
			"object": {
				"key": "index.html",
				"size": 295,
				"eTag": "d66c38317b6a28bdb1daceecddcbceb5",
				"contentType": "text/html",
				"userMetadata": {
					"X-Amz-Meta-Content_type": "text/html",
					"X-Amz-Meta-User_id": "72f91914-f3da-4f89-bc3a-b12fb9444cda",
					"X-Amz-Meta-Username": "oskar",
					"content-type": "text/html"
				},
				"versionId": "1",
				"sequencer": "1592110407485774"
			}
		},
		"source": {
			"host": "172.19.0.5",
			"port": "",
			"userAgent": "Minio (linux; amd64) minio-go/v6.0.21"
		}
	}]
}
*/
