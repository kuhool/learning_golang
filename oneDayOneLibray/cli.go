package oneDayOneLibray

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func TestCli() {
	app := cli.NewApp()
	app.Name = "My CLI App"
	app.Usage = "A sample CLI app built with urfave/cli"

	app.Commands = []*cli.Command{
		{
			Name:    "greet",
			Aliases: []string{"hello"},
			Usage:   "Greet someone",
			Action: func(c *cli.Context) error {
				//命令：go run main.go  hello 11 22 （11）
				//输出：Hello 11
				fmt.Println("Hello", c.Args().First())
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
