import express from 'express';
import { pingHandler } from "../../controllers/ping.controller";
import { validateRequestBody } from '../../validator';
import { pingSchema } from '../../validator/ping.validator';

const pingRouter = express.Router();

pingRouter.get("/", validateRequestBody(pingSchema), pingHandler);
// Also possible: pingRouter.get("/ping", [middleware1, middleware2], pingHandler);

pingRouter.get("/health", (req, res)=>{
  res.status(200).send("OK");
})

export default pingRouter;