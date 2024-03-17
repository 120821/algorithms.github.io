# 归并排序
# 原始数组分成较小的数组，直到每个小数组只有一个位置，接着将小数组合并成较大的数组，最终合并为排序好的数组。
# 复杂度：
# 时间复杂度：O(nlogn)O(nlogn)。
# 空间复杂度：O(n)O(n)。

def merge_sort(arr)
  return arr if arr.size <= 1
  mid = arr.size / 2
  left = merge_sort(arr[0...mid])
  right = merge_sort(arr[mid..])
  merge(left, right)
end

def merge(left, right)
  sorted = []
  until left.empty? || right.empty?
    sorted << (left.first <= right.first ? left.shift : right.shift)
  end
  sorted + left + right
end

arr = [4, 2, 1, 22, 21, 5, 29, 28, 35]
sorted_arr = merge_sort(arr)
puts sorted_arr
