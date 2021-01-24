public class 第k个数 {
    public int getKthMagicNumber(int k) {
        int[] dp = new int[k];
        int p3 = 0, p5 = 0, p7 = 0;
        dp[0] =1;
        for (int i = 1; i < k; i++) {
            // 选出最小的数字
            dp[i] = Math.min(dp[p3] * 3, Math.min(dp[p5] * 5, dp[p7] * 7));
            // 将该数字对应的指针向前移动一步。
            if( dp[i] == dp[p3] * 3 ) p3++;
            if( dp[i] == dp[p5] * 5 ) p5++;
            if( dp[i] == dp[p7] * 7 ) p7++;
        }
        return dp[k-1];

    }
}
