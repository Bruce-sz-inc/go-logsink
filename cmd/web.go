// Copyright © 2017 Sascha Andres <sascha.andres@outlook.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"log"

	"github.com/nightlyone/lockfile"
	"github.com/sascha-andres/go-logsink/web"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// webCmd represents the web command
var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Start a server instance with a web interface",
	Long: `Use web to start a web server. Navigate with your favorite
browser to localhost:8080 ( change the binding definition )
to see the logs in your browser.

  go-logsink web --serve ":80" --bind ":50051"`,
	Run: func(cmd *cobra.Command, args []string) {
		if "" != viper.GetString("lockfile") {
			lock, err := lockfile.New(viper.GetString("lockfile"))
			if err != nil {
				log.Fatal(err) // handle properly please!
			}
			err = lock.TryLock()

			// Error handling is essential, as we only try to get the lock.
			if err != nil {
				log.Fatal(fmt.Errorf("Cannot lock %q, reason: %v", lock, err))
			}

			defer lock.Unlock()
		}
		web.Start()
	},
}

func init() {
	RootCmd.AddCommand(webCmd)
	webCmd.Flags().StringP("bind", "b", ":50051", "Provide bind definition")
	webCmd.Flags().StringP("serve", "s", ":8080", "Provide bind definition")
	webCmd.Flags().Int32P("limit", "l", 0, "Limit number of visible lines")
	webCmd.Flags().Uint32P("break", "", 0, "Set this to the column # to break at")
	viper.BindPFlag("web.bind", webCmd.Flags().Lookup("bind"))
	viper.BindPFlag("web.serve", webCmd.Flags().Lookup("serve"))
	viper.BindPFlag("web.limit", webCmd.Flags().Lookup("limit"))
	viper.BindPFlag("web.break", webCmd.Flags().Lookup("break"))
}
