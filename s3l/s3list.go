package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
)

func main() {
	profile := flag.String("profile", "default", "Which aws profile")
	region := flag.String("region", "us-west-2", "Which aws region")
	flag.Parse()

	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithSharedConfigProfile(*profile),
		config.WithRegion(*region),
	)
	if err != nil {
		log.Fatalf("Unable to load aws config: %v", err)
	}
	s3Client := s3.NewFromConfig(cfg)
	result, err := s3Client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		log.Fatalf("Unable to List S3 buckets: %v", err)
	}
	fmt.Println("List of S3Buckets:")
	for _, bucket := range result.Buckets {
		fmt.Println("-", *bucket.Name)
	}
}
