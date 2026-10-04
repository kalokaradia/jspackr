import { config } from "./config/app";
import { UserService } from "./services/user-service";
import { calculateStats } from "./services/stats-service";
import { info, warn } from "./utils/logger";

async function main(): Promise<void> {
  info(`Starting ${config.name} v${config.version}`);
  info(`Environment: ${config.environment}`);

  const userService = new UserService();

  const users = userService.getActiveUsers();

  console.log("\nActive users:");

  for (const user of users) {
    console.log({
      ...user.toJSON(),
      normalizedScore: userService.calculateScore(user),
    });
  }

  const stats = calculateStats(users);

  console.log("\nStatistics:");
  console.table(stats);

  const targetUser = await userService.findById(1);

  if (targetUser) {
    console.log("\nFound user:");

    console.log(targetUser.toJSON());
  } else {
    warn("User not found");
  }
}

main().catch((error: unknown) => {
  console.error("Application failed:", error);
});