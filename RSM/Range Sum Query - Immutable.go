type NumArray struct {
    presum []int    // �洢 [0,i] �ĺ�
}


func Constructor(nums []int) NumArray {
    a := NumArray{}
    a.presum = append(a.presum, 0)  // ��ʼ�� presum ����
    for i:=1; i<=len(nums); i++ {
        t := a.presum[i-1] + nums[i-1]
        a.presum = append(a.presum, t)
    }
    return a
}


func (this *NumArray) SumRange(i int, j int) int {
    return this.presum[j+1] - this.presum[i]
}



/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
