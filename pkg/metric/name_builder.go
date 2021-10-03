package metric

import "strings"

// NameBuilder - build metrics name in prometheus format
// for instance -
//     * foo
//     * foo{bar="baz"}
//     * foo{bar="baz",aaa="b"}
// builder is not concurrent safe
type NameBuilder struct {
	name string
	m    tupleMap

	builder strings.Builder
}

const (
	separator = '='
	comma     = ','
	quote     = '"'
)

func (b *NameBuilder) Name(n string) *NameBuilder {
	b.name = n
	return b
}

func (b *NameBuilder) Add(k, v string) *NameBuilder {
	if contains(k, comma) || contains(v, comma) {
		return b
	}
	if contains(k, separator) || contains(v, separator) {
		return b
	}

	b.builder.Grow(len(v) + 2)
	b.builder.WriteByte(quote)
	b.builder.WriteString(v)
	b.builder.WriteByte(quote)

	b.m = append(b.m, tuple{Key: k, Value: b.builder.String()})

	b.builder.Reset()

	return b
}

func contains(s string, c byte) bool {
	for i := range s {
		if s[i] == c {
			return true
		}
	}

	return false
}

func (b *NameBuilder) String() string {
	b.builder.WriteString(b.name)
	b.builder.WriteByte('{')
	b.builder.WriteString(b.m.String())
	b.builder.WriteByte('}')

	defer func() {
		b.builder.Reset()
		b.m = nil
	}()

	return b.builder.String()
}

type tupleMap []tuple

func (t tupleMap) String() string {
	var builder strings.Builder

	for i := range t {
		builder.WriteString(t[i].Key)
		builder.WriteByte(separator)
		builder.WriteString(t[i].Value)
		if i != len(t)-1 {
			builder.WriteByte(comma)
		}
	}

	return builder.String()
}

type tuple struct {
	Key   string
	Value string
}

