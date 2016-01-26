package sliceop

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var prefillTests = []struct {
	input    int
	symbol   string
	expected []string
}{
	{9, "?", []string{"?", "?", "?", "?", "?", "?", "?", "?", "?"}},
	{4, "#", []string{"#", "#", "#", "#"}},
	{3, "  ", []string{"  ", "  ", "  "}},
	{1, "0", []string{"0"}},
	{0, "###", []string{}},
}

func TestPrefill(t *testing.T) {
	for _, test := range prefillTests {
		got := Prefill(test.input, test.symbol)
		assert.Equal(t, test.expected, got)
	}
}

var mapTests = []struct {
	input    []string
	expected []string
}{
	{[]string{"A", "b", "c"}, []string{"a", "b", "c"}},
	{[]string{"vll", "A"}, []string{"vll", "a"}},
	{[]string{}, []string{}},
}

func TestMap(t *testing.T) {
	mapFunc := func(s string) string {
		return strings.ToLower(s)
	}

	for _, test := range mapTests {
		got := Map(mapFunc, test.input...)

		assert.Equal(t, len(test.expected), len(test.input))
		assert.Equal(t, test.expected, got)
	}
}

var includesTests = []struct {
	input    []string
	key      string
	expected bool
}{
	{[]string{"a", "b", "c"}, "b", true},
	{[]string{"a", "b", "c"}, "d", false},
	{[]string{}, "d", false},
	{[]string{"D"}, "d", false},
}

func TestIncludes(t *testing.T) {
	for _, test := range includesTests {
		got := Includes(test.input, test.key)

		assert.Equal(t, test.expected, got,
			fmt.Sprintf("Expected %t for %v", test.expected, test.input))
	}
}

var notIncludesTests = []struct {
	input    []string
	key      string
	expected bool
}{
	{[]string{"a", "b", "c"}, "b", false},
	{[]string{"a", "b", "c"}, "d", true},
	{[]string{}, "d", true},
	{[]string{"D"}, "d", true},
}

func TestNotIncludes(t *testing.T) {
	for _, test := range notIncludesTests {
		got := NotIncludes(test.input, test.key)

		assert.Equal(t, test.expected, got,
			fmt.Sprintf("Expected %t for %v", test.expected, test.input))
	}
}

var rejectTests = []struct {
	input    []string
	toReject []string
	expected []string
}{
	{
		[]string{"A", "b", "c"},
		[]string{"b", "c"},
		[]string{"A"},
	},
	{
		[]string{"A", "b", "c"},
		[]string{},
		[]string{"A", "b", "c"},
	},
	{
		nil,
		[]string{"b", "c"},
		nil,
	},
}

func TestReject(t *testing.T) {
	for _, test := range rejectTests {
		got := Reject(test.input, test.toReject...)

		assert.Equal(t, test.expected, got)
	}
}

var selectTests = []struct {
	input    []string
	toSelect []string
	expected []string
}{
	{
		[]string{"A", "b", "c"},
		[]string{"b", "c"},
		[]string{"b", "c"},
	},
	{
		[]string{"A", "b", "c"},
		[]string{"b", ""},
		[]string{"b"},
	},
	{
		[]string{"A", "b", "c"},
		nil,
		[]string{"A", "b", "c"},
	},
	{
		nil,
		[]string{"b", "c"},
		nil,
	},
}

func TestSelect(t *testing.T) {
	for _, test := range selectTests {
		got := Select(test.input, test.toSelect...)

		assert.Equal(t, test.expected, got)
	}
}

var uniqueTests = []struct {
	input    []string
	expected []string
}{
	{[]string{"A", "b", "b", "b", "c"}, []string{"A", "b", "c"}},
	{[]string{"A", "1"}, []string{"A", "1"}},
	{[]string{}, []string{}},
}

func TestUnique(t *testing.T) {
	for _, test := range uniqueTests {
		got := Unique(test.input...)

		assert.Equal(t, test.expected, got)
	}
}

var intersectionTests = []struct {
	u1       []string
	u2       []string
	expected []string
}{
	{
		[]string{"a", "b1", "1"},
		[]string{"b1", "1"},
		[]string{"b1", "1"},
	},
	{
		[]string{"b2", "1"},
		[]string{"a", "b2", "1"},
		[]string{"b2", "1"},
	},
	{
		[]string{},
		[]string{"a3", "1"},
		nil,
	},
	{
		[]string{"1"},
		[]string{"2"},
		nil,
	},
	{
		[]string{"a4", "1"},
		[]string{},
		nil,
	},
	{
		[]string{},
		[]string{},
		nil,
	},
	{
		nil,
		nil,
		nil,
	},
}

func TestIntersection(t *testing.T) {
	for _, test := range intersectionTests {
		got := Intersection(test.u1, test.u2)
		assert.Equal(t, test.expected, got)
	}
}

var countFuncTests = []struct {
	input    []string
	f        func(string) bool
	expected int
}{
	{
		[]string{"some", "key", "gone", "here"},
		func(key string) bool {
			return key == "some"
		},
		1,
	},
	{
		[]string{"some", "some", "some", "some"},
		func(key string) bool {
			return key == "some"
		},
		4,
	},
	{
		[]string{"some", "some", "some", "some"},
		func(key string) bool {
			return false
		},
		0,
	},
	{
		nil,
		nil,
		0,
	},
}

func TestCountFunc(t *testing.T) {
	for _, test := range countFuncTests {
		got := CountFunc(test.f, test.input...)

		assert.Equal(t, test.expected, got)
	}
}
