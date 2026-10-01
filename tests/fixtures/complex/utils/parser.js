export function parseUsers(users) {
  return users
    .filter((user) => user.active)
    .map((user) => ({
      ...user,
      name: user.name.trim(),
    }));
}
