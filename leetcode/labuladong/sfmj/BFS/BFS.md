# BFS 算法框架
func BFS(start, target node) {
    queue := make([]node, 0)
    visit := make(map[int]struct{}, 0)
    // 使用队列记录每层的节点
    queue = append(queue, start)
    // 记录走过的路径
    visit = append(visit, start)
    // 记录步数
    step := 0

    for len(queue) != 0 {
        sz := len(queue)
        for i := 0; i < sz; i++ {
            cur := queue[len(queue)-1]
            queue = queue[:len(queue)-1]
            // 到达终点
            if cur == target {
                return step
            }
            // 遍历cur的邻居，如果未被访问，加入queue和visisted
        }
        step++
    }
}