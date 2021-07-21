package main

import (
  "fmt"
)

type TreeNode1 struct {
  Val int
  Left *TreeNode
  Right *TreeNode
}

func main() {
  t := &TreeNode{Val: 1,
                 Left: &TreeNode{Val: 2,
                                 Left: &TreeNode{Val: 4},
                                },
                 Right: &TreeNode{Val: 5,
                                  Left: &TreeNode{Val: 6},
                                  Right: &TreeNode{Val: 8}},

                  }
  t2 := &TreeNode{Val: 1,
    Left: &TreeNode{Val: 2,
      Left: &TreeNode{Val: 4},
    },
    Right: &TreeNode{Val: 5,
      Left: &TreeNode{Val: 6},
      Right: &TreeNode{Val: 9}},

  }
                  fmt.Println(leafSimilar(t,t2))
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {

  sqe := make([]int,0)
  flag := true
  walk := func(*TreeNode,) {}
  walk = func(r *TreeNode) {
    if r.Left ==nil && r.Right ==nil {
      sqe = append(sqe, r.Val)
      return
    }
    if  r.Left!=nil {
      walk(r.Left)
    }
    if r.Right !=nil {
      walk(r.Right)
    }
  }
  walk(root1)
  tmp := make([]int,len(sqe))
  copy(tmp,sqe)
  sqe = make([]int,0)
  walk(root2)
  if len(tmp)!=len(sqe){
    return false
  }
  for i,v:=range sqe {
    if v!=tmp[i]{
      return false
    }
  }

  return flag
}

