# 快速排序
# 选择一个值作为标准，然后进行一次遍历，把元素分为2部分，一部分是大于标准值，部分是小于标准值
# 然后重复这个操作，也就是把这两部分的元素再次进行基于新的基准值进行排序(递归)
# 直到整个数据都成为有序序列

# 复杂度：
# 时间复杂度：平均 O(nlog⁡n)O(nlogn)，最坏 O(n2)O(n2)（极端情况，如已经排序的数组）。
# 空间复杂度：O(log⁡n)O(logn)（递归调用栈的大小）。


def quick_sort(arr)
  return arr if arr.length <= 1
  pivot = arr.delete_at(arr.length / 2)
  left, right = arr.partition(&pivot.method(:>))
  quick_sort(left) + [pivot] + quick_sort(right)
end

array = [1, 3, 5, 3, 9, 10]
puts quick_sort(array)
