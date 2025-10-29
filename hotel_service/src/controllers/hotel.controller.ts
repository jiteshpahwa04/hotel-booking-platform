import { NextFunction, Request, Response } from "express";
import {
  createHotelService,
  getAllHotelsService,
  getHotelByIdService,
  softDeleteHotelService,
} from "../services/hotel.service";
import { StatusCodes } from "http-status-codes";

export const createHotelHandler = async (
  req: Request,
  res: Response,
  next: NextFunction
) => {
  const hotelResponse = await createHotelService(req.body);
  res.status(StatusCodes.CREATED).json({
    message: "Hotel created successfully",
    data: hotelResponse,
    success: true,
  });
};

export async function getHotelByIdHanlder(
  req: Request,
  res: Response,
  next: NextFunction
) {
  const hotelResponse = await getHotelByIdService(Number(req.params.id));

  res.status(StatusCodes.OK).json({
    message: "Hotel found successfully",
    data: hotelResponse,
    success: true,
  });
}

export async function getAllHotelsHanlder(
  req: Request,
  res: Response,
  next: NextFunction
) {
  const hotelResponse = await getAllHotelsService();

  res.status(StatusCodes.OK).json({
    message: "Hotels found successfully",
    data: hotelResponse,
    success: true,
  });
}

export async function softDeleteHotelHanlder(
  req: Request,
  res: Response,
  next: NextFunction
) {
  const hotelResponse = await softDeleteHotelService(Number(req.params.id));

  res.status(StatusCodes.OK).json({
    message: "Hotels deleted successfully",
    data: hotelResponse,
    success: true,
  });
}