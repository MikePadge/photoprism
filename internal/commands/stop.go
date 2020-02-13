package commands

import (
	"syscall"

	"github.com/mikepadge/photoprism/internal/config"
	"github.com/sevlyar/go-daemon"
	"github.com/urfave/cli"
)

// StopCommand stops the daemon if running.
var StopCommand = cli.Command{
	Name:    "stop",
	Aliases: []string{"down"},
	Usage:   "Stops web server (only in daemon mode)",
	Action:  stopAction,
}

func stopAction(ctx *cli.Context) error {
	conf := config.NewConfig(ctx)

	log.Infof("looking for pid in \"%s\"", conf.PIDFilename())

	dcxt := new(daemon.Context)
	dcxt.PidFileName = conf.PIDFilename()
	child, err := dcxt.Search()

	if err != nil {
		log.Fatal(err)
	}

	err = child.Signal(syscall.SIGTERM)

	if err != nil {
		log.Fatal(err)
	}

	st, err := child.Wait()

	if err != nil {
		log.Info("daemon exited successfully")
		return nil
	}

	log.Infof("daemon[%v] exited[%v]? successfully[%v]?\n", st.Pid(), st.Exited(), st.Success())

	return nil
}
