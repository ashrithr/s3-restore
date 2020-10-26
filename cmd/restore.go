package cmd

import (
	"errors"
	"fmt"

	"github.com/ashrithr/s3-restore/internal/restore"
	"github.com/spf13/cobra"
)

var (
	copy      bool
	del       bool
	dstBucket string
	dstPrefix string
)

// restoreCmd represents the restore command
var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restores deleted objects.",
	Long: `Restores deleted objects either by removing delete markers or by
copying from previous versions. Use:

	--del to remove the delete markers
	--copy to copy latest previous version of the object`,
	Args: func(cmd *cobra.Command, args []string) error {
		if !copy && !del {
			return errors.New("requires either '--del' or '--copy' flags")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Dry run enabled: %v\n", !noDryrun)

		if del {
			restore.ObjsUsingDel(bucket, prefix, region, noDryrun)
		} else {
			restore.ObjsUingCopy(bucket, prefix, region, dstBucket, dstPrefix, noDryrun)
		}
	},
}

func init() {
	rootCmd.AddCommand(restoreCmd)

	restoreCmd.Flags().BoolVar(&del, "del", false, "restore objects by removing delete markers")
	restoreCmd.Flags().BoolVar(&copy, "copy", false, "restore objects by copying previous latest version of the objects")
	restoreCmd.Flags().StringVar(&dstBucket, "dstBucket", "", "Optionally specify another destination bucket to copy the objects to")
	restoreCmd.Flags().StringVar(&dstPrefix, "dstPrefix", "", "Optionally specify destination prefix")
}
