package pkg

type MailchimpConfig struct {
	ApiKey       string
	ServerPrefix string
}

type CampaignSettings struct {
	SubjectLine string `json:"subject_line"`
	Title       string `json:"title"`
	FromName    string `json:"from_name"`
	ReplyTo     string `json:"reply_to"`
	TemplateId  int    `json:"template_id"`
}

type CampaignRecipients struct {
	ListID string `json:"list_id"`
}

type CampaignCreateRequest struct {
	Type       string             `json:"type"`
	Recipients CampaignRecipients `json:"recipients"`
	Settings   CampaignSettings   `json:"settings"`
}

type Response struct {
	Message string `json:"message"`
}

type AudienceContact struct {
	Company  string `json:"company"`
	Address1 string `json:"address1"`
	City     string `json:"city"`
	State    string `json:"state"`
	Zip      string `json:"zip"`
	Country  string `json:"country"`
}

type CampaignDefaults struct {
	FromName  string `json:"from_name"`
	FromEmail string `json:"from_email"`
	Subject   string `json:"subject"`
	Language  string `json:"language"`
}

type AudienceRequest struct {
	Name               string           `json:"name"`
	Contact            AudienceContact  `json:"contact"`
	PermissionReminder string           `json:"permission_reminder"`
	CampaignDefaults   CampaignDefaults `json:"campaign_defaults"`
	EmailTypeOption    bool             `json:"email_type_option"`
}

type AudienceResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type MergeFields struct {
	FNAME string `json:"fname"`
	LNAME string `json:"lname"`
}

type MemberRequest struct {
	EmailAddress string      `json:"email_address"`
	Status       string      `json:"status"`
	MergeFields  MergeFields `json:"merge_fields"`
}
