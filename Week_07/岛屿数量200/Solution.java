package 岛屿数量200;

public class Solution {
    //    private int m;
//    private int n;
    public int NumIslands(char[][] grid) {
        int m = grid.length;
        if (m == 0) return 0;
        int n = grid[0].length;
        if (n == 0) return 0;
        int count = 0;
        for (int i = 0; i < grid.length; i++) {
            for (int j = 0; j < grid[i].length; j++) {
                if (grid[i][j] == '1') {
                    dfs(m, n, i, j, grid);
                    count++;
                }
            }
        }
        System.out.println(count);
        return count;

    }

    private void dfs(int m, int n, int i, int j, char[][] grid) {
        if (i < 0 || j < 0 || i >= m || j >= n || grid[i][j] != '1') {
            return;
        }
        //process logic
        grid[i][j] = 0;
        //drill down
        dfs(m, n, i + 1, j, grid);
        dfs(m, n, i - 1, j, grid);
        dfs(m, n, i, j + 1, grid);
        dfs(m, n, i, j - 1, grid);
    }


    //第二遍
    public int numIslands2(char[][] grid) {

        int m = grid.length;
        if (m == 0) return m;
        int n = grid[0].length;
        if (n == 0) return n;
        int count = 0;
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (grid[i][j] == '1') {
                    count++;
                    dfs2(m, n, i, j, grid);
                }
            }
        }

        return count;

    }

    private void dfs2(int m, int n, int i, int j, char[][] grid) {
        if (i < 0 || j < 0 || i >= m || j >= n || grid[i][j] != '1') {
            return;
        }
        grid[i][j] = '0';
        dfs2(m, n, i + 1, j, grid);
        dfs2(m, n, i - 1, j, grid);
        dfs2(m, n, i, j + 1, grid);
        dfs2(m, n, i, j - 1, grid);
    }

}
