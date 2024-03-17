# 冒泡排序
# 比较相邻的2个元素，直到排序完成
# 复杂度
# 时间复杂度：最坏情况下 O(n2)O(n2)，最好情况下 O(n)O(n)（已经排序的数组）。
# 空间复杂度：O(1)O(1)。

def bubble_sort arr
  len = arr.size
  for i in 0..len-1
    for j in 1..len-2
      if arr[j] > arr[j+1]
        arr[j], arr[j+1] = arr[j+1], arr[j]
      end
    end
  end
  return arr
end

arr = [5, 4, 7, 5, 10, 28, 33,2, 11]
puts bubble_sort(arr)
