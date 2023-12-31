package utils

import (
	"go/build"
	"path/filepath"
	"reflect"
	"testing"
)

func TestRemoveCodeBlocks(t *testing.T) {
	fi, err := GetFileInfo("../testdata/md/input.md")
	if err != nil {
		t.Fail()
	}
	expected, err := GetFileInfo("../testdata/md/output1.md")
	if err != nil {
		t.Fail()
	}
	content := RemoveCodeBlocks(fi.Content)

	if content != expected.Content {
		t.Errorf("Expected:\n%s\n\nGot:\n%s\n", expected.Content, content)
	}
}

func TestRemoveCodeBlocksAndEmptyLines(t *testing.T) {
	fi, err := GetFileInfo("../testdata/md/input.md")
	if err != nil {
		t.Fail()
	}
	expected, err := GetFileInfo("../testdata/md/output2.md")
	if err != nil {
		t.Fail()
	}
	content := RemoveCodeBlocks(fi.Content)
	contentRemovedEmptyLines := RemoveEmptyLines(content)
	if contentRemovedEmptyLines != expected.Content {
		t.Errorf("Expected:\n%s\n\nGot:\n%s\n", expected.Content, contentRemovedEmptyLines)
	}
}

func TestParseDescription(t *testing.T) {
	desc := "DESCRIBE: 用于测试的描述\nTAG: example, test\nDATE: 2023/10/13"
	d := ParseDescription(desc)
	expected := Description{Describe: "用于测试的描述", Tag: "example, test", Date: "2023/10/13"}
	if d != expected {
		t.Errorf("Expected:\n%s\n\nGot:\n%s\n", expected, d)
	}
}

func TestDescriptionToString(t *testing.T) {
	d := Description{Describe: "用于测试的描述", Tag: "example, test", Date: "2023/10/13"}
	actual := d.String()
	expected := "DESCRIBE: 用于测试的描述\nTAG: example, test\nDATE: 2023/10/13"
	if actual != expected {
		t.Errorf("Expected:\n%s\n\nGot:\n%s\n", expected, actual)
	}

	d = Description{Describe: "用于测试的描述", Date: "2023/10/13"}
	actual = d.String()
	expected = "DESCRIBE: 用于测试的描述\nDATE: 2023/10/13"
	if actual != expected {
		t.Errorf("Expected:\n%s\n\nGot:\n%s\n", expected, actual)
	}

	d = Description{}
	if d.String() != "" {
		t.Errorf("Expected:\n%s\n\nGot:\n%s\n", "", d)
	}
}

func TestDescriptionToCode(t *testing.T) {
	desc := Description{Describe: "用于测试的描述", Tag: "example, test", Date: "2023/10/13"}
	actual := desc.Code()
	expected := "```\nDESCRIBE: 用于测试的描述\nTAG: example, test\nDATE: 2023/10/13\n```"
	if actual != expected {
		t.Errorf("Expected:\n%s\n\nGot:\n%s\n", expected, actual)
	}
}

func TestFindFirstCodeBlock(t *testing.T) {
	fi, err := GetFileInfo("../testdata/md/input2.md")
	if err != nil {
		t.Fail()
	}
	actual := FindFirstCodeBlock(fi.Content)
	expected := "```\nDESCRIBE: 用于测试的描述\nTAG: example, test\nDATE: 2023/10/13\n```"
	if actual != expected {
		t.Errorf("Expected:\n%s\n\nGot:\n%s\n", expected, actual)
	}

	d := ParseDescription(expected)
	e := Description{Describe: "用于测试的描述", Tag: "example, test", Date: "2023/10/13"}
	if d != e {
		t.Errorf("Expected:\n%s\n\nGot:\n%s\n", e, d)
	}
}

func TestGetWordCount(t *testing.T) {
	text := "测试示例的示例字符串：这是一个测试示例，示例内容用于测试中文分词功能是否正常，英文分词暂时不支持。"

	pkg, err := build.Import("github.com/yanyiwu/gojieba", "", build.FindOnly)
	if err != nil {
		t.Fail()
	}

	dictDir := filepath.Join(pkg.Dir, "dict")
	jb := JieBa{
		filepath.Join(dictDir, "jieba.dict.utf8"),
		filepath.Join(dictDir, "hmm_model.utf8"),
		filepath.Join(dictDir, "user.dict.utf8"),
		filepath.Join(dictDir, "idf.utf8"),
		filepath.Join(dictDir, "stop_words.utf8"),
	}
	actual := jb.GetWordCount(text, "n", "nr", "nz", "ns", "nt", "nw", "vn")[:3]
	expected := []WordCount{
		{Word: "示例", Count: 4},
		{Word: "测试", Count: 3},
		{Word: "分词", Count: 2},
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected:\n%v\n\nGot:\n%v\n", expected, actual)
	}
}
