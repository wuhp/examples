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
