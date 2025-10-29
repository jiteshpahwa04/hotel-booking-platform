import express from "express";
import { Express } from "express";
import {serverConfig} from "./config";
import v1Router from "./routers/v1/index.router";
import v2Router from "./routers/v2/index.router";
import { genericErrorHandler } from "./middlewares/error.middleware";
import logger from "./config/logger.config";
import { attachCorrelationMiddleWare } from "./middlewares/correlation.middleware";
import sequelize from "./db/models/sequelize";
import { setupRoomGenerationWorker } from "./processors/roomGeneration.processor";

const app: Express = express();

const PORT = serverConfig.PORT;

app.use(express.json());

app.use(attachCorrelationMiddleWare);
// Registering all the routers and their corresponding routes with our app server object
app.use('/api/v1', v1Router);
app.use('/api/v2', v2Router);

// Error handler middleware
app.use(genericErrorHandler);

app.listen(PORT, async ()=>{
    logger.info(`Server is running on port: ${PORT}`);
    await sequelize.authenticate();
    logger.info("Database connected successfully!");
    setupRoomGenerationWorker();
    logger.info("Room generation worker setup completed!");
});