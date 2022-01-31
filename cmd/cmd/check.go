package cmd

import (
	"errors"
	"fmt"
	"os"

	sonarcheck "github.com/helloticket/sonar-check/internal"
	"github.com/helloticket/sonar-check/internal/eval"
	"github.com/urfave/cli/v2"
	"gopkg.in/resty.v1"
)

var msg = `
----------------------------------------
check failed
----------------------------------------
quality gates status (%v)
bugs (%v)
code smell (%v)
vulnerability (%v)
----------------------------------------
`

var Check = &cli.Command{
	Name: "check",
	Subcommands: []*cli.Command{
		{
			Name: "quality",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "url",
					Aliases: []string{"u"},
				},
				&cli.StringFlag{
					Name:    "token",
					Aliases: []string{"t"},
				},
				&cli.StringFlag{
					Name:    "project_id",
					Aliases: []string{"p"},
				},
			},
			Action: func(c *cli.Context) error {
				request, err := makeClient(c)
				if err != nil {
					return err
				}

				if skip := os.Getenv("SONARQUBE_SKIP"); skip == "true" {
					fmt.Println("skipped check")
					return nil
				}

				chk := eval.NewEvalBugVulnerability(request)

				status, err := chk.Run()
				if err != nil {
					return err
				}

				if status.Failed() {
					return fmt.Errorf(msg, status.QualityGateStatus, status.Bugs, status.CodeSmell, status.Vulnerability)
				}

				fmt.Println("check successful")
				return nil
			},
		},
	},
}

func makeClient(c *cli.Context) (*sonarcheck.SonarClient, error) {
	host := c.String("url")
	if host == "" {
		host = os.Getenv("SONARQUBE_URL")
	}

	token := c.String("token")
	if token == "" {
		token = os.Getenv("SONARQUBE_TOKEN")
	}

	projectId := c.String("project_id")
	if projectId == "" {
		host = os.Getenv("SONARQUBE_PROJECT_ID")
	}

	if host == "" || token == "" || projectId == "" {
		return nil, errors.New("token, url or project is empty")
	}

	cli := resty.New()
	cli.SetHostURL(host)
	cli.SetBasicAuth(token, "")

	return &sonarcheck.SonarClient{
		Client:    cli.R(),
		ProjectID: projectId,
	}, nil
}
