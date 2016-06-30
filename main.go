package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {

	var bucketName = flag.String("bucket", "", "bucket name")
	var artifactName = flag.String("name", "", "object name")
	var acl = flag.String("acl", "", "S3 ACL to set on object")
	flag.Parse()

	if len(*bucketName) == 0 {
		fmt.Fprintln(os.Stderr, "Must pass a bucket name")
		os.Exit(1)
	}

	if len(*artifactName) == 0 {
		fmt.Fprintln(os.Stderr, "Must pass name of artifact")
		os.Exit(1)
	}

	var body []byte
	if len(flag.Args()) > 0 {
		var err error
		body, err = ioutil.ReadFile(flag.Arg(0))
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	}

	svc := s3.New(session.New())

	params := &s3.PutObjectInput{
		Bucket: aws.String(*bucketName),
		Key:    aws.String(*artifactName),
		Body:   bytes.NewReader(body),
	}

	if len(*acl) > 0 {
		params.ACL = aws.String(*acl)
	}

	_, err := svc.PutObject(params)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
