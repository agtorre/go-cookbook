package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

// CountTheWays is a custom type that
// we'll read a flag into
type CountTheWays []int

func (c *CountTheWays) String() string {
	result := ""
	for _, v := range *c {
		if len(result) > 0 {
			result += " ... "
		}
		result += fmt.Sprint(v)
	}
	return result
}

// Set will be used by the flag package
func (c *CountTheWays) Set(value string) error {
	values := strings.Split(value, ",")

	for _, v := range values {
		i, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		*c = append(*c, i)
	}

	return nil
}

// Config will be the holder for our flags
type Config struct {
	subject      string
	isAwesome    bool
	howAwesome   int
	countTheWays CountTheWays
}

// Setup initializes a config from flags that
// are passed in
func (c *Config) Setup() {
	// you can set a flag directly like so:
	// var someVar = flag.String("flag_name", "default_val", "description")
	// but in practice putting it in a struct is generally better

	// longhand
	flag.StringVar(&c.subject, "subject", "", "subject is a string, it defaults to empty")
	// shorthand
	flag.StringVar(&c.subject, "s", "", "subject is a string, it defaults to empty (shorthand)")

	flag.BoolVar(&c.isAwesome, "isawesome", true, "is go awesome or what?")
	flag.IntVar(&c.howAwesome, "howawesome", 10, "how awesome out of 10?")

	// custom variable type
	flag.Var(&c.countTheWays, "c", "comma separated list of integers")
}

// GetMessage uses all of the internal
// config vars and returns a sentence
func (c *Config) GetMessage() string {
	msg := c.subject
	if c.isAwesome {
		msg += " is awesome"
	} else {
		msg += " is NOT awesome"
	}

	msg = fmt.Sprintf("%s with a certainty of %d out of 10. Let me count the ways %s", msg, c.howAwesome, c.countTheWays.String())
	return msg
}

func main() {
	// initialize our setup
	c := Config{}
	c.Setup()

	// generally call this from main
	flag.Parse()

	fmt.Println(c.GetMessage())

}
