import { numbers, sum } from "./utils/numbers.js";
import { formatNumber } from "./utils/format.js";

const total = sum(numbers);

console.log("Numbers:", numbers);
console.log("Total:", formatNumber(total));
