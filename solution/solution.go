package main
import (
	"strconv"
	"fmt"
)

// solution 1
func sockMerchant(n int32, ar []int32) {
    var pairs int32

    for i := 0 ; int32(i) < n; i++ {
        for j := i + 1; int32(j) < n; j++ {
            if (ar[i] == ar[j] && ar[i] != 0) {
                pairs++
                ar[i] = 0
                ar[j] = 0
                break    
            }
        }
    }

    fmt.Println(pairs) 
}

// solution 2
func countingValleys(steps int32, path string) {  
    var v int32
    var pos int
    var isDown bool = false
    
    for i := 0; i < int(steps); i++ {
        if (path[i] == 'U') {
            if (pos == -1 && isDown) {
                v += 1
                isDown = false
            }
            pos += 1
        }  
        if (path[i] == 'D') {
            if (pos == 0) {
                isDown = true
            }
            pos -= 1;
        }
    }
    
    fmt.Println(v)
}

// solution 3
func separateNums(num int) {
	str := strconv.Itoa(num)
	for i := 0; i < len(str); i++ {
		var nums string = string(str[i])
		for z := len(str)-i; z > 1; z-- {
			nums += "0"
		}
		fmt.Println(nums)
		nums = ""
	}
}

// solution 4
func lampSwitches() {
	var on int
	var lamps [100]bool
	for i:= 0; i < 100; i++ {
		lamps[i] = true
	}

	for n:= 1; n < 100; n++ {
		var current int = n + 1	
		for i:= current; i <= 100; i++ {
			if (i % current == 0) {
				lamps[i-1] = !lamps[i-1]
			}
		}
	}

	for i:= 0; i < 100; i++ {
		if (lamps[i] == true) {
			on++
		}
	}

	fmt.Println(on)
}

func main() {
	// problem 1
	var s1_n int32 = 9
	var s1_arr = []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}
	fmt.Println("Problem 1, output:")
	sockMerchant(s1_n, s1_arr)
	fmt.Println("")

	// problem 2
	var s2_steps int32 = 8
	var s2_path string = "UDDDUDUU"
	fmt.Println("Problem 2, output:")
	countingValleys(s2_steps, s2_path)
	fmt.Println("")

	// problem 3
	var s3_num int = 1345679
	fmt.Println("Problem 3, output:")
	separateNums(s3_num)
	fmt.Println("")

	// problem 4
	fmt.Println("Problem 4, output:")
	lampSwitches()
	fmt.Println("")
}