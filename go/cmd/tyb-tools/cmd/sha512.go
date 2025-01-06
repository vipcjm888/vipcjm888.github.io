package cmd

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"

	"github.com/spf13/cobra"
)

var sha512Cmd = &cobra.Command{
	Use:   "sha512",
	Short: "sha512 encode",
	Long:  `sha512 encode`,
	Run: func(cmd *cobra.Command, args []string) {
		str, _ := cmd.Flags().GetString("str")
		hash := sha512.Sum512([]byte(str))
		hashStr := hex.EncodeToString(hash[:])
		fmt.Println(hashStr)
	},
}

func init() {
	rootCmd.AddCommand(sha512Cmd)
	sha512Cmd.Flags().StringP("str", "S", "", "Operated Str")
	sha512Cmd.MarkFlagRequired("str")
}
