/**
 * @Author:      leafney
 * @GitHub:      https://github.com/leafney
 * @Project:     rose-email
 * @Date:        2024-04-16 15:53
 * @Description:
 */

package remail

import (
	"errors"
	"github.com/wneessen/go-mail"
	"sync"
)

type GoMail struct {
	host     string
	port     int
	userName string
	passWord string
	debug    bool
	client   *mail.Client
	once     sync.Once
	fromAddr string
	fromName string
	to       []string
	subject  string
	bodyType int // 1:str 2:html
}

func NewGoMail(host string, port int, userName, passWord string) *GoMail {
	return &GoMail{
		host:     host,
		port:     port,
		userName: userName,
		passWord: passWord,
	}
}

func (r *GoMail) send(str string) error {
	var err error
	r.once.Do(func() {
		r.client, err = mail.NewClient(
			r.host,
			mail.WithPort(r.port),
			mail.WithSSLPort(true),
			mail.WithUsername(r.userName),
			mail.WithPassword(r.passWord),
			mail.WithSMTPAuth(mail.SMTPAuthPlain),
		)
	})
	if err != nil {
		return err
	}

	m := mail.NewMsg()

	if IsEmpty(r.fromAddr) {
		return errors.New("mail from address empty")
	}

	if IsNotEmpty(r.fromName) {
		err = m.FromFormat(r.fromName, r.fromAddr)
	} else {
		err = m.From(r.fromAddr)
	}
	if err != nil {
		return err
	}

	if len(r.to) == 0 {
		return errors.New("mail to address empty")
	}
	err = m.To(r.to...)
	if err != nil {
		return err
	}

	m.Subject(r.subject)

	if r.bodyType == 2 {
		m.SetBodyString(mail.TypeTextHTML, str)
	} else {
		m.SetBodyString(mail.TypeTextPlain, str)
	}

	err = r.client.DialAndSend(m)
	return err
}

func (r *GoMail) SetDebug(debug bool) *GoMail {
	r.debug = debug
	return r
}

func (r *GoMail) SetFrom(addr string) *GoMail {
	r.fromAddr = addr
	return r
}

func (r *GoMail) SetFromFormat(addr, name string) *GoMail {
	r.fromAddr = addr
	r.fromName = name
	return r
}

func (r *GoMail) SetTo(addr string) *GoMail {
	r.to = []string{addr}
	return r
}
func (r *GoMail) SetTos(addr []string) *GoMail {
	r.to = addr
	return r
}

func (r *GoMail) SetSubject(s string) *GoMail {
	r.subject = s
	return r
}

//func (r *GoMail) SetBodyString(body string) *GoMail {
//	r.bodyStr = body
//	return r
//}

func (r *GoMail) SendStr(str string) error {
	r.bodyType = 1
	return r.send(str)
}

func (r *GoMail) SendHtml(html string) error {
	r.bodyType = 2
	return r.send(html)
}
