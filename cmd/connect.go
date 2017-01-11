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

	"github.com/sascha-andres/go-logsink/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to a go-logsink server and forward stdin",
	Long: `This command is used to connect to a go-logsink server.
Call it to forward data piped ito this application to the server.`,
	Run: func(cmd *cobra.Command, args []string) {
		address := viper.GetString("address")
		fmt.Printf("Connecting to %s\n", address)
		client.Connect(address)
	},
}

func init() {
	RootCmd.AddCommand(connectCmd)
	connectCmd.Flags().StringP("address", "a", "localhost:50051", "Provide server address")
	viper.BindPFlag("address", listenCmd.Flags().Lookup("address"))
}
