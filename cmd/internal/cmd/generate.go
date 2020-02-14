// Copyright (c) 2019 IoTeX Foundation
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/iotexproject/cooladdress/cooladdress"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "gen [suffix] [timeout]",
	Short: `Generates address with suffix,timeout such as "300ms", "-1.5h" or "2h45m". Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h"`,
	Long:  `Generates address with suffix,timeout such as "300ms", "-1.5h" or "2h45m". Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h"`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now().Unix()
		fmt.Println(start)
		fmt.Println(generate(args))
		end := time.Now().Unix()
		fmt.Println(end, ":", end-start)
	},
}

func generate(args []string) (string, error) {
	return cooladdress.Gen(args)
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
