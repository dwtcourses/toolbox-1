package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
	"github.com/liamg/clinch/task"
	"github.com/owenrumney/clinch/prompt"
	"strings"
)

type CloudWatchEvents struct {
	client *cloudwatchevents.CloudWatchEvents
}

type EventFlags struct {
	ShowEnabled  bool
	ShowDisabled bool
}

func NewCloudWatchEventClient() *CloudWatchEvents {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	return &CloudWatchEvents{
		client: cloudwatchevents.New(sess),
	}
}

func (c *CloudWatchEvents) GetEvents(filter string, flags EventFlags) {
	ruleNames, rules := c.getEvents(filter, flags)
	if len(rules) == 0 {
		println("No choices returned.")
		return
	}
	index, _, _ := prompt.ChooseFromList("Rules", ruleNames)
	if index == -1 {
		fmt.Println("Cancelled")
		return
	}

	chosenRule := rules[index]
	c.showRuleOptions(*chosenRule)

}

func (c *CloudWatchEvents) getEvents(filter string, flags EventFlags) ([]string, []*cloudwatchevents.Rule) {
	var rulesOutput *cloudwatchevents.ListRulesOutput
	var ruleNames []string
	var err error
	var rules []*cloudwatchevents.Rule
	showEnabled := flags.ShowEnabled == flags.ShowDisabled || flags.ShowEnabled
	showDisabled := flags.ShowEnabled == flags.ShowDisabled || flags.ShowDisabled
	more := true

	for more {
		input := &cloudwatchevents.ListRulesInput{}

		if rulesOutput != nil {
			input.NextToken = rulesOutput.NextToken
		}

		rulesOutput, err = c.client.ListRules(input)
		if err != nil {
			panic(err)
		}
		for _, rule := range rulesOutput.Rules {
			ruleName := aws.String(*rule.Name)
			if strings.Contains(*ruleName, filter) &&
				((showEnabled && string(*rule.State) == "ENABLED") ||
					(showDisabled && string(*rule.State) == "DISABLED")) {
				ruleNames = append(ruleNames, *ruleName)
				rules = append(rules, rule)
			}
		}
		more = rulesOutput.NextToken != nil
	}

	return ruleNames, rules
}

func (c *CloudWatchEvents) showRuleOptions(rule cloudwatchevents.Rule) {
	fmt.Println("Rule Name: ", string(*rule.Name))
	fmt.Println("Rule Description: ", string(*rule.Description))
	state := string(*rule.State)
	fmt.Println("Rule State: ", state)

	toggle := prompt.EnterInputWithDefault("\nToggle rule", "no")

	if strings.ToLower(toggle) == "yes" {
		if state == "ENABLED" {
			c.toggleRule(rule, true)
		} else {
			c.toggleRule(rule, false)
		}
	}
}

func (c *CloudWatchEvents) toggleRule(rule cloudwatchevents.Rule, enable bool) {
	if enable {
		_, err := c.client.EnableRule(&cloudwatchevents.EnableRuleInput{
			Name:         rule.Name,
			EventBusName: rule.EventBusName,
		})
		if err != nil {
			panic(err)
		}
	}
}

func (c *CloudWatchEvents) DisableEvents(filter string) {
	flags := &EventFlags{
		ShowEnabled: true,
	}
	ruleNames, rules := c.getEvents(filter, *flags)
	ids, _, _ := prompt.ChooseFromMultiList("Select events to snooze", ruleNames)
	var toDisable []cloudwatchevents.Rule
	for _, i := range ids {
		toDisable = append(toDisable, *rules[i])
	}
	c.disableEvents(toDisable)
}

func (c *CloudWatchEvents) EnableEvents(filter string) {
	flags := &EventFlags{
		ShowDisabled: true,
	}
	ruleNames, rules := c.getEvents(filter, *flags)
	ids, _, _ := prompt.ChooseFromMultiList("Select events to snooze", ruleNames)
	var toEnable []cloudwatchevents.Rule
	for _, i := range ids {
		toEnable = append(toEnable, *rules[i])
	}
	c.enableEvents(toEnable)
}

func (c *CloudWatchEvents) enableEvents(rules []cloudwatchevents.Rule) {
	for _, rule := range rules {
		err := task.New(
			"enabling",
			fmt.Sprintf("%v", aws.StringValue(rule.Name)),
			func(t *task.Task) error {
				_, err := c.client.EnableRule(&cloudwatchevents.EnableRuleInput{
					EventBusName: rule.EventBusName,
					Name:         rule.Name,
				})
				return err
			},
		).Run()
		if err != nil {
			panic(err)
		}
	}
}

func (c *CloudWatchEvents) disableEvents(rules []cloudwatchevents.Rule) {
	for _, rule := range rules {
		err := task.New(
			"disabling",
			fmt.Sprintf("%v", aws.StringValue(rule.Name)),
			func(t *task.Task) error {
				_, err := c.client.DisableRule(&cloudwatchevents.DisableRuleInput{
					EventBusName: rule.EventBusName,
					Name:         rule.Name,
				})
				return err
			},
		).Run()
		if err != nil {
			panic(err)
		}
	}
}
