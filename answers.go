package main

import (
	"bytes"
	"gopkg.in/AlecAivazis/survey.v1"
	"strings"
)

type Answers struct {
	Type    string
	Scope   string
	Subject string
	Body    string
	Footer  string
}

func (answers *Answers) AssembleIntoMessage(buf *bytes.Buffer) {
	buf.WriteString(answers.Type)
	if answers.Scope != "" {
		buf.WriteString("(" + strings.TrimSpace(answers.Scope) + ")")
	}
	buf.WriteString(": " + strings.TrimSpace(answers.Subject))
	if answers.Body != "" {
		buf.WriteString("\n\n" + answers.Body)
	}
	if answers.Footer != "" {
		buf.WriteString("\n\nCloses " + answers.Footer)
	}
}

func typeFromOption(option interface{}) interface{} {
	text, ok := option.(string)
	if !ok {
		return nil
	}
	return strings.Split(text, ":")[0]
}

// the questions to ask
var commitQs = []*survey.Question{
	// type of header
	{
		Name: "type",
		Prompt: &survey.Select{
			Message: "Select the type of change you are committing:",
			Options: []string{
				"feat:      A new feature",
				"fix:       A bug fix",
				"docs:      Documentation only changes",
				"style:     Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)",
				"refactor:  A code change that neither fixes a bug nor adds a feature",
				"perf:      A code change that improves performance",
				"test:      Adding missing or correcting existing tests",
				"build:     Changes that affect the build system or external dependencies (example scopes: gulp, broccoli, npm, docker)",
				"ci:        Changes to our CI configuration files and scripts (example scopes: travis, circle, gitlab",
				"chore:     Changes to the build process or auxiliary tools and libraries such as documentation generation",
				"revert:    Reverts a previous commit",
			},
			PageSize: 11,
		},
		Transform: typeFromOption,
		Validate:  survey.Required,
	},
	// scope of header
	{
		Name: "scope",
		Prompt: &survey.Input{
			Message: "Scope. Could be anything specifying place of the commit change (users, db, poll):",
		},
		Validate: survey.MinLength(0),
	},
	// subject of header
	{
		Name: "subject",
		Prompt: &survey.Input{
			Message: "Subject. Concise description of the changes. Imperative, lower case and no final dot:",
		},
		Validate: survey.Required,
	},
	// body
	{
		Name: "body",
		Prompt: &survey.Multiline{
			Message: "Body. Motivation for the change and contrast this with previous behavior:",
		},
	},
	// footer
	{
		Name: "footer",
		Prompt: &survey.Multiline{
			Message: "Closed issues devided comma. Example(AT-21,AT-22):",
		},
	},
}

func AskForCommitMessage(answers *Answers) error {
	return survey.Ask(commitQs, answers)
}
