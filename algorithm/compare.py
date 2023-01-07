# coding:utf-8

# 自定义比较
# https://leetcode.cn/problems/sort-integers-by-the-number-of-1-bits/description/?envType=study-plan&id=programming-skills-beginner&plan=programming-skills&plan_progress=chlh3ac

def num(x):
    return len([c for c in bin(x) if c == '1'])

def compare(x, y):
    if num(x) < num(y):
        return -1
    if num(x) > num(y):
        return 1
    if x < y:
        return -1
    if x > y:
        return 1
    return 0

class Solution(object):
    def sortByBits(self, arr):
        """
        :type arr: List[int]
        :rtype: List[int]
        """
        arr.sort(key=functools.cmp_to_key(compare))
        return arr
