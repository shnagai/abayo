package main

import (
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
)

func main() {
	var repository string

	flag.StringVar(&repository, "r", "", "repository name")
	// parse flag value
	flag.Parse()
	// new aws session
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(endpoints.ApNortheast1RegionID),
	}))

	// ecr connection
	svc := ecr.New(sess)

	input := &ecr.ListImagesInput{
		RepositoryName: aws.String(repository),
		Filter: &ecr.ListImagesFilter{
			TagStatus: aws.String("UNTAGGED"),
		},
	}

	result, err := svc.ListImages(input)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ecr.ErrCodeServerException:
				fmt.Println(ecr.ErrCodeServerException, aerr.Error())
			case ecr.ErrCodeInvalidParameterException:
				fmt.Println(ecr.ErrCodeInvalidParameterException, aerr.Error())
			case ecr.ErrCodeRepositoryNotFoundException:
				fmt.Println(ecr.ErrCodeRepositoryNotFoundException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}
	// delete listを作る
	var delete_list []int

	for num, res := range result.ImageIds {
		delete_image_input := &ecr.BatchDeleteImageInput{
			ImageIds: []*ecr.ImageIdentifier{
				{
					ImageDigest: aws.String(*res.ImageDigest),
				},
			},
			RepositoryName: aws.String(repository),
		}
		// 配列に追加していく(最後に数を出すため)
		delete_list = append(delete_list, num)

		_, err := svc.BatchDeleteImage(delete_image_input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case ecr.ErrCodeServerException:
					fmt.Println(ecr.ErrCodeServerException, aerr.Error())
				case ecr.ErrCodeInvalidParameterException:
					fmt.Println(ecr.ErrCodeInvalidParameterException, aerr.Error())
				case ecr.ErrCodeRepositoryNotFoundException:
					fmt.Println(ecr.ErrCodeRepositoryNotFoundException, aerr.Error())
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				fmt.Println(err.Error())
			}
			return
		}
		//number := num
	}

	fmt.Printf("あばよ!!%d個のイメージ達!!\n", len(delete_list))
}
