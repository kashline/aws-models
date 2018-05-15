package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
)

func GetAmi() {
	client := ec2metadata.New(session.New())
	test, err := client.GetMetadata("")
	if err != nil {
		log.Println(err)
	}
	log.Println(test)
}
