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

	switch e.Type {

	case "PushEvent":
		return fmt.Sprintf("- Pushed to repo: %s branch: %s by user: %s", repo, ref, actor)
	case "PullRequestEvent":
		return fmt.Sprintf("- Pull Request: %s", repo)
	case "IssuesEvent":
		return fmt.Sprintf("- Issue in repo: %s", repo)
	case "WatchEvent":
		if e.Payload.Action == "started" {
			return fmt.Sprintf("- User: %s just starred repo: %s", actor, repo)
		}
	case "CreateEvent":
		if e.Payload.RefType == "repository" {
			return fmt.Sprintf("- Created repository %s", repo)
		}
		if e.Payload.RefType != "" && ref != "" {
			return fmt.Sprintf("- Created %s '%s' in %s", e.Payload.RefType, ref, repo)
		}
		return fmt.Sprintf("- Created a new resource in %s", repo)
	case "PublicEvent":
		return fmt.Sprintf("- The Repo: %s was made public by user: %s", repo, actor)
	case "DeleteEvent":
		return fmt.Sprintf("- Deleted %s: %s in repo: %s", e.Payload.RefType, ref, repo)
	}

	return fmt.Sprintf("-[u] %s in %s", e.Type, repo)

}
