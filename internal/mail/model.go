package mail

type Mail struct {
	From string
	To   []string
	CC   []string

	Subject     string
	Body        string
	Attach      string
	Attachments []string
}
