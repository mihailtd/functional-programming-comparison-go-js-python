def curry(fn):
    def curried(*args):
        if len(args) >= fn.__code__.co_argcount:
            return fn(*args)
        else:
            return lambda *args2: curried(*(args + args2))

    return curried


def add(a: int, b: int, c: int) -> int:
    return a + b + c


result = curry(add)(1)(2)(3)
print(result)  # Output: 6
