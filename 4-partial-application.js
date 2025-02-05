const createGreeting = (greeting) => {
  return (name) => {
    return `${greeting} ${name}`;
  };
};

const firstGreeting = createGreeting("Well, hello there ");
const secondGreeting = createGreeting("Hola ");
console.log(firstGreeting("Remi"));
console.log(secondGreeting("Ana"));