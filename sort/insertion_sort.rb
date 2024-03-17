# 插入排序
# 对于没有排序的元素，在已经排序的序列里从后向前扫描，然后在相应的位置插入。
# 复杂度：
# 时间复杂度：最坏情况下 O(n2)O(n2)，最好情况下 O(n)O(n)。
# 空间复杂度：O(1)O(1)。
def insertion_sort arr
  n = arr.length
  n.times do |i|
    key = arr[i]
    j = i-1
    while j >=0 and arr[j] > key
      arr[j+1] = arr[j]
      j-=1
    end
    arr[j+1] = key
  end
  arr
end

arr = [4, 2, 1, 5, 29, 25, 22, 29]
insertion_sort arr
puts arr
