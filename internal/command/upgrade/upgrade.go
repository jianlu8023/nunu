package upgrade

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/AlecAivazis/survey/v2"
	"github.com/jianlu8023/nunu/config"
	"github.com/spf13/cobra"
)

var CmdUpgrade = &cobra.Command{
	Use:     "upgrade",
	Short:   "Upgrade the nunu command.",
	Long:    "Upgrade the nunu command.",
	Example: "nunu upgrade",
	Run: func(_ *cobra.Command, _ []string) {
		var upgrade = false

		prompt := &survey.Confirm{
			Message: fmt.Sprintf("do you want to upgrade it to offical ?"),
		}
		err := survey.AskOne(prompt, &upgrade)
		if err != nil {
			log.Fatalf("internal error %s\n", err)
		}

		if upgrade {
			fmt.Printf("go install %s\n", config.NunuCmd)
			cmd := exec.Command("go", "install", config.NunuCmd)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				log.Fatalf("go install %s error\n", err)
			}
			fmt.Printf("\n🎉 Nunu upgrade successfully!\n\n")
		} else {
			fmt.Printf("\n Nothing to do \n\n")
		}
	},
}
