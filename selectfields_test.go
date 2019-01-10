package dbhelper

import (
	"testing"
)

type T struct {
	ID   int    `db:"id"`
	Name string `db:"name" json:"name"`
}

func TestSelectFields(t *testing.T) {

	var structA = new(T)

	tests := []struct {
		Input  interface{}
		Prefix string
		Want   string
	}{
		{structA, "", "id,name"},
		{structA, "a", "a.id,a.name"},
	}

	for _, test := range tests {
		if got := SelectFields(test.Input, test.Prefix); got != test.Want {
			t.Errorf("want %q, got %q", test.Want, got)
		}
	}
}
