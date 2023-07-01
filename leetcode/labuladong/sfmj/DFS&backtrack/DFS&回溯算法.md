# DFS&回溯算法
回溯算法效率一般不高，但却是最好用的算法

## 回溯算法框架
result = []
func backtrack(路径, 选择列表):
    if 满足结束条件： {
        result = append(result, 路径)
        return
    }
    for 选择 := range 选择列表：{
        做选择
        backtrack(路径，选择列表)
        撤销选择
    }

## 子集 元素无重复不可复选
res := make([][]int, 0)
path := make([]int, 0)
var backtrack func(nums []int, index int)
backtrack = func(nums []int, index int) {
    res = append(res, append([]int{}, path...))
    for i := index; i < len(nums); i++ {
        path = append(path, nums[i])
        backtrack(nums, i+1)
        path = path[:len(path)-1]
    }
}
backtrack(nums, 0)
return res

用一个索引作为指针遍历整个数组，保证不会重复

## 组合 元素无重复不可复选
求子集长度为某个值的所有子集，
在上述代码上res添加处修改下即可。

res := make([][]int, 0)
path := make([]int, 0)
var backtrack func(nums []int, index int)
backtrack = func(nums []int, index int) {
    if len(path) == k {
        res = append(res, append([]int{}, path...))
        return
    }
    for i := index; i < len(nums); i++ {
        path = append(path, nums[i])
        backtrack(nums, i+1)
        path = path[:len(path)-1]
    }
}
backtrack(nums, 0)
return res

## 排列 元素无重复不可复选
求子集长度等于数组长度的所有子集，
在上述代码上res添加处修改下,并使用一个数组记录下每个元素是否使用即可。

res := make([][]int, 0)
path := make([]int, 0)
used := make([]bool, len(nums))
var backtrack func(nums []int)
backtrack = func(nums []int) {
    if len(path) == len(nums) {
        res = append(res, append([]int{}, path...))
        return
    }
    for i := 0; i < len(nums); i++ {
        if used[i] {
            continue
        }
        path = append(path, nums[i])
        used[i] = true
        backtrack(nums)
        path = path[:len(path)-1]
        used[i] = false
    }
}
backtrack(nums)
return res

使用不使用used，取决于能否出现nums[i]之后能不能出现nums[i]左边的元素，可以就使用，不可以就不需要

## 子集/组合 元素可重复不可复选
先排序，然后跳过相邻重复的元素

### 子集
res := make([][]int, 0)
path := make([]int, 0)
var backtrack func(nums []int, index int)
backtrack = func(nums []int, index int) {
    res = append(res, append([]int{}, path...))
    for i := index; i < len(nums); i++ {
        if i > index && nums[i] == nums[i-1] {
            continue
        }
        path = append(path, nums[i])
        backtrack(nums, i+1)
        path = path[:len(path)-1]
    }
}
sort.Ints(nums)
backtrack(nums, 0)
return res

### 组合
需要用额外的sum来记录和,只有当sum等于target时才加入

res := make([][]int, 0)
path := make([]int, 0)
sum := 0
var backtrack func(nums []int, index int)
backtrack = func(nums []int, index int) {
    if sum == target {
        res = append(res, append([]int{}, path...))
        return
    }
    // 优化，避免多次遍历
    if sum > target {
        return
    }
    for i := index; i < len(nums); i++ {
        if i > index && nums[i] == nums[i-1] {
            continue
        }
        // 优化，避免多次遍历
        if sum+nums[i] > target {
            return
        }
        path = append(path, nums[i])
        sum += nums[i]
        backtrack(nums, i+1)
        path = path[:len(path)-1]
        sum -= nums[i]
    }
}
sort.Ints(nums)
backtrack(nums, 0)
return res

## 排列 元素可重复不可复选
根据上述的排列问题，做下可重复处理即先将数组排序，然后跳过相邻的元素

res := make([][]int, 0)
path := make([]int, 0)
used := make([]bool, len(nums))
var backtrack func(nums []int)
backtrack = func(nums []int) {
    if len(path) == len(nums) {
        res = append(res, append([]int{}, path...))
        return
    }
    for i := 0; i < len(nums); i++ {
        if used[i] {
            continue
        }
        // 相邻元素跳过
        if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
            continue
        }
        path = append(path, nums[i])
        used[i] = true
        backtrack(nums)
        path = path[:len(path)-1]
        used[i] = false
    }
}
// 排序
sort.Ints(nums)
backtrack(nums)
return res

## 子集/组合 元素无重复可复选
子集和组合问题是通过一个start变量来保证不重复，
下次递归从start+1开始，所以不重复

### 组合
res := make([][]int, 0)
path := make([]int, 0)
sum := 0
var backtrack func(nums []int, start int)
backtrack = func(nums []int, start int) {
    if sum == target {
        res = append(res, append([]int{}, path...))
        return
    }
    if sum > target {
        return
    }
    for i := start; i < len(nums); i++ {
        path = append(path, nums[i])
        sum += nums[i]
        backtrack(nums, i)
        path = path[:len(path)-1]
        sum -= nums[i]
    }
}
backtrack(candidates, 0)
return res

### 子集
将上述代码的sum部分去掉,start也去掉

res := make([][]int, 0)
path := make([]int, 0)
var backtrack func(nums []int)
backtrack = func(nums []int) {
    if len(path) == len(nums) {
        res = append(res, append([]int{}, path...))
        return
    }
    for i := 0; i < len(nums); i++ {
        path = append(path, nums[i])
        backtrack(nums, i)
        path = path[:len(path)-1]
    }
}
backtrack(candidates)
return res

## 总结
子集问题和组合问题本质相同，只是结束条件有区别
框架代码见上

# DFS遍历二维数组 框架
var dfs func(grid [][]int, i, j int)
dfs = func(grid [][]int, i, j int) {
    if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
        return
    }
    if grid[i][j] == 1 {
        return
    }
    grid[i][j] = 1
    dfs(grid, i-1, j)
    dfs(grid, i+1, j)
    dfs(grid, i, j-1)
    dfs(grid, i, j+1)
}

## 判断岛屿数量
直接套用上述的框架，遍历一遍，计数，然后用上述的dfs淹掉

## 计算飞地面积
先将四周边界的淹掉，然后再计数， 计数之后不需要淹掉岛屿
// 把上下边界的岛屿淹掉
for i := 0; i < len(grid[0]); i++ {
    dfs(grid, 0, i)
    dfs(grid, len(grid)-1, i)
}
// 把左右边界的岛屿淹掉
for i := 0; i < len(grid); i++ {
    dfs(grid, i, 0)
    dfs(grid, i, len(grid[0])-1)
}

## 判断封闭岛屿的数量
先将四周边界的淹掉，然后再计数，计数之后需要淹掉岛屿

## 判断最大岛屿面积
修改上述的dfs，返回值增加岛屿面积，遍历的时候遇到岛屿更新面积
var dfs func(grid [][]int, i, j int) int
dfs = func(grid [][]int, i, j int) int {
    if i < 0 || j < 0 || i >= m || j >= n {
        return 0
    }
    if grid[i][j] == 0 {
        return 0
    }
    grid[i][j] = 0
    return dfs(grid, i-1, j) + dfs(grid, i+1, j) + dfs(grid, i, j-1) + dfs(grid, i, j+1) + 1
}