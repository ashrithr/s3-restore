package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	bucket   string
	prefix   string
	region   string
	noDryrun bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "s3-restore",
	Short: "Restores deleted objects of an S3 version enabled bucket.",
	Long:  `Restores deleted objects of an S3 version enabled bucket.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&bucket, "bucket", "", "S3 bucket name to restore the objects from")
	rootCmd.PersistentFlags().StringVar(&prefix, "prefix", "", "S3 prefix to look for delete objects")
	rootCmd.PersistentFlags().StringVar(&region, "region", "us-east-1", "AWS region where S3 bucket is located")
	rootCmd.PersistentFlags().BoolVar(&noDryrun, "noDryrun", false, "Disable dryrun, default behavior is to run in drymode")

	rootCmd.MarkPersistentFlagRequired("bucket")
}
