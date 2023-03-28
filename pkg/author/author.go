package author

import (
	"github.com/urfave/cli/v2"
)

var Author cli.Author

func init() {
	Author.Name = "oink"
	Author.Email = "wx11211@hotmail.com"
}
