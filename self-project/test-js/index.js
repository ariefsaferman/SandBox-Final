const myNumbers = [4, 1, -20, -7, 5, 9, -6];

console.log(removeNeg(myNumbers, (x) => x < 0));

function removeNeg(arr, callback) {
  const arr2 = [];

  arr.forEach((element) => {
    if (callback(element)) {
      arr2.push(element);
    }
  });

  return arr2;
}
