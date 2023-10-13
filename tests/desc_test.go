package tests

import (
	"testing"

	"github.com/cphovo/note/utils"
)

func TestParseDescription(t *testing.T) {
	desc := `DESCRIBE: 用于测试的描述
TAG: example, test
DATE: 2023/10/13`
	d := utils.ParseDescription(desc)
	expected := utils.Description{Describe: "用于测试的描述", Tag: "example, test", Date: "2023/10/13"}
	if d != expected {
		t.Errorf("Expected:\n%s\n\nGot:\n%s\n", expected, d)
	}
}

func TestDescriptionToString(t *testing.T) {
	d := utils.Description{Describe: "用于测试的描述", Tag: "example, test", Date: "2023/10/13"}
	actual := d.String()
	expected := "DESCRIBE: 用于测试的描述\nTAG: example, test\nDATE: 2023/10/13"
	if actual != expected {
		t.Errorf("Expected:\n%s\n\nGot:\n%s\n", expected, actual)
	}

	d = utils.Description{Describe: "用于测试的描述", Date: "2023/10/13"}
	actual = d.String()
	expected = "DESCRIBE: 用于测试的描述\nDATE: 2023/10/13"
	if actual != expected {
		t.Errorf("Expected:\n%s\n\nGot:\n%s\n", expected, actual)
	}

	d = utils.Description{}
	if d.String() != "" {
		t.Errorf("Expected:\n%s\n\nGot:\n%s\n", "", d)
	}
}
