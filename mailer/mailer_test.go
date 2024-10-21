package mailer

import "testing"

func TestMail_SendDMTPMessage(t *testing.T) {
	mailMessage := Message{
		From:        "floridait@gmail.com",
		FromName:    "drezilla",
		To:          "tobi@heaven.com",
		Subject:     "integration test",
		Template:    "test",
		Attachments: []string{"./testdata/mail/test.html.tmpl", "./testdata/mail/test.plain.tmpl"},
	}

	if err := mailer.SendSMTPMessage(mailMessage); err != nil {
		t.Error(err)
	}
}

func TestMail_SendUsingChan(t *testing.T) {
	mailMessage := Message{
		From:        "floridait@gmail.com",
		FromName:    "drezilla",
		To:          "tobi@heaven.com",
		Subject:     "integration test",
		Template:    "test",
		Attachments: []string{"./testdata/mail/test.html.tmpl", "./testdata/mail/test.plain.tmpl"},
	}

	mailer.Jobs <- mailMessage
	res := <-mailer.Results
	if res.Error != nil {
		t.Error("failed to run mail over channels")
	}

	mailMessage.To = "not_an_email_addres"
	mailer.Jobs <- mailMessage
	res = <-mailer.Results
	if res.Error == nil {
		t.Error("no error received with invalid address")
	}
}

func TestEmail_SendUsingAPI(t *testing.T) {
	mailMessage := Message{
		From:        "floridait@gmail.com",
		FromName:    "drezilla",
		To:          "tobi@heaven.com",
		Subject:     "integration test",
		Template:    "test",
		Attachments: []string{"./testdata/mail/test.html.tmpl", "./testdata/mail/test.plain.tmpl"},
	}

	mailer.API = "unknow"
	mailer.APIKey = "abc123"
	mailer.APIUrl = "https://test-shit.com"
	// this should through the test to fail so we negate it with ==
	if err := mailer.SendUsingAPI(mailMessage, "smtp"); err == nil {
		t.Error(err)
	}
	mailer.API = ""
	mailer.APIKey = ""
	mailer.APIUrl = ""
}

func TestMail_buildHTMLMessage(t *testing.T) {
	mailMessage := Message{
		From:        "floridait@gmail.com",
		FromName:    "drezilla",
		To:          "tobi@heaven.com",
		Subject:     "integration test",
		Template:    "test",
		Attachments: []string{"./testdata/mail/test.html.tmpl", "./testdata/mail/test.html.tmpl"},
	}

	_, err := mailer.buildHTMLMessage(mailMessage)
	if err != nil {
		t.Error(err)
	}
}

func TestMail_buildPlainMessage(t *testing.T) {
	mailMessage := Message{
		From:        "floridait@gmail.com",
		FromName:    "drezilla",
		To:          "tobi@heaven.com",
		Subject:     "integration test",
		Template:    "test",
		Attachments: []string{"./testdata/mail/test.plain.tmpl", "./testdata/mail/test.plain.tmpl"},
	}

	_, err := mailer.buildPlainTextMessage(mailMessage)
	if err != nil {
		t.Error(err)
	}
}

func TestMail_send(t *testing.T) {
	mailMessage := Message{
		From:        "floridait@gmail.com",
		FromName:    "drezilla",
		To:          "tobi@heaven.com",
		Subject:     "integration test",
		Template:    "test",
		Attachments: []string{"./testdata/mail/test.plain.tmpl", "./testdata/mail/test.plain.tmpl"},
	}

	if err := mailer.Send(mailMessage); err != nil {
		t.Error(err)
	}

	mailer.API = "unknow"
	mailer.APIKey = "abc123"
	mailer.APIUrl = "https://www.fake.com"

	if err := mailer.Send(mailMessage); err == nil {
		t.Errorf("did not get an error although we should have: %s", err)
	}

	mailer.API = ""
	mailer.APIKey = ""
	mailer.APIUrl = ""
}

func TestMail_ChooseAPI(t *testing.T) {
	mailMessage := Message{
		From:        "floridait@gmail.com",
		FromName:    "drezilla",
		To:          "tobi@heaven.com",
		Subject:     "integration test",
		Template:    "test",
		Attachments: []string{"./testdata/mail/test.plain.tmpl", "./testdata/mail/test.plain.tmpl"},
	}

	mailer.API = "unknown"
	if err := mailer.ChooseAPI(mailMessage); err == nil {
		t.Error("if we dont get an error the the api chooser is not working properly")
	}
}
