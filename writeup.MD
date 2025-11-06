# Write Up Questions

## Q: How long did you spend working on the problem? What did you find to be the most difficult part?
**A:** I spent about 3 hours getting a solution running that passed the simulator tests. The hardest part was definitely learning Go as I went. The night before starting, I completed the [Go Tour tutorial](https://go.dev/tour/list) to become familiar with the language, and I relied heavily on AI assistance to map my understanding from JavaScript/Express/NestJS services to Go concepts, patterns, conventions, and syntax.

## Q: How would you modify your data model or code to account for more kinds of metrics?
**A:** I've already isolated a model file for the `Device`, with a sub-struct for `UploadStats`, for example. I would continue with that pattern for additional simple metrics, or consider more fully-featured models if complexity or data volume increases. This approach keeps the code modular, maintainable, and easy to extend.

## Q: Discuss your solution’s runtime complexity
**A:** While the solution works well for the purposes of this assignment, storing all heartbeat and stat data in memory would become unmanageable as the number of devices and events grows. At a high level, storing the data in a database would help with scalability. Time complexity for the GET endpoint (calculating uptime and average upload stats) is linear in the number of events per device, which is acceptable for small datasets. For larger deployments, maintaining in-memory aggregates or implementing a data retention/purging strategy would help manage memory usage and improve performance.
