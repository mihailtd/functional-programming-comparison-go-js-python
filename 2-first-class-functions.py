def applyOperation(operation, a: int, b: int) -> int:
    return operation(a, b)


def add(a: int, b: int) -> int:
    return a + b


def subtract(a: int, b: int) -> int:
    return a - b


result = applyOperation(add, 10, 5)
print(result)  # Output: 15

result = applyOperation(subtract, 10, 5)
print(result)  # Output: 5
