arr = []

arr[0] = 8
arr[1] = 5
arr[2] = 1
arr[3] = 4
arr[4] = 10
arr[5] = 6
arr[6] = 8
arr[7] = 9

n = 8
i = 0
temp = 0
j = 0
swapped = 0

while i < n-1
  j = 0
  swapped = 0
  while j < n - i - 1
    if arr[j] > arr[j+1]
      temp = arr[j]
      arr[j] = arr[j+1]
      arr[j+1] = temp
      swapped = swapped + 1
    end
    j = j + 1
  end

  if swapped == 0
    return arr[0]
  end
  i = i+1
end

return arr[1]
