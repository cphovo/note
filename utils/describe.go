package utils

import (
	"fmt"
	"strings"
)

type Description struct {
	Describe string
	Tag      string
	Date     string
}

func (d *Description) String() string {
	var parts []string

	if d.Describe != "" {
		parts = append(parts, fmt.Sprintf("DESCRIBE: %s", d.Describe))
	}
	if d.Tag != "" {
		parts = append(parts, fmt.Sprintf("TAG: %s", d.Tag))
	}
	if d.Date != "" {
		parts = append(parts, fmt.Sprintf("DATE: %s", d.Date))
	}

	return strings.Join(parts, "\n")
}

func (d *Description) Code() string {
	return fmt.Sprintf("```\n%s\n```", d)
}

func ParseDescription(desc string) Description {
	describe, tag, date := "", "", ""

	lines := strings.Split(desc, "\n")

	for _, line := range lines {
		if strings.HasPrefix(line, "DESCRIBE") {
			words := strings.Split(line, ":")
			if len(words) >= 2 {
				describe = strings.TrimSpace(strings.Join(words[1:], ":"))
			}
		}
		if strings.HasPrefix(line, "TAG") {
			words := strings.Split(line, ":")
			if len(words) >= 2 {
				tag = strings.TrimSpace(strings.Join(words[1:], ":"))
			}
		}
		if strings.HasPrefix(line, "DATE") {
			words := strings.Split(line, ":")
			if len(words) == 2 {
				date = strings.TrimSpace(words[1])
			}
		}
	}
	return Description{Describe: describe, Tag: tag, Date: date}
}

func CheckIfContainsDescription(content string) (Description, bool) {
	firstCodeBlock := FindFirstCodeBlock(content)
	desc := ParseDescription(firstCodeBlock)
	if desc.Describe != "" || desc.Tag != "" || desc.Date != "" {
		return desc, true
	}
	return Description{}, false
}
