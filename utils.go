/**
 * @Author:      leafney
 * @GitHub:      https://github.com/leafney
 * @Project:     rose-email
 * @Date:        2024-04-16 16:30
 * @Description:
 */

package remail

import "strings"

func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func IsNotEmpty(s string) bool {
	return len(strings.TrimSpace(s)) > 0
}
