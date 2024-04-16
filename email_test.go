/**
 * @Author:      leafney
 * @GitHub:      https://github.com/leafney
 * @Project:     rose-email
 * @Date:        2024-04-16 16:40
 * @Description:
 */

package remail

import "testing"

func TestNewGoMail(t *testing.T) {
	m := NewGoMail(
		"smtp.exmail.qq.com",
		465,
		"",
		"",
	)

	err := m.
		SetFrom("").
		SetTo("").
		SetSubject("测试邮件发送").
		SendStr("这是一条测试邮件")
	if err != nil {
		t.Log(err)
	}
}
