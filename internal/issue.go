package sonarcheck

type Issue struct {
	Total  int `json:"total"`
	P      int `json:"p"`
	Ps     int `json:"ps"`
	Paging struct {
		PageIndex int `json:"pageIndex"`
		PageSize  int `json:"pageSize"`
		Total     int `json:"total"`
	} `json:"paging"`
	EffortTotal int `json:"effortTotal"`
	Issues      []struct {
		Key       string `json:"key"`
		Rule      string `json:"rule"`
		Severity  string `json:"severity"`
		Component string `json:"component"`
		Project   string `json:"project"`
		Line      int    `json:"line"`
		Hash      string `json:"hash"`
		TextRange struct {
			StartLine   int `json:"startLine"`
			EndLine     int `json:"endLine"`
			StartOffset int `json:"startOffset"`
			EndOffset   int `json:"endOffset"`
		} `json:"textRange"`
		Flows              []interface{} `json:"flows"`
		Status             string        `json:"status"`
		Message            string        `json:"message"`
		Effort             string        `json:"effort"`
		Debt               string        `json:"debt"`
		Assignee           string        `json:"assignee"`
		Author             string        `json:"author"`
		Tags               []interface{} `json:"tags"`
		CreationDate       string        `json:"creationDate"`
		UpdateDate         string        `json:"updateDate"`
		Type               string        `json:"type"`
		ExternalRuleEngine string        `json:"externalRuleEngine"`
		Scope              string        `json:"scope"`
	} `json:"issues"`
	Components []struct {
		Key       string `json:"key"`
		Enabled   bool   `json:"enabled"`
		Qualifier string `json:"qualifier"`
		Name      string `json:"name"`
		LongName  string `json:"longName"`
		Path      string `json:"path,omitempty"`
	} `json:"components"`
	Facets []struct {
		Property string `json:"property"`
		Values   []struct {
			Val   string `json:"val"`
			Count int    `json:"count"`
		} `json:"values"`
	} `json:"facets"`
}
