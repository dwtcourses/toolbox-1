package asg

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/liamg/clinch/prompt"
	log "github.com/sirupsen/logrus"
)

func GetAutoScalingGroups() {
	client := autoscaling.New(session.Must(session.NewSession(&aws.Config{Region: aws.String("eu-west-1")})))
	groups, err := getAutoScalingGroups(client)
	if err != nil {
		panic(err)

	}

	var groupNames []string

	for _, group := range groups {
		groupNames = append(groupNames, aws.StringValue(group.AutoScalingGroupName))
	}

	_, asg, err := prompt.ChooseFromList("Select ASG", groupNames)
	if err != nil {
		panic(err)
	}
	lifecycleHook := getLifecycleHooks(client, asg)

	completeInstances(client, asg, lifecycleHook)

}

func completeInstances(client *autoscaling.AutoScaling, asgName, lifecycleHook string) {
	instances, err := client.DescribeAutoScalingInstances(&autoscaling.DescribeAutoScalingInstancesInput{})
	if err != nil {
		panic(err)
	}

	for _, instance := range instances.AutoScalingInstances {
		if aws.StringValue(instance.AutoScalingGroupName) == asgName {
			log.Infof("Completing lifecycle hook for %v", aws.StringValue(instance.InstanceId))
		}
	}
}

func getLifecycleHooks(client *autoscaling.AutoScaling, asgName string) string {
	hooks, err := client.DescribeLifecycleHooks(&autoscaling.DescribeLifecycleHooksInput{
		AutoScalingGroupName: aws.String(asgName),
	})
	if err != nil {
		panic(err)
	}
	var hookNames []string

	for _, hook := range hooks.LifecycleHooks {
		hookNames = append(hookNames, aws.StringValue(hook.LifecycleHookName))
	}

	_, lifecycleHook, err := prompt.ChooseFromList("Select Lifecycle hook", hookNames)
	if err != nil {
		panic(err)
	}
	return lifecycleHook
}

func getAutoScalingGroups(client *autoscaling.AutoScaling) ([]*autoscaling.Group, error) {
	var result []*autoscaling.Group
	var output *autoscaling.DescribeAutoScalingGroupsOutput
	next := true

	for next {
		var err error
		input := &autoscaling.DescribeAutoScalingGroupsInput{}

		if output != nil {
			input.NextToken = output.NextToken
		}

		output, err = client.DescribeAutoScalingGroups(input)
		if err != nil {
			return nil, fmt.Errorf("error getting autoscaling groups: %v", err)
		}

		next = output.NextToken != nil

		result = append(result, output.AutoScalingGroups...)
	}

	return result, nil
}
