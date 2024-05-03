package cmd 

import (
	"github.com/spf13/cobra"
	"os"
	"fmt"
)

var rootCmd = &cobra.Command{
	Use: "TextToSpeech",
	Short: "This app will convert text to speech.",
	Long: "This cli app take a text as input and then convert it into speech . ",
}

func Execute(){
	if err:= rootCmd.Execute(); err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}
}