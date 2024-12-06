package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// GreetCmd represents the greet command
var GreetCmd = &cobra.Command{
	Use:   "greet",
	Short: "Say hello to the user",
	Long:  `This command will greet the user with a message.`,
	Run: func(cmd *cobra.Command, args []string) {
		subject := "World"
		if len(args) > 0 {
			subject = args[0]
		}
		fmt.Printf("Hello, %s\n", subject)
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// GreetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// GreetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
