package main
import "fmt"

func arr (x []int,a []int){

	for t:=0;t<=len(x)-1;t++{
	    var ans bool=x[t]==a[0]
         fmt.Println (ans)
	}
	
	}
func main(){
var y[]int=[]int{1,2,3,4,5}
var z[]int=[]int{3}
arr(y,z)

}

