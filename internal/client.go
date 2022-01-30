package sonarcheck

import "gopkg.in/resty.v1"

type SonarClient struct {
	Client    *resty.Request
	ProjectID string
}
