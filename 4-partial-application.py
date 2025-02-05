def createGreeting(greeting: str):
    def greetingFn(name: str) -> str:
        return f"{greeting} {name}"

    return greetingFn


firstGreeting = createGreeting("Well, hello there ")
secondGreeting = createGreeting("Hola ")
print(firstGreeting("Remi"))
print(secondGreeting("Ana"))
