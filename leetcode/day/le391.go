package main
func isRectangleCover(rectangles [][]int) bool {
	type point struct{
		x,y int
	}
	area,minX,minY,maxX,maxY :=0,rectangles[0][0],rectangles[0][1],rectangles[0][2],rectangles[0][3]
	cnt :=map[point]int{}
	for _,rct := range rectangles{
	x,y,a,b := rct[0],rct[1],rct[2],rct[3]

	area += (a-x)*(b-y)
	minX = min(minX,x)
	minY = min(minY,y)
	maxX = max(maxX,a)
	maxY = max(maxY,b)

	cnt[point{x,y}]++
	cnt[point{a,b}]++
	cnt[point{x,b}]++
	cnt[point{a,y}]++
	
	}
	if area != (maxX-minX) * (maxY - minY) || cnt[point{minX,minY}] !=1 || cnt[point{maxX,maxY}] !=1 || cnt[point{minX,maxY}]!=1 || cnt[point{maxX,minY}]!=1{
		return false
	}

	delete(cnt,point{minX,minY})
	delete(cnt,point{maxX,maxY})
	delete(cnt,point{minX,maxY})
	delete(cnt,point{maxX,minY})

	for _,c :=range cnt{
		if c!=2 && c!=4{
			return false
		}
	}
	return true
}

func min(x,y int)int {
	if x < y{
		return x
	}
	return y
}

func max(x,y int)int {
	if x>y{
		return x
	}
	return y
}