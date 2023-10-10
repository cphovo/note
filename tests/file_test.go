package tests

import (
	"testing"

	"github.com/cphovo/note/utils"
)

func TestRemoveCodeBlocks(t *testing.T) {
	fi, err := utils.GetFileInfo("./input.md")
	if err != nil {
		t.Fail()
	}
	expected, err := utils.GetFileInfo("./output1.md")
	if err != nil {
		t.Fail()
	}
	content := utils.RemoveCodeBlocks(fi.Content)

	if content != expected.Content {
		t.Errorf("Expected:\n%s\n\nGot:\n%s\n", expected.Content, content)
	}
}

func TestRemoveCodeBlocksAndEmptyLines(t *testing.T) {
	fi, err := utils.GetFileInfo("./input.md")
	if err != nil {
		t.Fail()
	}
	expected, err := utils.GetFileInfo("./output2.md")
	if err != nil {
		t.Fail()
	}
	content := utils.RemoveCodeBlocks(fi.Content)
	contentRemovedEmptyLines := utils.RemoveEmptyLines(content)
	if contentRemovedEmptyLines != expected.Content {
		t.Errorf("Expected:\n%s\n\nGot:\n%s\n", expected.Content, contentRemovedEmptyLines)
	}
}
