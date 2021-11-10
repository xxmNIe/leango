package main

import (
	"fmt"
	"sort"
)

func analysisHistogram(heights []int, cnt int) []int {
	sort.Sort(sort.IntSlice(heights))
	n, i, j := len(heights), 0, 0
	for ; i+cnt <= n; i++ {
		if heights[i+cnt-1]-heights[i] < heights[j+cnt-1]-heights[j] {
			j = i
		}
	}
	ans := []int{}
	ans = append(ans, heights[j:j+cnt]...)
	return ans
}

// class Solution {
// 	public:
// 		vector<int> analysisHistogram(vector<int>& heights, int cnt) {
// 			sort(heights.begin(),heights.end());
// 			int n=heights.size(),i,j;
// 			for(i=j=0;i+cnt<=n;i++)if(heights[i+cnt-1]-heights[i]<heights[j+cnt-1]-heights[j])j=i;
// 			vector<int> ans;
// 			for(i=0;i<cnt;i++)ans.push_back(heights[i+j]);
// 			return ans;
// 		}
// 	};

func main() {
	fmt.Println(analysisHistogram([]int{3, 2, 7, 6, 1, 8}, 3))
}
