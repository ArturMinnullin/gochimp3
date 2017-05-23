package gochimp3

const (
	reports_path = "/reports"
)

type ReportQueryParams struct {
	ExtendedQueryParams
}

func (api API) GetReports(params *ReportQueryParams) (*ListOfReports, error) {
	response := new(ListOfReports)

	err := api.Request("GET", reports_path, params, nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type ListOfReports struct {
	TotalItems int      `json:"total_items"`
	Reports    []Report `json:"reports"`
}

type Report struct {
	ID            string `json:"total_items"`
	CampaignTitle string `json:"campaign_title"`
	EmailsSent    int    `json:"emails_sent"`
	Opens         Opens  `json:"opens"`
	Clicks        Clicks `json:"clicks"`
}

type Opens struct {
	OpensTotal  int     `json:"opens_total"`
	UniqueOpens int     `json:"unique_opens"`
	OpenRate    float64 `json:"open_rate"`
}

type Clicks struct {
	ClicksTotal  int     `json:"clicks_total"`
	UniqueClicks int     `json:"unique_clicks"`
	ClickRate    float64 `json:"click_rate"`
}
