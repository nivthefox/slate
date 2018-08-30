// Copyright (c) 2018 Kevin Kragenbrink, II
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package roll

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// FormatCofDResults
type FormatCofDResultsSuite struct {
	suite.Suite
	cmd *Command
}

func (s *FormatCofDResultsSuite) SetupTest() {
	cmd := new(Command)
	cmd.flags.system = "cofd"
	cmd.flags.verbose = false
	cmd.flags.cofd.exceptional = 5
	cmd.flags.cofd.again = 10
	cmd.flags.cofd.rote = false
	cmd.flags.cofd.weakness = false
	s.cmd = cmd
}

func (s *FormatCofDResultsSuite) TestNormal() {
	dice := 2
	rolls := []int{10, 7}
	rerolls := []int{8}
	successes := 2

	assert.Equal(s.T(), "rolled 2 CofD dice for 2 successes.", s.cmd.formatCofDResults(dice, rolls, rerolls, successes))
}

func (s *FormatCofDResultsSuite) TestVerbose() {
	dice := 2
	rolls := []int{10, 7}
	rerolls := []int{8}
	successes := 2

	s.cmd.flags.verbose = true
	assert.Equal(s.T(), "rolled 2 CofD dice for 2 successes. rolls: [10 7] rerolls: [8]", s.cmd.formatCofDResults(dice, rolls, rerolls, successes))
}

func (s *FormatCofDResultsSuite) TestExceptional() {
	dice := 2
	rolls := []int{10, 7}
	rerolls := []int{8}
	successes := 2

	s.cmd.flags.cofd.exceptional = 2
	assert.Equal(s.T(), "rolled 2 CofD dice for 2 successes. Exceptional success!", s.cmd.formatCofDResults(dice, rolls, rerolls, successes))
}

func (s *FormatCofDResultsSuite) TestAgain() {
	dice := 2
	rolls := []int{10, 7}
	rerolls := []int{8}
	successes := 2

	s.cmd.flags.cofd.again = 9
	assert.Equal(s.T(), "rolled 2 CofD dice (with 9-again) for 2 successes.", s.cmd.formatCofDResults(dice, rolls, rerolls, successes))
}

func (s *FormatCofDResultsSuite) TestRote() {
	dice := 2
	rolls := []int{10, 7}
	rerolls := []int{8}
	successes := 2

	s.cmd.flags.cofd.rote = true
	assert.Equal(s.T(), "rolled 2 CofD dice (with rote) for 2 successes.", s.cmd.formatCofDResults(dice, rolls, rerolls, successes))
}

func (s *FormatCofDResultsSuite) TestWeak() {
	dice := 2
	rolls := []int{10, 7}
	rerolls := []int{8}
	successes := 2

	s.cmd.flags.cofd.weakness = true
	assert.Equal(s.T(), "rolled 2 CofD dice (with weakness) for 2 successes.", s.cmd.formatCofDResults(dice, rolls, rerolls, successes))
}

func (s *FormatCofDResultsSuite) TestChance() {
	dice := 0
	rolls := []int{5}
	rerolls := []int(nil)
	successes := 0

	assert.Equal(s.T(), "rolled 0 CofD dice (Chance Die) for 0 successes.", s.cmd.formatCofDResults(dice, rolls, rerolls, successes))
}

func (s *FormatCofDResultsSuite) TestCriticalFailure() {
	dice := 0
	rolls := []int{1}
	rerolls := []int(nil)
	successes := 0

	assert.Equal(s.T(), "rolled 0 CofD dice (Chance Die) for 0 successes. Critical failure!", s.cmd.formatCofDResults(dice, rolls, rerolls, successes))
}

func TestFormatCofDResult(t *testing.T) {
	suite.Run(t, new(FormatCofDResultsSuite))
}

// CofDRuoll
func TestCofDRoll(t *testing.T) {
	rolls := 0
	roll := func(n int) int {
		rolls++
		if rolls < 5 {
			return 5
		}
		return 7
	}

	r, rr, s := cofdroll(5, 10, false, false, roll)
	assert.Equal(t, []int{6, 6, 6, 6, 8}, r)
	assert.Equal(t, []int(nil), rr)
	assert.Equal(t, 1, s)
}

func TestCofDRollAgain(t *testing.T) {
	rolls := 0
	roll := func(n int) int {
		rolls++
		if rolls < 5 {
			return 5
		}
		if rolls < 6 {
			return 9
		}
		return 8
	}

	r, rr, s := cofdroll(5, 10, false, false, roll)
	assert.Equal(t, []int{6, 6, 6, 6, 10}, r)
	assert.Equal(t, []int{9}, rr)
	assert.Equal(t, 2, s)
}

func TestCofDRollRote(t *testing.T) {
	rolls := 0
	roll := func(n int) int {
		rolls++
		if rolls%2 == 1 {
			return 5
		}
		return 7
	}

	r, rr, s := cofdroll(5, 10, true, false, roll)
	assert.Equal(t, []int{6, 6, 6, 6, 6}, r)
	assert.Equal(t, []int{8, 8, 8, 8, 8}, rr)
	assert.Equal(t, 5, s)
}

func TestCofDRollWeak(t *testing.T) {
	rolls := 0
	roll := func(n int) int {
		rolls++
		if rolls < 5 {
			return 8
		}
		return 0
	}

	r, rr, s := cofdroll(5, 10, false, true, roll)
	assert.Equal(t, []int{9, 9, 9, 9, 1}, r)
	assert.Equal(t, []int(nil), rr)
	assert.Equal(t, 3, s)
}

func TestCofDRollChance(t *testing.T) {
	roll := func(n int) int {
		return 8
	}
	r, rr, s := cofdroll(0, 10, false, false, roll)
	assert.Equal(t, []int{9}, r)
	assert.Equal(t, []int(nil), rr)
	assert.Equal(t, 0, s)

	roll = func(n int) int {
		return 9
	}
	r, rr, s = cofdroll(0, 10, false, false, roll)
	assert.Equal(t, []int{10}, r)
	assert.Equal(t, []int(nil), rr)
	assert.Equal(t, 1, s)
}
