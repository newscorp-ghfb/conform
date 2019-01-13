// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	"os"

	"github.com/autonomy/conform/internal/enforcer"
	"github.com/autonomy/conform/internal/policy"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// enforceCmd represents the enforce command
var enforceCmd = &cobra.Command{
	Use:   "enforce",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			err := errors.Errorf("The enforce command does not take arguments")

			fmt.Println(err)
			os.Exit(1)
		}
		e, err := enforcer.New()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		opts := []policy.Option{}

		if commitMsgFile := cmd.Flags().Lookup("commit-msg-file").Value.String(); commitMsgFile != "" {
			opts = append(opts, policy.WithCommitMsgFile(&commitMsgFile))
		}

		e.Enforce(opts...)
	},
}

func init() {
	enforceCmd.Flags().String("commit-msg-file", "", "the path to the temporary commit message file")
	RootCmd.AddCommand(enforceCmd)
}
