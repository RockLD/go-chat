package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var AppCmdServe = &cobra.Command{
	Use:   "run",
	Short: "启动chat服务",
	Long:  `启动chat服务`,
	Run:   Run,
}

func Run(cmd *cobra.Command, arg []string) {
	fmt.Println("start service")
}
