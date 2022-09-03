# coding:utf-8

import heapq

from Queue import LifoQueue
from copy import deepcopy

### 链表排序 - 归并思路
def sort_linked_list(head):
    if not head.next:
        return head

    first = second = head
    while True:
        if not second.next:
            break
        second = second.next
        if not second.next:
            break
        second = second.next
        first = first.next
    tmp = first.next
    first.next = None
    first = tmp
    h1 = sort_linked_list(head)
    h2 = sort_linked_list(first)
    return merge_sorted_linked_list(h1, h2)

def merge_sorted_linked_list(h1, h2):
    if not h1:
        return h2
    if not h2:
        return h1

    head = rear = None
    while True:
        if not h1:
            rear.next = h2
            break
        if not h2:
            rear.next = h1
            break
        if h1.val > h2.val:
            if not head:
                head = rear = h2
            else:
                rear.next = h2
                rear = rear.next
            h2 = h2.next
            rear.next = None
            continue
        if not head:
            head = rear = h1
        else:
            rear.next = h1
            rear = rear.next
        h1 = h1.next
        rear.next = None
    return head

### 拓扑排序
### https://leetcode.cn/problems/course-schedule-ii/
def find_order(self, numCourses, prerequisites):
    """
    :type numCourses: int
    :type prerequisites: List[List[int]]
    :rtype: List[int]
    """
    courses = set(range(numCourses))
    result = []
    while len(courses) > 0:
        has_dependency_set = set([p[0] for p in prerequisites])
        outputs = courses - has_dependency_set
        if len(outputs) == 0:
            return []
        result.extend(list(outputs))
        courses = has_dependency_set
        prerequisites = [[a,b] for a, b in prerequisites if a not in outputs and b not in outputs]
    return result

### 动态规划（背包问题）
### https://leetcode.cn/problems/coin-change/
def coin_change(self, coins, amount):
    """
    :type coins: List[int]
    :type amount: int
    :rtype: int
    """
    result = [0]*(amount+1)
    for val in range(1, amount+1):
        num = -1
        for coin in coins:
            if val < coin:
                continue
            if val == coin:
                num = 1
                break
            if result[val-coin] == -1:
                continue
            if num < 0:
                num = result[val-coin] + 1
                continue
            if result[val-coin] + 1 > num:
                continue
            num = result[val-coin] + 1
        result[val] = num
    return result[-1]

### 双指针
### https://leetcode.cn/problems/3sum-closest/
def three_sum_closest(nums, target):
    """
    :type nums: List[int]
    :type target: int
    :rtype: int
    """
    nums.sort()
    result = nums[0] + nums[1] + nums[2]
    for i in range(len(nums)-2):
        sub_target = target - nums[i]
        left = i+1
        right = len(nums)-1
        sub_result = nums[left]+nums[right]
        while left != right:
            if nums[left]+nums[right] == sub_target:
                return target
            if abs(nums[left]+nums[right]-sub_target) < abs(sub_result-sub_target):
                sub_result = nums[left]+nums[right]
            if nums[left]+nums[right] > sub_target:
                right = right-1
                continue
            left = left+1
        if abs(sub_result+nums[i]-target) < abs(result-target):
            result = sub_result+nums[i]
    return result
