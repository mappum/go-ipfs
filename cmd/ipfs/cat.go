package main

import (
	"os"

	"github.com/jbenet/go-ipfs/Godeps/_workspace/src/github.com/gonuts/flag"
	"github.com/jbenet/go-ipfs/Godeps/_workspace/src/github.com/jbenet/commander"
	"github.com/jbenet/go-ipfs/daemon"
	u "github.com/jbenet/go-ipfs/util"
)

var cmdIpfsCat = &commander.Command{
	UsageLine: "cat",
	Short:     "Show ipfs object data.",
	Long: `ipfs cat <ipfs-path> - Show ipfs object data.

    Retrieves the object named by <ipfs-path> and displays the Data
    it contains.
`,
	Run:  catCmd,
	Flag: *flag.NewFlagSet("ipfs-cat", flag.ExitOnError),
}

func catCmd(c *commander.Command, inp []string) error {
	u.Debug = true
	if len(inp) < 1 {
		u.POut(c.Long)
		return nil
	}

	com := daemon.NewCommand()
	com.Command = "cat"
	com.Args = inp

	err := daemon.SendCommand(com, "localhost:12345")
	if err != nil {
		n, err := localNode(false)
		if err != nil {
			return err
		}

		daemon.ExecuteCommand(com, n, os.Stdout)
	}
	return nil
}
