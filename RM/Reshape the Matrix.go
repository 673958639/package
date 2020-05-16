func matrixReshape(nums [][]int, r int, c int) [][]int {
    row,col:=len(nums),len(nums[0])
    if row*col!=r*c{
        return nums
    }
    res:=make([][]int,r)
    for i:=0;i<len(res);i++{
        res[i]=make([]int,c)
    }
    
    for i:=0;i<r*c;i++{
        res[i/c][i%c]=nums[i/col][i%col]
    }
    return res
}

