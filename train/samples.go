package main

import (
	"crypto/md5"
	"encoding/json"
	"io/ioutil"

	"github.com/unixpickle/poeturn/model"
	"github.com/unixpickle/sgd"
	"github.com/unixpickle/weakai/rnn/seqtoseq"
)

// A SampleSet is an sgd.SampleSet that generates
// vectovec.Sample samples for poems.
type SampleSet struct {
	Poems []string
}

// LoadSampleSet reads the samples from a JSON file.
func LoadSampleSet(f string, maxLen int) (*SampleSet, error) {
	data, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}
	var raw []string
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}
	var res SampleSet
	for _, x := range raw {
		if len(x) <= maxLen {
			res.Poems = append(res.Poems, x)
		}
	}
	return &res, nil
}

// Len returns the number of poems.
func (s *SampleSet) Len() int {
	return len(s.Poems)
}

// Swap swaps two poems.
func (s *SampleSet) Swap(i, j int) {
	s.Poems[i], s.Poems[j] = s.Poems[j], s.Poems[i]
}

// GetSample returns the sample at the given index as a
// vectovec.Sample object.
func (s *SampleSet) GetSample(i int) interface{} {
	p := s.Poems[i]
	var res seqtoseq.Sample
	if len(p) == 0 {
		return res
	}
	res.Inputs = append(res.Inputs, model.OneHot('\n'))
	for _, b := range []byte(p) {
		res.Outputs = append(res.Outputs, model.OneHot(b))
	}
	res.Inputs = append(res.Inputs, res.Outputs[:len(res.Outputs)-1]...)
	return res
}

// Subset generates a slice of the SampleSet.
func (s *SampleSet) Subset(i, j int) sgd.SampleSet {
	return &SampleSet{Poems: s.Poems[i:j]}
}

// Copy generates a copy of the SampleSet.
func (s *SampleSet) Copy() sgd.SampleSet {
	return &SampleSet{Poems: append([]string{}, s.Poems...)}
}

// Hash generates a hash for the poem.
func (s *SampleSet) Hash(i int) []byte {
	data := md5.Sum([]byte(s.Poems[i]))
	return data[:]
}
