const fs = require("fs");

// Read the input from stdin synchronously
const input = fs.readFileSync(0, "utf8"); // 0 represents stdin

// Split the input based on new lines
const lines = input.trim().split("\n");
const arraySize = parseInt(lines[0], 10);
const arr = lines[1].split(" ").map(Number);

function elementInMiddle(arr) {
  // ... function logic remains the same
  const N = arr.length;
  let response = -1;

  for (let i = 1; i < N - 1; i++) {
    let leftLess = true;
    let rightMore = true;

    for (let j = 0; j < i; j++) {
      if (arr[j] > arr[i]) {
        leftLess = false;
        break;
      }
    }

    for (let k = i + 1; k < N; k++) {
      if (arr[i] > arr[k]) {
        rightMore = false;
        break;
      }
    }

    if (leftLess && rightMore) {
      response = arr[i];
      break;
    }
  }

  return response;
}

// Check if the array length matches the provided size
if (arr.length !== arraySize) {
  console.error(
    "The number of elements does not match the specified array size."
  );
  process.exit(1);
}

// Call the function and output the result
console.log(elementInMiddle(arr));
