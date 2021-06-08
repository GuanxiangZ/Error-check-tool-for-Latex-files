package command

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/uoa-compsci-399/error-correction-tool-for-latex/config"
	"github.com/uoa-compsci-399/error-correction-tool-for-latex/server"
	"os"
	"path/filepath"
)

var (
	ccmd *cobra.Command
)


func serverCmd() *cobra.Command {

	cfg := config.GetServeCmdConfig()

	ccmd = &cobra.Command{
		Use:     "serve",
		Short:   "Start or Stop the error correction of LaTeX files Web server. The server will run in background.",
		Example: "\n\ttexcheck serve \n\ttexcheck serve -p 11399\n\ttexcheck serve -p 10399 --parser-server-url=localhost:50051",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			var (
				logfile *os.File
			)

			// create directory to store log file
			_, err = os.Stat(cfg.LogDir)
			if os.IsNotExist(err) {
				err = os.MkdirAll(cfg.LogDir, 0755)
				if err != nil {
					logrus.Errorf("Service fails to start")
					logrus.Errorf(err.Error())
					if err == os.ErrPermission {
						logrus.Errorf("please try sudo")
					}
					return
				}
			}

			// create a log file if not exists
			logfile, err = os.OpenFile(filepath.Join(cfg.LogDir, "texcheck.log"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				logrus.Errorf("Service fails to start")
				logrus.Errorf(err.Error())
				if err == os.ErrPermission {
					logrus.Errorf("please try sudo")
				}

				return
			}
			defer logfile.Close()

			// print server information
			logrus.Infof("Error-Correction of LaTeX files Web server starts successfully...\n")
			logrus.Infof("Listening on port %d...\n", cfg.Port)
			logrus.Infof("POST\thttp://127.0.0.1:%d/api/upload", cfg.Port)

			logrus.Infof("Remote Parser Server URL: \t%s", cfg.ParserUrl)

			err = server.Serve(cfg.Port, logfile)
			if err != nil {
				return
			}

			return nil
		},
	}

	ccmd.Flags().Uint32VarP(&cfg.Port, "port", "p", 10399, "Specify the Web server running port")

	ccmd.Flags().StringVar(&cfg.ParserUrl, "parser-server-addr",  "localhost:50051", "Specify the LaTeX parser server address")
	ccmd.Flags().StringVar(&cfg.LogDir, "log-directory",  ".log", "Specify the log file directory")

	return ccmd
}
