package nippo

import (
	"fmt"
	"time"

	"github.com/akkyie/nippo/api"
)

// MakeNippo returns nippo template as a string
func MakeNippo(issues []api.Issue, now time.Time) string {
	issueList := ""
	for _, issue := range issues {
		issueList += fmt.Sprintf("• %s %s\n", issue.Title, issue.URL)
	}

	template := `📅 日報 %s
*今日やること*
• …

*昨日やったこと*
• …
%s

*業務で気づいたこと*
• …
`

	today := now.Format("2006/01/02")
	return fmt.Sprintf(template, today, issueList)
}
