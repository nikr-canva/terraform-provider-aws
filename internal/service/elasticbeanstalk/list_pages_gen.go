// Code generated by "internal/generate/listpages/main.go -ListOps=DescribeApplicationVersions,DescribeEnvironments"; DO NOT EDIT.

package elasticbeanstalk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
)

func describeApplicationVersionsPages(ctx context.Context, conn *elasticbeanstalk.Client, input *elasticbeanstalk.DescribeApplicationVersionsInput, fn func(*elasticbeanstalk.DescribeApplicationVersionsOutput, bool) bool, optFns ...func(*elasticbeanstalk.Options)) error {
	for {
		output, err := conn.DescribeApplicationVersions(ctx, input, optFns...)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextToken) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextToken = output.NextToken
	}
	return nil
}
func describeEnvironmentsPages(ctx context.Context, conn *elasticbeanstalk.Client, input *elasticbeanstalk.DescribeEnvironmentsInput, fn func(*elasticbeanstalk.DescribeEnvironmentsOutput, bool) bool, optFns ...func(*elasticbeanstalk.Options)) error {
	for {
		output, err := conn.DescribeEnvironments(ctx, input, optFns...)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextToken) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextToken = output.NextToken
	}
	return nil
}
