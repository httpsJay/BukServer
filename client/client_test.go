package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestSet(t *testing.T) {
	testTable := []struct {
		x   string
		y   string
		err error
	}{
		{"jay", "software", nil},
		{"jack", "operation", nil},
		{"rakeshk", "software", nil},
		{"abc-1", "abc", nil},
		{"abc-2", "abc", nil},
		{"xyz-1", "xyz", nil},
		{"xyz-2", "xyz", nil},
	}

	for _, table := range testTable {
		resp := Set(table.x, table.y)
		if resp != table.err {
			t.Errorf("Key [%s] Value [%s] - Return[%v].", table.x, table.y, table.err)
		}
	}
}

func TestGet(t *testing.T) {
	testTable := []struct {
		x   string
		res string
		// err error
	}{
		{"jay", "software"},
		{"jack", "operation"},
		{"rakeshk", "software"},
		{"abc-1", "abc"},
		{"abc-2", "abc"},
		{"xyz-1", "xyz"},
		{"xyz-2", "xyz"},
	}

	for _, table := range testTable {
		resp, err := Get(table.x)
		if err != nil {
			t.Errorf("error. %s", err)
		}
		if resp != table.res {
			t.Errorf("Key [%s] - Return[%v].", table.x, table.res)
		}
	}
}

func TestSearchPrefix(t *testing.T) {
	testTable := []struct {
		x   string
		res []string
		// err error
	}{
		{"j", []string{"jay", "jack"}},
		{"ja", []string{"jay", "jack"}},
		{"abc", []string{"abc-1", "abc-2"}},
		{"xyz", []string{"xyz-1", "xyz-2"}},
	}

	for _, table := range testTable {
		resp, err := SearchPrefix(table.x)
		if err != nil {
			t.Errorf("error. %s", err)
		}

		// checking without order
		sort.Strings(resp)
		sort.Strings(table.res)

		// TODO: check each result in list
		if !reflect.DeepEqual(resp, table.res) {
			t.Errorf("Key [%s] - Return[%v]. - RESPONSE - [%v]", table.x, table.res, resp)
		}
	}
}

func TestSearchSuffix(t *testing.T) {
	testTable := []struct {
		x   string
		res []string
		// err error
	}{
		{"k", []string{"jack", "rakeshk"}},
		{"-1", []string{"xyz-1", "abc-1"}},
	}

	for _, table := range testTable {
		resp, err := SearchSuffix(table.x)
		if err != nil {
			t.Errorf("error. %s", err)
		}
		// checking without order
		sort.Strings(resp)
		sort.Strings(table.res)

		// TODO: check each result in list
		if !reflect.DeepEqual(resp, table.res) {
			t.Errorf("Key [%s] - Return[%v]. - RESPONSE - [%v]", table.x, table.res, resp)
		}
	}
}
