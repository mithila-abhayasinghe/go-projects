package internal

import (
	"fmt"
	"strings"
	"time"
)

/* This is the subset of GitHub events cli will support

PushEvent:
PullRequestEvent:
IssuesEvent:
WatchEvent:
CreateEvent
PublicEvent:

*/

type GitHubEvent struct {
	Id        string    `json:"id"`
	Type      string    `json:"type"`
	Actor     Actor     `json:"actor"`
	Repo      Repo      `json:"repo"`
	Payload   Payload   `json:"payload"`
	CreatedAt time.Time `json:"created_at"`
}

type Repo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Actor struct {
	Id           int    `json:"id"`
	Login        string `json:"login"`
	DisplayLogin string `json:"display_login"`
}

type Payload struct {
	Action  string `json:"action"`
	Ref     string `json:"ref"`
	RefType string `json:"ref_type"`
}

func (e GitHubEvent) String() string {
	actor := e.Actor.Login
	repo := e.Repo.Name
	ref := strings.TrimPrefix(e.Payload.Ref, "refs/heads/")

	timeStamp := e.CreatedAt.Format("2006-01-02")

	switch e.Type {

	case "PushEvent":
		return fmt.Sprintf("- Pushed to repo: %s branch: %s by user: %s at: %s", repo, ref, actor, timeStamp)
	case "PullRequestEvent":
		return fmt.Sprintf("- Pull Request: %s at: %s", repo, timeStamp)
	case "IssuesEvent":
		return fmt.Sprintf("- Issue in repo: %s at: %s", repo, timeStamp)
	case "WatchEvent":
		if e.Payload.Action == "started" {
			return fmt.Sprintf("- User: %s just starred repo: %s at: %s", actor, repo, timeStamp)
		}
	case "CreateEvent":
		if e.Payload.RefType == "repository" {
			return fmt.Sprintf("- Created repository %s at: %s", repo, timeStamp)
		}
		if e.Payload.RefType != "" && ref != "" {
			return fmt.Sprintf("- Created %s '%s' in %s at: %s", e.Payload.RefType, ref, repo, timeStamp)
		}
		return fmt.Sprintf("- Created a new resource in %s at: %s", repo, timeStamp)
	case "PublicEvent":
		return fmt.Sprintf("- The Repo: %s was made public by user: %s at: %s", repo, actor, timeStamp)
	case "DeleteEvent":
		return fmt.Sprintf("- Deleted %s: %s in repo: %s at: %s", e.Payload.RefType, ref, repo, timeStamp)
	}

	return fmt.Sprintf("-[u] %s in %s at: %s", e.Type, repo, timeStamp)

}
