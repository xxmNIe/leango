package main

type TreeNode struct {
  Val int
  Left *TreeNode
  Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
//
    var xparent,yparent *TreeNode
    var xdep,ydep int
    var xfond,yfond bool
    var dfs func(node,parent *TreeNode,deep int)
    dfs = func(node, parent *TreeNode, deep int) {
        if node ==nil {
            return
        }
        if node.Val ==x{
            xparent,xdep,xfond = parent,deep,true
        }else if node.Val == y  {

            yparent,ydep,yfond = parent,deep,true

        }
        if xfond && yfond {
            return
        }
        dfs(node.Left,node,deep+1)
        if xfond && yfond {
            return
        }
        dfs(node.Right,node,deep+1)
    }
    dfs(root,nil,0)
    return xparent!=yparent && xdep==ydep
    //link :=make([]*TreeNode,64)
    //pr :=0
    //link[pr] = root
    //pr++
    //p :=link[pr-1]
    //for p!=nil {
    //    if p
    //}
}

func ab(root *TreeNode, x int, y int)bool{
    var xParent, yParent *TreeNode
    var xDepth, yDepth int
    var xFound, yFound bool

    var dfs func(node, parent *TreeNode, depth int)
    dfs = func(node, parent *TreeNode, depth int) {
        if node == nil {
            return
        }

        if node.Val == x {
            xParent, xDepth, xFound = parent, depth, true
        } else if node.Val == y {
            yParent, yDepth, yFound = parent, depth, true
        }

        // 如果两个节点都找到了，就可以提前退出遍历
        // 即使不提前退出，对最坏情况下的时间复杂度也不会有影响
        if xFound && yFound {
            return
        }

        dfs(node.Left, node, depth+1)

        if xFound && yFound {
            return
        }

        dfs(node.Right, node, depth+1)
    }
    dfs(root, nil, 0)

    return xDepth == yDepth && xParent != yParent

}