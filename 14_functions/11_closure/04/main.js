wrapper = () => {
  let x = 0;
  return () => {
    x++;
    return x;
  }
}

const increment = wrapper();
console.log(increment());
console.log(increment());
