package author

import (
	"github.com/urfave/cli/v2"
)

var Author cli.Author

func init() {
	Author.Name = "oink"
	Author.Email = "https://github.com/berryalen02/resonance"
}
