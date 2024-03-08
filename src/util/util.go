package util

import "unicode"

const DEFAULT_FS_PERM = 0700
const DEFAULT_FILE_PERM = 0666

func ToTitleCase(name string) string {
	r := []rune(name)
	r[0] = unicode.ToTitle(r[0])
	return string(r)
}
