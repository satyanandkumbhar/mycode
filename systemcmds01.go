/* Alta3 Research | RZFeeser
   Executing system commands  */

package main

import (
	"fmt"
    "log"
    "os/exec"
)

func main() {

    // prepares a "cmd" struct
    out, err := exec.Command("ls").Output()

    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%s", out)
}


