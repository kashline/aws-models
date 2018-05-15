package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// NewInstance creates a new instance based on the params provided
func NewInstance(imageID, instanceType, name string, minCount, maxCount int64) {
	svc := ec2.New(session.New())
	runResult, err := svc.RunInstances(&ec2.RunInstancesInput{
		ImageId:      aws.String(imageID),
		InstanceType: aws.String(instanceType),
		MinCount:     aws.Int64(minCount),
		MaxCount:     aws.Int64(maxCount),
	})
	if err != nil {
		log.Println("Failed to create instances ", err)
	}

	log.Println("Created instance: ", runResult.Instances[0].InstanceId)

}

// TagInstance tags the given instance(s) with the name provided
func TagInstance(runResult *ec2.Reservation, name string) {
	svc := ec2.New(session.New())
	_, err := svc.CreateTags(&ec2.CreateTagsInput{
		Resources: []*string{runResult.Instances[0].InstanceId},
		Tags: []*ec2.Tag{
			{
				Key:   aws.String("Name"),
				Value: aws.String(name),
			},
		},
	})
	if err != nil {
		log.Println(err)
	}
}

// RemoveInstances terminates the specified instances
func RemoveInstances() {

}
