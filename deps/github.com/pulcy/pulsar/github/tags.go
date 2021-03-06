package github

import (
	"fmt"

	logging "github.com/op/go-logging"
)

const (
	TAGS_URI = "/repos/%s/%s/tags%s"
)

type Tag struct {
	Name       string `json:"name"`
	Commit     Commit `json:"commit"`
	ZipBallUrl string `json:"zipball_url"`
	TarBallUrl string `json:"tarball_url"`
}

func (t *Tag) String() string {
	return t.Name + " (commit: " + t.Commit.Url + ")"
}

/* get the tags associated with a repo */
func Tags(log *logging.Logger, user, repo, token string) ([]Tag, error) {
	var tags []Tag

	if token != "" {
		token = "?access_token=" + token
	}

	err := GithubGet(log, fmt.Sprintf(TAGS_URI, user, repo, token), &tags)
	if err != nil {
		return nil, err
	}

	return tags, nil
}
