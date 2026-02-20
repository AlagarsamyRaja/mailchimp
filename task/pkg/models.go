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

type CampaignResponse struct {
	ID     string      `json:"id"`
	Type   string      `json:"type"`
	Status interface{} `json:"status"`
}

type Response struct {
	Message string `json:"message"`
}

// ✅ Audience Contact Information
type AudienceContact struct {
	Company  string `json:"company"`  // required
	Address1 string `json:"address1"` // required
	City     string `json:"city"`     // required
	State    string `json:"state"`    // required
	Zip      string `json:"zip"`      // required
	Country  string `json:"country"`  // required (two-letter code, e.g. "IN")
}

// ✅ Campaign Defaults for the Audience
type CampaignDefaults struct {
	FromName  string `json:"from_name"`  // required
	FromEmail string `json:"from_email"` // required
	Subject   string `json:"subject"`    // optional
	Language  string `json:"language"`   // required (e.g., "EN")
}

// ✅ Request Payload for Create/Update Audience
type AudienceRequest struct {
	Name               string           `json:"name"`                // required
	Contact            AudienceContact  `json:"contact"`             // required
	PermissionReminder string           `json:"permission_reminder"` // required
	CampaignDefaults   CampaignDefaults `json:"campaign_defaults"`   // required
	EmailTypeOption    bool             `json:"email_type_option"`   // required
}

type AudienceResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
