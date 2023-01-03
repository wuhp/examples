# coding:utf-8

# 迭代器
# https://leetcode.cn/problems/reshape-the-matrix/description/?envType=study-plan&id=programming-skills-beginner&plan=programming-skills&plan_progress=chlh3ac

def generator(mat, m, n):
    for i in range(0, m):
        for j in range(0, n):
            yield mat[i][j]

class Solution(object):
    def matrixReshape(self, mat, r, c):
        """
        :type mat: List[List[int]]
        :type r: int
        :type c: int
        :rtype: List[List[int]]
        """
        m = len(mat)
        n = len(mat[0])
        if m*n != r*c:
            return mat
        result = []
        it = iter(generator(mat, m ,n))
        for i in range(0, r):
            line = []
            for j in range(0, c):
                line.append(next(it))
            result.append(line)
        return result
