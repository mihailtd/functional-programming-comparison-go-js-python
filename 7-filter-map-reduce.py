from functools import reduce

# Map: Square each number
numbers = [1, 2, 3, 4, 5]
squares = list(map(lambda n: n * n, numbers))
print("Squares:", squares)

# Filter: Get even numbers
evens = list(filter(lambda n: n % 2 == 0, numbers))
print("Evens:", evens)

# Reduce: Sum all numbers
sum_of_numbers = reduce(lambda acc, n: acc + n, numbers, 0)  # 0 is the initial value
print("Sum:", sum_of_numbers)
