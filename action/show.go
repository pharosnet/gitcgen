package action

import (
	"fmt"
)

func Show(path string) (err error) {

	r, openErr := getRepository(path)
	if openErr != nil {
		err = openErr
		return
	}

	id, tag, cmtIdErr := getLatestCommitId(r, true)
	if cmtIdErr != nil {
		err = cmtIdErr
		return
	}

	fmt.Println("id:", id, "tag:", tag)

	return
}
