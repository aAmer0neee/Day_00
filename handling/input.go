package handling
import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"github.com/aAmer0neee/Day_00/stats"
)

func reader() []int {
	NumReader := bufio.NewReader(os.Stdin)
	var nums []int

	fmt.Println("Enter the numbers (press Ctrl+D to finish):")

	for {
		input, err := NumReader.ReadString('\n')

		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Printf("Error: '%s' not a number\n", input)
			continue
		}

		input = input[:len(input)-1]
		if input == "" {
			fmt.Println("Error: empty string")
			continue
		}
		num, err := strconv.Atoi(input)

		if err != nil {
			fmt.Printf("Error: '%s' not a number\n", input)
			continue
		}

		if -100000 < num && num < 100000 {
			nums = append(nums, num)
		} else {
			fmt.Printf("Error: '%d' out of range\n", num)
			continue
		}

	}
	if len(nums) == 1 {
		nums = stats.SingleNumberSequence(nums)
	}
	sort.Ints(nums)

	return nums
}



func MetricFlag() {
	MeanFlag := flag.Bool("mean", false, "display Mean")
	MedianFlag := flag.Bool("median", false, "display Median")
	ModeFlag := flag.Bool("mode", false, "display Mode")
	SDFlag := flag.Bool("sd", false, "display Standard Deviation")

	flag.Parse()

	result := reader()

	if 0 == flag.NFlag() {
		fmt.Printf("Mean: %.2f\n",  stats.Mean(result))
		fmt.Printf("Median: %.2f\n", stats.Median(result))
		fmt.Printf("Mode: %.2f\n", stats.Median(result))
		fmt.Printf("SD: %.2f\n", stats.StandardDeviation(result))
	} else {
		if *MeanFlag {
			fmt.Printf("Mean: %.2f\n", stats.Mean(result))
		}
		if *MedianFlag {
			fmt.Printf("Median: %.2f\n", stats.Median(result))
		}
		if *ModeFlag {
			fmt.Printf("Mode: %.2f\n", stats.Mode(result))
		}
		if *SDFlag {
			fmt.Printf("SD: %.2f\n", stats.StandardDeviation(result))
		}
	}
}