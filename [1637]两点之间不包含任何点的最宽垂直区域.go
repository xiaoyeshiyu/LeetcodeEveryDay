func maxWidthOfVerticalArea(points [][]int) int {
    sort.Slice(points,func(i,j int)bool {
        return points[i][0] < points[j][0]
    })

    var ans int 
    for i:= 1 ;i < len(points);i++ {
        tmp := points[i][0] - points[i-1][0]
        if tmp > ans {
            ans = tmp
        }
    }
    return ans 
}
