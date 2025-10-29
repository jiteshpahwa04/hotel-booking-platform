import express from "express";
import {
  createHotelHandler,
  getAllHotelsHanlder,
  getHotelByIdHanlder,
  softDeleteHotelHanlder,
} from "../../controllers/hotel.controller";
import { validateRequestBody } from "../../validator";
import { hotelSchema } from "../../validator/hotel.validator";

const hotelRouter = express.Router();

hotelRouter.post("/", validateRequestBody(hotelSchema), createHotelHandler);
hotelRouter.get("/:id", getHotelByIdHanlder);
hotelRouter.get("/", getAllHotelsHanlder);
hotelRouter.delete("/:id", softDeleteHotelHanlder);

export default hotelRouter;
