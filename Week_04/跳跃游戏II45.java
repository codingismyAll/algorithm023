public class 跳跃游戏II45 {
    //时间复杂度：O(n^2)，其中 n 是数组长度。有两层嵌套循环，在最坏的情况下，例如数组中的所有元素都是 1，
    // position 需要遍历数组中的每个位置，对于 position 的每个值都有一次循环。
    //空间复杂度：O(1)。
    //反向查找
    public int jump(int[] nums) {
        int position = nums.length - 1;
        int step = 0;
        while (position > 0) {
            for (int i = 0; i < position; i++) {
                if ((nums[i] + i) >= position) {
                    step++;
                    position = i;
                    break;
                }
            }
        }
        return step;
    }

//时间复杂度：O(n)，其中 nn 是数组长度。
//
//空间复杂度：O(1)。
    public int Jump2(int[] nums) {
        int length = nums.length;
        int end = 0;
        int maxPosition = 0;
        int steps = 0;
        for (int i = 0; i < length-1; i++) {
            maxPosition = Math.max(maxPosition, i + nums[i]);
            if (i == end) {
                end = maxPosition;
                steps++;
            }
        }
        return steps;
    }

    public static void main(String[] args) {
        int[] arr = {2,3,1,2,4,2,3};
        跳跃游戏II45 t = new 跳跃游戏II45();
        t.Jump2(arr);
    }

}
