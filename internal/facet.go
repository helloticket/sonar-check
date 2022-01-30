package sonarcheck

type FacetStatus struct {
	CodeSmell         int
	Vulnerability     int
	Bugs              int
	QualityGateError  bool
	QualityGateStatus string
}

func (f FacetStatus) Failed() bool {
	return f.Bugs > 0 || f.Vulnerability > 0 || f.QualityGateError
}
