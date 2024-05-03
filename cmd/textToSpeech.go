package cmd

import (
	"fmt"
	"os/exec"


	htgotts "github.com/hegedustibor/htgo-tts"
	voices "github.com/hegedustibor/htgo-tts/voices"
)

func tts_converter(text string){
	speech := htgotts.Speech{
		Folder: "audio", 
		Language: voices.English,
	}
	speech.Speak(text)
    cmd := exec.Command("./script.sh")

    // Run the command and capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing script:", err)
	}
	fmt.Println("Script output:", string(output))
}