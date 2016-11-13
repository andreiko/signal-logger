package main

import (
	"github.com/kdar/logrus-cloudwatchlogs"
	"github.com/aws/aws-sdk-go/aws"
)

func InstallCloudwatch(cloudwatchGroup, cloudwatchStream string) error {
	hook, err := logrus_cloudwatchlogs.NewHook(
		cloudwatchGroup,
		cloudwatchStream,
		aws.NewConfig(),
	)

	if err != nil {
		return err
	}

	logger.Hooks.Add(hook)
	return nil
}
