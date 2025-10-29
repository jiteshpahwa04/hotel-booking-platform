import { NextFunction, Request, Response } from "express";
import { InternalServerError } from "../utils/errors/app.error";

export const pingHandler = async(req: Request, res: Response, next: NextFunction): Promise<void> => {
  console.log("request body: ", req.body);
  console.log("query params: ", req.query);
  console.log("request params: ", req.params);

  try{
    // this throw only works with async functions
  }catch(err){
    throw new InternalServerError("Something went wrong!");
  }
  let err = "2";
  if(err=="1"){
    next(err); // this is how you handle errors in async calls
  }

  res.status(200).json({
    success: true,
    message: "PONG"
  })
};