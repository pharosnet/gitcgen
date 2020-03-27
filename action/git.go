package action

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"time"
)

func getRepository(projectPath string) (r *git.Repository, err error) {

	r0, openErr := git.PlainOpen(projectPath)

	if openErr != nil {
		err = fmt.Errorf("git open at %s failed, %v", projectPath, openErr)
		return
	}

	r = r0

	return
}

func getLatestCommit(r *git.Repository) (cmt *object.Commit, err error) {

	since := time.Time{}
	until := time.Now()

	commits, logErr := r.Log(&git.LogOptions{From: plumbing.ZeroHash, Since: &since, Until: &until})
	if logErr != nil {
		err = fmt.Errorf("git log failed, %v", commits)
		return
	}

	if commits == nil {
		err = fmt.Errorf("git log failed, no commits")
		return
	}

	cmt0, cmtErr := commits.Next()
	defer commits.Close()
	if cmtErr != nil {
		err = fmt.Errorf("git get latest commit failed, %v", cmtErr)
		return
	}
	cmt = cmt0

	return
}

func getLatestCommitId(r *git.Repository, short bool) (id string, tag string, err error) {

	cmt, cmtErr := getLatestCommit(r)
	if cmtErr != nil {
		err = cmtErr
		return
	}

	id0 := cmt.Hash.String()

	if id0 == "" {
		id0 = "0000000000000000000000000000000000000000"
	}

	if short {
		id0 = id0[0:7]
	}

	id = id0

	tag, _ = getTagOnCommitId(r, cmt)

	return
}

func getTagOnCommitId(r *git.Repository, cmt *object.Commit) (name string, err error) {

	tags, tagsErr := r.Tags()
	if tagsErr != nil {
		err = tagsErr
		return
	}
	_ = tags.ForEach(func(reference *plumbing.Reference) error {
		if name != "" {
			return nil
		}
		if reference.Hash().String() == cmt.Hash.String() {
			name = reference.Name().Short()
		}
		return nil
	})

	return
}
