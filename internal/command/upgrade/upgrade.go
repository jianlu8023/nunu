package upgrade

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CmdUpgrade = &cobra.Command{
	Use:     "upgrade",
	Short:   "Upgrade the nunu command.",
	Long:    "Upgrade the nunu command.",
	Example: "nunu upgrade",
	Run: func(_ *cobra.Command, _ []string) {
		// fmt.Printf("go install %s\n", config.NunuCmd)
		// cmd := exec.Command("go", "install", config.NunuCmd)
		// cmd.Stdout = os.Stdout
		// cmd.Stderr = os.Stderr
		// if err := cmd.Run(); err != nil {
		// 	log.Fatalf("go install %s error\n", err)
		// }
		// fmt.Printf("\n🎉 Nunu upgrade successfully!\n\n")
		fmt.Printf("\n Nunu ban upgrade \n\n")
	},
}
