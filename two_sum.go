package main

import (
	"fmt"
	"sort"
)

func twoSum(nums []int, target int) []int {
	for i:=0;i<len(nums);i++{
		for j:=i+1;j<len(nums);j++{
			if nums[i]+nums[j] == target{
				return []int{i,j}
			}
		}
	}
	sort.Ints(nums)

	return nums;
}

func longestPalindrome(s string) string {

	return s
}



func main() {
	fmt.Print("hello")
	//var cmd = &cobra.Command{
	//	Use:   "greet [name]",
	//	Short: "A brief description of your command",
	//	Long: `A longer description that spans multiple lines and provides
    //    more detailed information about the command.`,
	//	Args: cobra.ExactArgs(1),
	//	Run: func(cmd *cobra.Command, args []string) {
	//		fmt.Printf("Hello, %s!\n", args[0])
	//	},
	//}
	//
	//cmd.Execute()
}


