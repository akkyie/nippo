package nippo_test

import (
	"testing"
	"time"

	"github.com/akkyie/nippo/api"
	"github.com/akkyie/nippo/nippo"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func TestMakeNippo(t *testing.T) {
	now := time.Date(2019, 5, 1, 0, 0, 0, 0, time.Local)
	issues := []api.Issue{
		api.Issue{Title: "Hoge", URL: "http://hoge"},
		api.Issue{Title: "Fuga", URL: "http://fuga"},
	}
	actual := nippo.MakeNippo(issues, now)
	expected := `📅 日報 2019/05/01
*今日やること*
• …

*昨日やったこと*
• …
• Hoge http://hoge
• Fuga http://fuga


*業務で気づいたこと*
• …
`

	if actual != expected {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(actual, expected, false)
		t.Errorf("unexpected nippo: %v", diffs)
	}
}
