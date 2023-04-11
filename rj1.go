package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano() + int64(os.Getpid()))

	outputmap(os.Stdout)
}

func outputarray(out io.Writer) {
	fmt.Fprintf(out, "[")
	leader := ' '
	v := rand.Intn(4)
	for elements := rand.Intn(10) + 1; elements >= 0; elements-- {
		fmt.Fprintf(out, "%c%v", leader, value(v))
		leader = ','
	}
	fmt.Fprintf(out, "]")
}

func value(typ int) interface{} {
	var r interface{}
	switch typ {
	case 0: // bool
		switch rand.Intn(2) {
		case 0:
			r = true
		case 1:
			r = false
		}
	case 1: // string
		r = fmt.Sprintf("%q", randomstring())
	case 2: // float
		r = rand.NormFloat64()
	default:
		n := rand.Int()
		if rand.Intn(2) == 1 {
			n = -n
		}
		r = n
	}
	return r
}

func outputmap(out io.Writer) {
	fmt.Fprintf(out, "{")

	leader := ' '

	for elements := rand.Intn(10) + 1; elements >= 0; elements-- {
		fmt.Fprintf(out, "%c%q: ", leader, randomstring())
		switch rand.Intn(7) {
		case 0: // integer
			fmt.Fprintf(out, "%d", rand.Int())
		case 1: // string
			fmt.Fprintf(out, "%q", randomstring())
		case 2: // true
			fmt.Fprintf(out, "true")
		case 3: // false
			fmt.Fprintf(out, "false")
		case 4: // floating point number
			fmt.Fprintf(out, "%f", rand.NormFloat64())
		case 5: // array
			outputarray(out)
		case 6: // map
			outputmap(out)
		}

		leader = ','
	}
	fmt.Fprintf(out, "}\n")
}

func randomstring() string {
	ary := []rune{
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
		'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
		'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
		' ',
	}

	max := rand.Intn(17) + 3
	rstring := make([]rune, 0, 20)
	l := len(ary)

	for i := 0; i < max; i++ {
		letter := ary[rand.Intn(l)]
		rstring = append(rstring, letter)
	}
	return string(rstring)
}
