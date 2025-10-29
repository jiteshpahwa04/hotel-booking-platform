import { NextFunction, Request, Response } from "express";
import {v4 as uuidV4} from "uuid";
import { asyncLocalStorage } from "../utils/helpers/request.helper";

export const attachCorrelationMiddleWare = async(req: Request, res: Response, next: NextFunction)=>{
    const correlationId = uuidV4();

    req.headers['x-Correlation-ID'] = correlationId;

    asyncLocalStorage.run({correlationId: correlationId}, ()=>{
        next();
    });
}