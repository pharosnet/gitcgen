package action

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path"
	"strings"
)

func Generate(path0 string, short bool, output string) (err error) {

	id, tag, idErr := getCommitId(path0, short)
	if idErr != nil {
		err = idErr
		return
	}

	genErr := generateFile(id, tag, output)
	if genErr != nil {
		err = genErr
		return
	}

	return
}

func getCommitId(path string, short bool) (id string, tag string, err error) {

	r, openErr := getRepository(path)
	if openErr != nil {
		err = openErr
		return
	}

	id0, tag0, cmtIdErr := getLatestCommitId(r, short)
	if cmtIdErr != nil {
		err = cmtIdErr
		return
	}

	id = id0
	tag = tag0

	return
}

func generateFile(id string, tag string, path0 string) (err error) {

	dir0 := path.Dir(path0)
	goPkgName := strings.TrimSpace(dir0[strings.LastIndex(dir0, "/")+1:])

	b := bytes.NewBufferString("")
	b.WriteString(fmt.Sprintf("package %s", goPkgName))
	b.WriteString("\n")
	b.WriteString("\n")

	b.WriteString(fmt.Sprintf(`var gitVersion = "%s"`, id))
	b.WriteString("\n")
	b.WriteString("\n")

	b.WriteString("func GitVersion() string {")
	b.WriteString("\n")
	b.WriteString("	return gitVersion")
	b.WriteString("\n")
	b.WriteString("}")
	b.WriteString("\n")
	b.WriteString("\n")

	b.WriteString(fmt.Sprintf(`var gitTag = "%s"`, tag))
	b.WriteString("\n")
	b.WriteString("\n")

	b.WriteString("func GitTag() string {")
	b.WriteString("\n")
	b.WriteString("	return gitTag")
	b.WriteString("\n")
	b.WriteString("}")
	b.WriteString("\n")

	content := b.Bytes()

	wErr := ioutil.WriteFile(path0, content, 066)

	if wErr != nil {
		err = fmt.Errorf("write go file failed, %v", wErr)
		return
	}

	return
}
