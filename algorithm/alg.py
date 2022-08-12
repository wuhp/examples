# coding:utf-8

from Queue import LifoQueue
from copy import deepcopy

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

### 排序
## 归并排序
def merge_sort(arr):
  if len(arr) <= 1:
    return arr
  mid = len(arr)/2
  arr1 = merge_sort(arr[0:mid])
  arr2 = merge_sort(arr[mid:])
  return _merge(arr1,arr2)

def _merge(arr1, arr2):
  result = []
  i = j = 0
  while True:
    if i == len(arr1):
      result.extend(arr2[j:])
      break
    if j == len(arr2):
      result.extend(arr1[i:])
      break
    if arr1[i] < arr2[j]:
      result.append(arr1[i])
      i = i+1
      continue
    result.append(arr2[j])
    j = j+1
  return result

## 快速排序
def quick_sort(arr):
  _quick_sort(arr, 0, len(arr)-1)

def _quick_sort(arr, start, end):
  if start >= end:
    return
  p1 = start
  p2 = end
  while p1 <= p2:
    if arr[p2] >= arr[start]:
      p2 = p2-1
      continue
    if arr[p1] <= arr[start]:
      p1 = p1+1
      continue
    arr[p1], arr[p2] = arr[p2], arr[p1]
    p2 = p2-1
    p1 = p1+1
  if p2 < start:
    p2 = start
  arr[p2], arr[start] = arr[start], arr[p2]
  _quick_sort(arr, start, p2-1)
  _quick_sort(arr, p2+1, end)

### 并查集 & 图的连通区域个数
### https://leetcode.cn/problems/number-of-provinces/
### 有 n 个城市，其中一些彼此相连，另一些没有相连，省份是一组直接或间接相连的城市，组内不含其他没有相连的城市。
### 给你一个 n*n 的矩阵 isConnected，其中 isConnected[i][j] = 1 表示第 i 个城市和第 j 个城市直接相连，而 isConnected[i][j] = 0 表示二者不直接相连。
### 返回矩阵中省份的数量。
def findCircleNum(isConnected):
  """
  :type isConnected: List[List[int]]
  :rtype: int
  """
  n = len(isConnected)
  search_arr = [-1]*n
  for i in range(n):
    for j in range(i):
      if isConnected[i][j] == 1:
        _fill_search_arr(search_arr, i, j)
  return len([i for i in search_arr if i == -1])

def _fill_search_arr(arr, i, j):
  x = _find_root(arr, i)
  y = _find_root(arr, j)
  if x == y:
    return
  arr[y] = x

def _find_root(arr, i):
  while arr[i] != -1:
    i = arr[i]
  return i

### 图的广度遍历(最短路径)
### https://leetcode.cn/problems/shortest-path-in-binary-matrix/
def shortestPathBinaryMatrix(self, grid):
  """
  :type grid: List[List[int]]
  :rtype: int
  """
  if grid[0][0] != 0:
    return -1

  visited = []
  for i in range(len(grid)):
    visited.append([0]*len(grid))

  queue = Queue()
  queue.put((0,0))
  queue.put((-1,-1))
  result = 1
  while True:
    x,y = queue.get()
    if x == -1 and y == -1:
      result = result+1
      if queue.empty():
        return -1
      queue.put((-1,-1))
      continue

    if x == len(grid)-1 and y == len(grid)-1:
      break

    if visited[x][y]:
      continue

    visited[x][y] = 1
    for cell in _next_cells(grid, x, y):
      if not visited[cell[0]][cell[1]]:
        queue.put(cell)

  return result

def _next_cells(grid, x, y):
  possibles = [
    (x+1,y),(x-1,y),
    (x,y+1),(x,y-1),
    (x+1,y+1),(x-1,y-1),
    (x+1,y-1),(x-1,y+1)
  ]

  result = []
  for a, b in possibles:
    if a < 0 or a >= len(grid) or b < 0 or b >= len(grid):
      continue
    if grid[a][b] == 0:
      result.append((a,b))

  return result

### 有向图的深度遍历
### https://leetcode.cn/problems/all-paths-from-source-to-target/
def allPathsSourceTarget(self, graph):
  """
  :type graph: List[List[int]]
  :rtype: List[List[int]]
  """
  return [p for p in _gen_paths(graph, 0, len(graph)-1)]

def _gen_paths(graph, start, end):
  for p in graph[start]:
    if p == end:
      yield [start, end]
      continue
    for path in gen_paths(graph, p, end):
      yield [start] + path

### 带重复数的全排列
def permuteUnique(self, nums):
  """
  :type nums: List[int]
  :rtype: List[List[int]]
  """
  mapping = {}
  for _, n in enumerate(nums):
    mapping[n] = mapping.get(n, 0) + 1
  return [x for x in _perm(mapping)]

def _perm(mapping):
  if len(mapping) == 1:
    yield [mapping.keys()[0]] * mapping.values()[0]
  else:
    for first in mapping.keys():
      m = deepcopy(mapping)
      if m[first] == 1:
        del m[first]
      else:
        m[first] = m[first] - 1
      for p in perm(m):
        yield [first] + p
