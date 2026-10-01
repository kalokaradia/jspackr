import { Cache } from "./lib/cache.js";
import { fetchData } from "./lib/api.js";
import { parseUsers } from "./utils/parser.js";
import { formatUser } from "./utils/formatter.js";

const cache = new Cache();

const users = [
  { id: 1, name: " Alice ", active: true },
  { id: 2, name: "Bob", active: false },
  { id: 3, name: " Charlie ", active: true }
];

const parsed = parseUsers(users);

for (const user of parsed) {
  console.log(formatUser(user));
}

cache.set("users", parsed);

console.log("Cached:", cache.has("users"));

const data = await fetchData();

console.log("API:", data);