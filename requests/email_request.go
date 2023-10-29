package requests

type EmailRequest struct {
	From        string
	To          []string
	Cc          []string
	Bcc         []string
	Subject     string
	Content     string
	Attachments []string
}
