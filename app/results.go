package app

import (
	"errors"

	"github.com/confio/weave"
)

// ResultsFromKeys returns a ResultSet of all keys
// given a set of models
func ResultsFromKeys(models []weave.Model) *ResultSet {
	res := make([][]byte, len(models))
	for i, m := range models {
		res[i] = m.Key
	}
	return &ResultSet{res}
}

// ResultsFromValues returns a ResultSet of all values
// given a set of models
func ResultsFromValues(models []weave.Model) *ResultSet {
	res := make([][]byte, len(models))
	for i, m := range models {
		res[i] = m.Value
	}
	return &ResultSet{res}
}

// JoinResults inverts ResultsFromKeys and ResultsFromValues
// and makes then a consistent whole again
func JoinResults(keys, values *ResultSet) ([]weave.Model, error) {
	kref, vref := keys.Refs, values.Refs
	if len(kref) != len(vref) {
		return nil, errors.New("Mismatches result set size")
	}
	mods := make([]weave.Model, len(kref))
	for i := range mods {
		mods[i] = weave.Model{
			Key:   kref[i],
			Value: vref[i],
		}
	}
	return mods, nil
}
