package pkg

import (
	"encoding/json"
	"sort"
	"strings"
)

type StringSlice []string

func (ref StringSlice) Len() int           { return len(ref) }
func (ref StringSlice) Swap(i, j int)      { ref[i], ref[j] = ref[j], ref[i] }
func (ref StringSlice) Less(i, j int) bool { return ref[i] < ref[j] }

func (ref *StringSlice) Filter(closure func(row string) bool) *StringSlice {
	rows := NewStringSlice()
	for _, _row := range *ref {
		if closure(_row) {
			rows.Push(_row)
		}
	}
	return rows
}

func (ref *StringSlice) Unique() *StringSlice {
	rows := NewStringSlice()
	for i := 0; i < len(*ref); i++ {
		repeat := false
		for j := i + 1; j < len(*ref); j++ {
			if (*ref)[i] == (*ref)[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			rows.Push((*ref)[i])
		}
	}
	return rows
}

func (ref *StringSlice) First(closure func(row string) bool) string {
	for i, _row := range *ref {
		if closure(_row) {
			return (*ref)[i]
		}
	}
	return ""
}

func (ref *StringSlice) Exists(closure func(row string) bool) bool {
	return len(*ref.Filter(closure)) > 0
}

func (ref *StringSlice) Each(closure func(row *string)) *StringSlice {
	for i, _ := range *ref {
		closure(&(*ref)[i])
	}
	return ref
}

func (ref *StringSlice) Has(s string) bool {
	for _, v := range *ref {
		if v == s {
			return true
		}
	}
	return false
}

func (ref *StringSlice) Contains(s string) bool {
	for _, v := range *ref {
		if strings.Contains(v, s) {
			return true
		}
	}
	return false
}

func (ref *StringSlice) IsEmpty() bool {
	return ref.Length() > 0
}

func (ref *StringSlice) Push(f ...string) *StringSlice {
	for _, s := range f {
		*ref = append(*ref, s)
	}
	return ref
}

func (ref *StringSlice) Sort() *StringSlice {
	sort.Sort(ref)
	return ref
}

func (ref *StringSlice) JsonString() string {
	res, err := json.Marshal(*ref)
	if err != nil {
		panic(err)
	}
	return string(res)
}

func (ref *StringSlice) Length() int {
	return len(*ref)
}

func NewStringSlice(strings ...string) *StringSlice {
	rows := StringSlice{}
	for _, s := range strings {
		rows.Push(s)
	}
	return &rows
}
