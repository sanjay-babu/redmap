package runner

import (
	"github.com/spf13/cobra"

	_ "github.com/praetorian-inc/redmap/pkg/plugins/all"
)

// ANSI escape codes
const (
	red   = "\x1b[31m"
	bold  = "\x1b[1m"
	reset = "\x1b[0m"
)

var banner = `
                           .___                     
_______   ____   __| _/_____ _____  ______  
\_  __ \_/ __ \ / __ |/     \\__  \ \____ \ 
 |  | \/\  ___// /_/ |  Y Y  \/ __ \|  |_> >
 |__|    \___  >____ |__|_|  (____  /   __/ 
             \/     \/     \/     \/|__|    
                          .___      _____   
           ____ ___.__. __| _/_____/ ____\  
  ______ _/ ___<   |  |/ __ |/ __ \   __\   
 /_____/ \  \___\___  / /_/ \  ___/|  |     
          \___  > ____\____ |\___  >__|     
              \/\/         \/    \/         
`

var rootCmd = &cobra.Command{
	Use:   "redmap",
	Short: "Organizational asset discovery tool",
	Long:  banner + "\nRedMap discovers CIDR blocks and domains associated with an organization using multiple OSINT data sources.",
}

// SetVersion sets the version string displayed by --version.
func SetVersion(v string) {
	rootCmd.Version = v
}

// Execute runs the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(newRunCmd())
	rootCmd.AddCommand(newListCmd())
}
