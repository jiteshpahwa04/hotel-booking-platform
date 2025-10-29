import { NextFunction, Request, Response } from "express";
import { StatusCodes } from "http-status-codes";
import { addRoomGenerationJobToQueue } from "../producers/roomGeneration.producer";

export async function generateRoomsHandler(req: Request, res: Response, next: NextFunction) {
    await addRoomGenerationJobToQueue(req.body);
    res.status(StatusCodes.OK).json({
      message: "Rooms generation request queued successfully",
      data: null,
      success: true,
    });
}