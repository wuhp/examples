import pytest
import alg

class TestBinarySerach:
  def test_success(self):
    pos = alg.binary_search([2,1,3], 0, 2, 1)
    assert pos == 1

  def test_fail(self):
    pos = alg.binary_search([2,1,3], 0, 2, 100)
    assert pos == -1

class TestExpression:
  def test_cal_post_expression_valid_001(self):
    result = alg.cal_post_expression(list("1253-*+32/-"))
    assert result == 4

  def test_cal_post_expression_valid_002(self):
    result = alg.cal_post_expression(["25", "10", "+"])
    assert result == 35

  def test_cal_post_expression_valid_003(self):
    result = alg.cal_post_expression(["25"])
    assert result == 25

  def test_expression_mid2post_001(self):
    result = alg.expression_mid2post(list("A+B*(C-D)-E/F"))
    assert "".join(result) == "ABCD-*+EF/-"

  def test_expression_mid2post_003(self):
    result = alg.expression_mid2post(["10", "-", "5"])
    assert result == ["10", "5", "-"]

  def test_expression_mid2post_003(self):
    result = alg.expression_mid2post(["10"])
    assert result == ["10"]

class TestPerm:
  def test_perm_generator_001(self):
    result = [
      [1,2,3],
      [1,3,2],
      [2,1,3],
      [2,3,1],
      [3,1,2],
      [3,2,1]
    ]
    n = 0
    for p in alg.perm_generator([1,2,3]):
      n = n+1
      assert p in result
    assert n == len(result)

  def test_perm_generator_002(self):
    result = [
      [1,2],
      [1,3],
      [2,1],
      [2,3],
      [3,1],
      [3,2]
    ]
    n = 0
    for p in alg.perm_generator([1,2,3],2):
      n = n+1
      assert p in result
    assert n == len(result)

  def test_perm_generator_003(self):
    result = [['a']]
    n = 0
    for p in alg.perm_generator(['a']):
      n = n+1
      assert p in result
    assert n == len(result)

class TestComb:
  def test_comb_generator_001(self):
    result = [
      set([1,2,3]),
      set([4,2,1]),
      set([3,4,1]),
      set([2,4,3]),
    ]
    n = 0
    for p in alg.comb_generator([1,2,3,4],3):
      n = n+1
      assert set(p) in result
    assert n == len(result)

  def test_comb_generator_002(self):
    result = [
      set([1,2,3,4])
    ]
    n = 0
    for p in alg.comb_generator([1,2,3,4]):
      n = n+1
      assert set(p) in result
    assert n == len(result)

  def test_comb_generator_003(self):
    result = [set(['a'])]
    n = 0
    for p in alg.comb_generator(['a']):
      n = n+1
      assert set(p) in result
    assert n == len(result)

@pytest.fixture
def tree_node():
  n14 = alg.TreeNode(14)
  n9 = alg.TreeNode(9)
  n6 = alg.TreeNode(6)
  n12 = alg.TreeNode(12, None, n14)
  n8 = alg.TreeNode(8, n6, n9)
  n15 = alg.TreeNode(15, n12, None)
  n5 = alg.TreeNode(5, None, n8)
  n10 = alg.TreeNode(10, n5, n15)
  return n10

class TestTree:
  def test_tree_recursive_mid_visit(self, tree_node):
    result = []
    alg.tree_recursive_mid_visit(tree_node, result)
    assert result == [5,6,8,9,10,12,14,15]

  def test_tree_mid_visit(self, tree_node):
    result = []
    alg.tree_mid_visit(tree_node, result)
    assert result == [5,6,8,9,10,12,14,15]

  def test_tree_front_visit(self, tree_node):
    result = []
    alg.tree_front_visit(tree_node, result)
    assert result == [10,5,8,6,9,15,12,14]

  def test_tree_rear_visit(self, tree_node):
    result = []
    alg.tree_rear_visit(tree_node, result)
    assert result == [6,9,8,5,14,12,15,10]

@pytest.fixture
def unsort_list():
  return [0,1,10,9,8,7,6,5,4,2,3]

class TestSort:
  def test_merge_sort_001(self, unsort_list):
    result = [i for i in range(11)]
    assert result == alg.merge_sort(unsort_list)

  def test_merge_sort_002(self):
    assert 0 == len(alg.merge_sort([]))

  def test_merge_sort_003(self):
    assert [1] == alg.merge_sort([1])

  def test_quick_sort_001(self, unsort_list):
    result = [i for i in range(11)]
    alg.quick_sort(unsort_list)
    assert result == unsort_list

  def test_quick_sort_002(self):
    arr = []
    alg.quick_sort(arr)
    assert 0 == len(arr)

  def test_quick_sort_003(self):
    arr = [5,4,4,5,5,5]
    alg.quick_sort(arr)
    assert [4,4,5,5,5,5] == arr

  def test_quick_sort_004(self):
    arr = [5,5,5,5]
    alg.quick_sort(arr)
    assert [5,5,5,5] == arr
