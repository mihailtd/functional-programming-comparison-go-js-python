from functools import reduce


def compose(*funcs):
    return lambda x: reduce(lambda acc, f: f(acc), reversed(funcs), x)


def add_one(x):
    return x + 1


def square(x):
    return x * x


def multiply_by_two(x):
    return x * 2


composed = compose(square, multiply_by_two, add_one)  # Order is correct
result = composed(5)  # ((5 + 1) * 2)^2 = 144
print("Composed result:", result)
