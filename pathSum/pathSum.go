package pathSum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	path = [][]int{}
)

func pathSum(root *TreeNode, target int) [][]int {
	path = make([][]int, 0)

	var dfs func(root *TreeNode, cur int, tmp []int)
	dfs = func(root *TreeNode, cur int, tmp []int) {
		if root == nil {
			return
		}
		cur += root.Val
		tmp = append(tmp, root.Val)
		if root.Left == nil && root.Right == nil {
			if cur == target {
				path = append(path, append([]int{}, tmp...))
			}
		}
		if root.Left != nil {
			dfs(root.Left, cur, tmp)
		}
		if root.Right != nil {
			dfs(root.Right, cur, tmp)
		}
	}
	dfs(root, 0, []int{})
	return path
}
