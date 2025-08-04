package ignorePathRegexs

import (
	"regexp"
	"strings"
)

type IgnorePathRegexs []*regexp.Regexp

func New(reStr ...string) IgnorePathRegexs {
	output := make([]*regexp.Regexp, len(reStr))
	for _, dir := range reStr {
		output = append(output, regexp.MustCompile(dir))
	}
	return output
}

func (ipr *IgnorePathRegexs) MatchPath(path string) bool {
	for _, re := range *ipr {
		if re.MatchString(path) {
			return true
		}
	}
	return false
}

func (ipr *IgnorePathRegexs) String() string {
	var strs []string
	for _, re := range *ipr {
		strs = append(strs, re.String())
	}
	return strings.Join(strs, ",")
}

func (ipr *IgnorePathRegexs) Set(value string) error {
	*ipr = append(*ipr, regexp.MustCompile(value))
	return nil
}
