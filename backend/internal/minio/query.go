package minio

import (
	"bytes"
	"compress/gzip"
	"context"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"my-archive/backend/models"
	"strings"

	minio "github.com/minio/minio-go"
	"github.com/tealeg/xlsx"
)

func Query(pipe io.Writer, userID, secret, bucket string, query *models.Query) error {
	mc, err := ConnectToMinio(userID, secret)
	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return err
	}
	op := minio.SelectObjectOutputSerialization{}
	if query.Output == "json" || query.Output == "JSON" {
		op.JSON = &minio.JSONOutputOptions{
			RecordDelimiter: query.RecordDelimiter,
		}
	} else {
		op.CSV = &minio.CSVOutputOptions{
			RecordDelimiter: query.RecordDelimiter,
			FieldDelimiter:  query.FieldDelimiter,
		}
	}
	compressionType := minio.SelectCompressionNONE
	if strings.Contains(query.Dataset, ".gz") {
		compressionType = minio.SelectCompressionGZIP
	}

	opts := minio.SelectObjectOptions{
		Expression:     query.Query,
		ExpressionType: minio.QueryExpressionTypeSQL,
		InputSerialization: minio.SelectObjectInputSerialization{
			CompressionType: compressionType,
			CSV: &minio.CSVInputOptions{
				FileHeaderInfo:  minio.CSVFileHeaderInfoUse,
				RecordDelimiter: query.RecordDelimiter,
				FieldDelimiter:  query.FieldDelimiter,
			},
		},
		OutputSerialization: op,
	}

	reader, err := mc.SelectObjectContent(context.Background(), bucket, query.Dataset, opts)
	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return err
	}
	defer reader.Close()

	if _, err := io.Copy(pipe, reader); err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return err
	}

	return nil
}

func Preview(pipe io.Writer, userID, secret, bucket, file string) error {
	mc, err := ConnectToMinio(userID, secret)
	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return err
	}

	obj, err := mc.GetObject(bucket, file, minio.GetObjectOptions{})
	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return err
	}
	s, err := obj.Stat()
	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return err
	}
	var b2 []byte
	if s.ContentType == "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" { //Read in plaintext first
		b2, err = previewExcel(obj, s)
		if err != nil {
			log.Printf("[ERROR] %s", err.Error())
			return err
		}
	} else {
		offset := int64(1000)
		if s.Size < 1000 {
			offset = s.Size - 1
		}
		for {
			b2, err = readTopRow(obj, s, offset)

			if err != nil {
				log.Printf("[ERROR] %s", err.Error())
				return err
			}
			if len(b2) > 0 {
				break
			}
			if s.Size < offset+500 {
				offset = s.Size - 1
			} else {
				offset = offset + 500
			}
		}
	}
	pipe.Write(b2)
	return nil
}

func readTopRow(obj *minio.Object, oi minio.ObjectInfo, offset int64) ([]byte, error) {
	b := make([]byte, offset)
	_, err := obj.Read(b)
	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return nil, err
	}
	if oi.ContentType == "application/x-gzip" {
		return previewGzip(b, offset)
	}
	return previewPlaintext(b)
}

func previewGzip(b []byte, offset int64) ([]byte, error) {
	// First uncompress
	pz := make([]byte, offset)
	r := bytes.NewReader(b[0:offset])
	zr, err := gzip.NewReader(r)
	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return nil, err
	}
	defer zr.Close()

	for {
		pz, err = ioutil.ReadAll(zr)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			break
		} else if err != nil {
			log.Printf("[ERROR] %s", err.Error())
			break
		}
	}

	return previewPlaintext(pz)
}

func previewPlaintext(b []byte) ([]byte, error) {
	firstRowOfset := 0

	for k, v := range b {
		if string(v) == "\n" {
			firstRowOfset = k
			break
		}
	}
	if firstRowOfset != 0 {
		return b[0:firstRowOfset], nil
	}
	return nil, nil
}

func previewExcel(obj *minio.Object, oi minio.ObjectInfo) ([]byte, error) {
	f, err := xlsx.OpenReaderAtWithRowLimit(obj, oi.Size, 1)
	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return nil, err
	}
	var str string
	if f.Sheets[0] != nil && f.Sheets[0].Rows[0] != nil {
		for _, v := range f.Sheets[0].Rows[0].Cells {
			str = str + "," + v.String()
		}
	} else {
		return nil, errors.New("No sheets in file")
	}
	return []byte(str), nil
}
