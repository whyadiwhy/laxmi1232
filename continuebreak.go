
package main
import (
	"fmt"
)

func main() {
  var a int = 19
  for a<20{
    if  a == 15{
      a++
      continue
      }
    fmt.Printf("value of a is %d\n",a)
    break
    }
  
  }
  