package cmd 
import (
	"fmt"
	"github.com/spf13/cobra"
)

var message string
var defaultmessage string
var msgCmd = &cobra.Command{
	Use: "msg",
	Short: "This command will take the text as input.",
	Run: func(cmd *cobra.Command, args []string){
		
		if message!=""{
		fmt.Println("Input message : ",message)
		tts_converter(message)
		}else{
			defaultmessage = "hello give a custom input"
			fmt.Println("Input message : ",defaultmessage)
			tts_converter(defaultmessage)
		}
	},
}

func init(){
	rootCmd.AddCommand(msgCmd)
    msgCmd.Flags().StringVarP(&message, "message", "m", "", "Message to process")
}