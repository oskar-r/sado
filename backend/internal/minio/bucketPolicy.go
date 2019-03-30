package minio

type Stmt struct {
	Effect    string   `json:"Effect,omitempty"`
	Principal Princ    `json:"Principal,omitempty"`
	Action    []string `json:"Action,omitempty"`
	Resource  []string `json:"Resource,omitempty"`
}
type Princ struct {
	AWS []string `json:"AWS,omitempty"`
}

type BucketPolicy struct {
	Version   string `json:"Version,omitempty"`
	Statement []Stmt `json:"Statement,omitempty"`
}

func DefaultPolicy(bucketName string) BucketPolicy {
	return BucketPolicy{
		Version: "2012-10-17",
		Statement: []Stmt{
			{
				Effect: "Allow",
				Principal: Princ{
					AWS: []string{"*"},
				},
				Action:   []string{"s3:GetBucketLocation", "s3:ListBucket", "s3:ListBucketMultipartUploads"},
				Resource: []string{"arn:aws:s3:::" + bucketName},
			},
			{
				Effect: "Allow",
				Principal: Princ{
					AWS: []string{"*"},
				},
				Action:   []string{"s3:AbortMultipartUpload", "s3:DeleteObject", "s3:GetObject", "s3:ListMultipartUploadParts", "s3:PutObject"},
				Resource: []string{"arn:aws:s3:::" + bucketName + "/*"},
			},
		},
	}
}
