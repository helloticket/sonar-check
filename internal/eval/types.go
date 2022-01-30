package eval

import (
	"errors"

	sonarcheck "github.com/helloticket/sonar-check/internal"
)

type evalBugVulnerability struct {
	cli *sonarcheck.SonarClient
}

func NewEvalBugVulnerability(cli *sonarcheck.SonarClient) *evalBugVulnerability {
	return &evalBugVulnerability{cli: cli}
}

func (c *evalBugVulnerability) Run() (sonarcheck.FacetStatus, error) {
	issues, err := c.findIssues()
	if err != nil {
		return sonarcheck.FacetStatus{}, err
	}

	quality, err := c.findQualityGate()
	if err != nil {
		return sonarcheck.FacetStatus{}, err
	}

	status := sonarcheck.FacetStatus{}

	for _, i := range issues.Facets {
		if i.Property == "types" {
			for _, v := range i.Values {
				if v.Val == "BUG" {
					if v.Count > 0 {
						status.Bugs = status.Bugs + v.Count
					}
				}

				if v.Val == "VULNERABILITY" {
					if v.Count > 0 {
						status.Vulnerability = status.Vulnerability + v.Count
					}
				}

				if v.Val == "CODE_SMELL" {
					if v.Count > 0 {
						status.CodeSmell = status.CodeSmell + v.Count
					}
				}
			}
		}
	}

	status.QualityGateStatus = quality.ProjectStatus.Status

	if quality.ProjectStatus.Status == "ERROR" {
		status.QualityGateError = true
	}

	return status, nil
}

func (c *evalBugVulnerability) findIssues() (sonarcheck.Issue, error) {
	resp, err := c.cli.Client.
		SetResult(sonarcheck.Issue{}).
		SetQueryParams(map[string]string{
			"componentKeys": c.cli.ProjectID,
			"facets":        "types",
		}).
		Get("/api/issues/search")

	if err != nil {
		return sonarcheck.Issue{}, err
	}

	if resp.IsError() {
		return sonarcheck.Issue{}, errors.New(resp.String())
	}

	issues := resp.Result().(*sonarcheck.Issue)

	return *issues, nil
}

func (c *evalBugVulnerability) findQualityGate() (sonarcheck.QualityGateProjectStatus, error) {
	resp, err := c.cli.Client.
		SetResult(sonarcheck.QualityGateProjectStatus{}).
		SetQueryParams(map[string]string{
			"projectKey": c.cli.ProjectID,
		}).
		Get("/api/qualitygates/project_status")

	if err != nil {
		return sonarcheck.QualityGateProjectStatus{}, err
	}

	if resp.IsError() {
		return sonarcheck.QualityGateProjectStatus{}, errors.New(resp.String())
	}

	status := resp.Result().(*sonarcheck.QualityGateProjectStatus)

	return *status, nil
}
