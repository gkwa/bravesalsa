package cmd

import (
   "context"
   "fmt"
   "os"

   "github.com/go-logr/logr"
   "github.com/spf13/cobra"
   "github.com/spf13/viper"

   "github.com/gkwa/bravesalsa/core"
   "github.com/gkwa/bravesalsa/internal/logger"
)

var (
   cfgFile   string
   verbose   bool
   logFormat string
   cliLogger logr.Logger
   reverse   bool
)

var rootCmd = &cobra.Command{
   Use:   "bravesalsa",
   Short: "Sort file paths based on modification time",
   Long:  `A command-line tool to sort file paths based on their most recent modification time.`,
   RunE: func(cmd *cobra.Command, args []string) error {
   	logger := LoggerFrom(cmd.Context())
   	err := core.SortFiles(cmd.InOrStdin(), cmd.OutOrStdout(), reverse)
   	if err != nil {
   		logger.Error(err, "Failed to sort files")
   		return err
   	}
   	return nil
   },
}

func Execute() {
   err := rootCmd.Execute()
   if err != nil {
   	os.Exit(1)
   }
}

func init() {
   cobra.OnInitialize(initConfig)

   rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bravesalsa.yaml)")
   rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose mode")
   rootCmd.PersistentFlags().StringVar(&logFormat, "log-format", "", "json or text (default is text)")
   rootCmd.Flags().BoolVarP(&reverse, "reverse", "r", false, "Sort in reverse order")

   if err := viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose")); err != nil {
   	fmt.Printf("Error binding verbose flag: %v\n", err)
   	os.Exit(1)
   }
   if err := viper.BindPFlag("log-format", rootCmd.PersistentFlags().Lookup("log-format")); err != nil {
   	fmt.Printf("Error binding log-format flag: %v\n", err)
   	os.Exit(1)
   }
}

func initConfig() {
   if cfgFile != "" {
   	viper.SetConfigFile(cfgFile)
   } else {
   	home, err := os.UserHomeDir()
   	cobra.CheckErr(err)

   	viper.AddConfigPath(home)
   	viper.SetConfigType("yaml")
   	viper.SetConfigName(".bravesalsa")
   }

   viper.AutomaticEnv()

   if err := viper.ReadInConfig(); err == nil {
   	fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
   }

   logFormat = viper.GetString("log-format")
   verbose = viper.GetBool("verbose")
}

func LoggerFrom(ctx context.Context, keysAndValues ...interface{}) logr.Logger {
   if cliLogger.IsZero() {
   	cliLogger = logger.NewConsoleLogger(verbose, logFormat == "json")
   }
   newLogger := cliLogger
   if ctx != nil {
   	if l, err := logr.FromContext(ctx); err == nil {
   		newLogger = l
   	}
   }
   return newLogger.WithValues(keysAndValues...)
}

