package model

import (
	"math"
	"math/rand"

	"github.com/unixpickle/num-analysis/linalg"
	"github.com/unixpickle/weakai/rnn"
)

const maxLen = 1000

// A Session represents a writing session in which a user
// and a Model take turns writing pieces of a poem.
type Session struct {
	runner *rnn.Runner
}

// NewSession creates a fresh Session.
func NewSession(m *Model) *Session {
	return &Session{
		runner: &rnn.Runner{Block: m.Block},
	}
}

// Query requests the next line from the Session.
//
// The line will not include the final trailing newline,
// but it may include intermediate newlines.
func (s *Session) Query() string {
	last := oneHot('\n')
	var res string
	for {
		next := s.runner.StepTime(last)
		if len(res) == maxLen {
			break
		}
		ch := randomSelection(next)
		if ch == '\n' && res != "" && res != "\n" {
			break
		}
		last = oneHot(ch)
		res += string(ch)
	}
	return res
}

// Dictate tells the Session the next line in the poem.
//
// The line should not include a trailing newline.
func (s *Session) Dictate(line string) {
	s.runner.StepTime(oneHot('\n'))
	for _, b := range []byte(line) {
		s.runner.StepTime(oneHot(b))
	}
}

func oneHot(b byte) linalg.Vector {
	res := make(linalg.Vector, CharCount)
	res[int(b)] = 1
	return res
}

func randomSelection(logProbs linalg.Vector) byte {
	v := rand.Float64()
	for i, x := range logProbs {
		v -= math.Exp(x)
		if v < 0 {
			return byte(i)
		}
	}
	return 0xff
}
