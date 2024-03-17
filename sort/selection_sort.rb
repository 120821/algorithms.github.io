# 选择排序
# 通过n次查询，每次查询剩余部分的最小值，并且把最小值放在未排序的起始位置。
# 复杂度：
# 时间复杂度：总是 O(n2)O(n2)。
# 空间复杂度：O(1)O(1)。

def selection_sort arr
  length = arr.size
  length.times do |i|
    min_index = i
    (i+1...length).each do |j|
      min_index = j if arr[j] < arr[min_index]
    end
    arr[i], arr[min_index] = arr[min_index], arr[i] if min_index != i
  end
  arr
end
arr = [4, 5, 2, 1, 11, 29, 22, 34, 25]
puts selection_sort(arr)
