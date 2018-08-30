package roll

import (
	"bytes"
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func (c *Command) cofd(ctx context.Context, args []string) string {
	dice, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Sprintf("%s is not a valid number of dice.", args[0])
	}

	rolls, rerolls, successes := cofdroll(dice, c.flags.cofd.again, c.flags.cofd.rote, c.flags.cofd.weakness, rand.Intn)
	return c.formatCofDResults(dice, rolls, rerolls, successes)
}

func (c *Command) formatCofDResults(dice int, rolls []int, rerolls []int, successes int) string {
	// parse the results
	var buff bytes.Buffer
	var flags []string

	// setup response
	buff.WriteString(fmt.Sprintf("rolled %d CofD dice", dice))

	// add roll modifying flags
	if c.flags.cofd.again != 10 {
		flags = append(flags, fmt.Sprintf("%d-again", c.flags.cofd.again))
	}
	if c.flags.cofd.rote {
		flags = append(flags, "rote")
	}
	if c.flags.cofd.weakness {
		flags = append(flags, "weakness")
	}
	if len(flags) > 0 {
		buff.WriteString(fmt.Sprintf(" (with %s)", strings.Join(flags, ", ")))
	}

	// add successes
	buff.WriteString(fmt.Sprintf(" for %d successes.", successes))

	if successes >= c.flags.cofd.exceptional {
		buff.WriteString(" Exceptional success!")
	}

	// add rolls and rerolls if desired
	if c.flags.verbose {
		buff.WriteString(fmt.Sprintf(" rolls: %d", rolls))

		if len(rerolls) > 0 {
			buff.WriteString(fmt.Sprintf(" rerolls: %d", rerolls))
		}
	}

	// send it
	return buff.String()
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func cofdroll(dice, again int, rote bool, weak bool, roll func(int) int) (rolls []int, rerolls []int, successes int) {
	rrn := 0

	// roll some dice
	for i := 0; i < dice; i++ {
		die := roll(10) + 1
		rolls = append(rolls, die)

		if die < 8 && rote {
			die = roll(10) + 1
			rerolls = append(rerolls, die)
		}

		if die >= 8 {
			successes++
		}

		if die >= again {
			rrn++
		}

		if weak && die == 1 {
			successes--
		}
	}

	// handle rerolls
	if rrn > 0 {
		r, rr, s := cofdroll(rrn, again, rote, weak, roll)
		rerolls = append(rerolls, r...)
		rerolls = append(rerolls, rr...)
		successes = successes + s
	}

	return rolls, rerolls, max(successes, 0)
}
