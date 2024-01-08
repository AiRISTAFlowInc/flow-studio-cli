package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindGoModPackageSrc(t *testing.T) {
	str, ver, err := FindGoModPackageSrc("github.com/AiRISTAFlowInc/flow-studio-core", "", true)
	if err != nil {
		fmt.Println("err:", err)
		t.FailNow()
	}

	fmt.Println("path: ", str)
	fmt.Println("ver: ", ver)
}


func TestFindOldPackageSrc(t *testing.T) {
	str, ver, err := FindOldPackageSrc("github.com/AiRISTAFlowInc/flow-studio-cli")
	if err != nil {
		fmt.Println("err:", err)
		t.FailNow()
	}

	fmt.Println("path: ", str)
	fmt.Println("ver: ", ver)
}


func TestFindGoModPackageSrcNotFound(t *testing.T) {
	_, _, err := FindGoModPackageSrc("github.com/project-blah/core", "", true)
	assert.True(t, IsPkgNotFoundError(err))
	fmt.Println("err: ", err)
}


func TestFindOldPackageSrcNotFound(t *testing.T) {
	_, _, err := FindOldPackageSrc("github.com/project-blah/core")
	assert.True(t, IsPkgNotFoundError(err))
	fmt.Println("err: ", err)
}
