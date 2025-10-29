import express from "express";
import { validateRequestBody } from "../../validator";
import { generateRoomsHandler } from "../../controllers/roomGeneration.controller";
import { RoomGenerationJob } from "../../dtos/roomGeneration.dto";

const hotelRouter = express.Router();

hotelRouter.post("/", validateRequestBody(RoomGenerationJob), generateRoomsHandler);

export default hotelRouter;
