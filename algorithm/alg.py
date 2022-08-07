# coding:utf-8

from Queue import LifoQueue

### 二分查找（找到返回下标，否则返回 -1）
def binary_search(arr, start, end, target):
    while end >= start:
        mid = (start+end)/2
        if arr[mid] == target:
            return mid
        if arr[mid] > target:
            end = mid-1
            continue
        start = mid+1
    return -1

### 后缀表达式计算
### 符号支持: +,-,*,/,%
### 表达式用字符串数组表示，操作数为整数，输入表达式确保正确，且长度大于 0
def cal_post_expression(arr):
  stack = LifoQueue()
  ops = ['+', "-", "*", "/", "%"]
  for c in arr:
    if c in ops:
      b = stack.get(False)
      a = stack.get(False)
      stack.put(_cal(a,c,b))
      continue
    stack.put(int(c))
  return stack.get(False)

def _cal(a, op, b):
  if op == "+":
    return a+b
  if op == "-":
    return a-b
  if op == "*":
    return a*b
  if op == "/":
    return a/b
  return a%b

### 中缀表达式转换成后缀表达式
### 符号支持: +,-,*,/,%,(,)
### 表达式用字符串数组表示，输入表达式确保正确, 且长度大于 0
def expression_mid2post(arr):
  ops = ['+', "-", "*", "/", "%", "(", ")"]
  in_p = {
    "#": 0,
    "+": 3, "-": 3,
    "*": 5, "/": 5, "%": 5,
    "(": 1,
    ")": 6
  }
  out_p = {
    "#": 0,
    "+": 2, "-": 2,
    "*": 4, "/": 4, "%": 4,
    "(": 6,
    ")": 1
  }
  result = []
  stack = LifoQueue()
  stack.put('#')
  for c in arr+['#']:
    # 退栈直到最开始的'#'
    if c == '#':
      while True:
        x = stack.get(False)
        if x == '#':
          break
        result.append(x)
      return result
    # 退栈到第一个'('
    if c == ')':
      while True:
        x = stack.get(False) 
        if x == '(':
          break
        result.append(x)
      continue
    # 遇到操作符
    if c in ops:
      # LifoQueue doesn't have top() method
      top = stack.get(False)
      stack.put(top)
      # 如果 out_p[c] > in_p[top], c 入栈
      if out_p[c] > in_p[top]:
        stack.put(c)
        continue
      # 否则, 退栈直到满足 out_p[c] > in_p[top]
      while True:
        x = stack.get(False)
        if out_p[c] > in_p[x]:
          stack.put(x)
          stack.put(c)
          break
        result.append(x)
      continue
    # 遇到操作数，直接输出
    result.append(c)

### 排列
### arr 的元素各不相同, 元素个数大于 0, n 为排列个数
### 返回生成器
def perm_generator(arr, n=None):
  if not n:
    n = len(arr)
  for i in range(0, len(arr)):
    p = arr[i:i+1]
    if n == 1:
      yield p
      continue
    for j in perm_generator(arr[0:i]+arr[i+1:], n-1):
      yield p+j

### 组合
### arr 的元素各不相同, 元素个数大于 0, n 为组合个数
### 返回生成器
def comb_generator(arr, n=None):
  if not n:
    n = len(arr)
  for i in range(0, len(arr)+1-n):
    p = arr[i:i+1]
    if n == 1:
      yield p
      continue
    for j in comb_generator(arr[i+1:], n-1):
      yield p+j

### 二叉树遍历
class TreeNode(object):
  def __init__(self, val, left=None, right=None):
    self.val = val
    self.left = left
    self.right = right

## 递归中序遍历
def tree_recursive_mid_visit(root, result):
  if not root:
    return
  tree_recursive_mid_visit(root.left, result)
  result.append(root.val)
  tree_recursive_mid_visit(root.right, result)

## 迭代中序遍历
def tree_mid_visit(root, result):
  stack = LifoQueue()
  while root:
    if root.left:
      stack.put(root)
      root = root.left
      continue
    result.append(root.val)
    if root.right:
      root = root.right
      continue
    while True:
      if stack.empty():
        return
      root = stack.get()
      result.append(root.val)
      if root.right:
        root = root.right
        break

## 迭代前序遍历
def tree_front_visit(root, result):
  stack = LifoQueue()
  while root:
    result.append(root.val)
    if root.left:
      if root.right:
        stack.put(root.right)
      root = root.left
      continue
    if root.right:
      root = root.right
      continue
    if stack.empty():
      break
    root = stack.get()

## 迭代后序遍历
def tree_rear_visit(root, result):
  stack = LifoQueue()
  while root:
    if root.left:
      # (root, false) false 表示right还没有遍历
      stack.put((root, False))
      root = root.left
      continue
    if root.right:
      stack.put((root, True))
      root = root.right
      continue
    result.append(root.val)
    while True:
      if stack.empty():
        return
      root, right_visited = stack.get()
      if right_visited or not root.right:
        result.append(root.val)
        continue
      stack.put((root, True))
      root = root.right
      break
