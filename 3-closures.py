def createCounter():
    count = 0

    def counter():
        nonlocal count  # count is not a local variable,
        # it refers to the count variable in the outer scope
        count += 1
        return count

    return counter


counter = createCounter()
print(counter())
print(counter())
