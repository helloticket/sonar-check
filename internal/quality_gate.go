package sonarcheck

type QualityGateProjectStatus struct {
	ProjectStatus struct {
		Status     string `json:"status"`
		Conditions []struct {
			Status         string `json:"status"`
			MetricKey      string `json:"metricKey"`
			Comparator     string `json:"comparator"`
			PeriodIndex    int    `json:"periodIndex"`
			ErrorThreshold string `json:"errorThreshold"`
			ActualValue    string `json:"actualValue"`
		} `json:"conditions"`
		Periods []struct {
			Index int    `json:"index"`
			Mode  string `json:"mode"`
			Date  string `json:"date"`
		} `json:"periods"`
		IgnoredConditions bool `json:"ignoredConditions"`
		Period            struct {
			Mode string `json:"mode"`
			Date string `json:"date"`
		} `json:"period"`
	} `json:"projectStatus"`
}
