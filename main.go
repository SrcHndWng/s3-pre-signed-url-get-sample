package main

import (
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	region := os.Args[1]
	bucket := os.Args[2]
	key := os.Args[3]

	log.Printf("your region = %s, bucket = %s, key = %s\n", region, bucket, key)

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)

	svc := s3.New(sess)

	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	urlStr, err := req.Presign(2 * time.Minute)

	if err != nil {
		log.Println("Failed to sign request", err)
	}

	log.Println("Pre signed url : ", urlStr)
}
