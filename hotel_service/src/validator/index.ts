import { NextFunction, Request, Response } from "express";
import { AnyZodObject } from "zod";

export const validateRequestBody = (schema: AnyZodObject) => {
    return async (req: Request, res: Response, next: NextFunction)=>{
        try {
            await schema.parseAsync(req.body);
            console.log("Request body is valid");
            next();
        } catch (error) {
            res.status(400).json({
                success: false,
                message: "Invalid json body",
                error: error
            })
        }
    }
}

export const validateQueryParams = (schema: AnyZodObject) => {
    return async (req: Request, res: Response, next: NextFunction)=>{
        try {
            await schema.parseAsync(req.query);
            console.log("Query params is valid");
            next();
        } catch (error) {
            res.status(400).json({
                success: false,
                message: "Invalid json query params",
                error: error
            })
        }
    }
}