package cmd

import (
	"github.com/shanedoc/gohub/pkg/cache"
	"github.com/shanedoc/gohub/pkg/console"
	"github.com/spf13/cobra"
)

var CmdCache = &cobra.Command{
	Use:   "cache",
	Short: "Cache management",
}

var CmdCacheClear = &cobra.Command{
	Use:   "clear",
	Short: "Clear cache",
	Run:   runClearCache,
}

func init() {
	//注册cache命令的字命令
	CmdCache.AddCommand(CmdCacheClear)
}

func runClearCache(cmd *cobra.Command, args []string) {
	cache.Flush()
	console.Success("Cache cleared")
}
