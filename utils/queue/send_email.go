package queue

import (
	Email "evermos/utils/email"

	"github.com/gocraft/work"
)

func (c *Context) SendEmail(job *work.Job) error {
	subject := job.ArgString("subject")
	email := job.ArgString("email")
	message := job.ArgString("message")
	//"Hello, <b>have a nice day</b>"
	if err := job.ArgError(); err != nil {
		return err
	}
	Email.Send(email, subject, message)
	//Email.Send([]string{email}, subject)
	return nil
}
