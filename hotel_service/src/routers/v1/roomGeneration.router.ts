import express from "express";
import { validateRequestBody } from "../../validator";
import { generateRoomsHandler } from "../../controllers/roomGeneration.controller";
import { RoomGenerationJobSchema } from "../../dtos/roomGeneration.dto";

const roomGenerationRouter = express.Router();

roomGenerationRouter.post("/", validateRequestBody(RoomGenerationJobSchema), generateRoomsHandler);

export default roomGenerationRouter;
