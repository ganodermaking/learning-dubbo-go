package http

import (
	"fmt"
	"learning-go/cmd/http/routers"
	"learning-go/initializers"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// HTTPCmd represents the http command
var Cmd = &cobra.Command{
	Use:   "http",
	Short: "http server",
	Long:  `run your http server`,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("http called")
		initializers.Init()

		r := routers.Init()
		addr := fmt.Sprintf(":%d", 5000)
		if err := r.Run(addr); err != nil {
			logrus.Panicf("start http server failed: %s", err)
		}
	},
}
