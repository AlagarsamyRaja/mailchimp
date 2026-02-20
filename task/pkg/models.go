package pkg

type MailchimpConfig struct {
	ApiKey       string
	ServerPrefix string
	//CampaignId   string
	//ListID       string
}

type CampaignSettings struct {
	SubjectLine string `json:"subject_line"`
	Title       string `json:"title"`
	FromName    string `json:"from_name"`
	ReplyTo     string `json:"reply_to"`
}

type CampaignRecipients struct {
	ListID string `json:"list_id"`
}

type CampaignCreateRequest struct {
	Type       string             `json:"type"`
	Recipients CampaignRecipients `json:"recipients"`
	Settings   CampaignSettings   `json:"settings"`
}

type CampaignResponse struct {
	ID     string `json:"id"`
	Type   string `json:"type"`
	Status string `json:"status"`
}

type Response struct {
	Message string `json:"message"`
}
