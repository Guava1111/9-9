package main
import "fmt"

func arr (x []int,a []int){
	fmt.Println(x,a)
	for t:=0;t<=len(x)-1;t++{
		for e:=0;e<=len(a)-1;e++{
			if x[t]==a[e]{
			fmt.Println("是")
			fmt.Println("相同數:",x[t])
		} else {
			
		}
	}

}
}
func main(){
var y[]int=[]int{1,2,3,4,5}
var z[]int=[]int{2,5,10}
arr(y,z)

}