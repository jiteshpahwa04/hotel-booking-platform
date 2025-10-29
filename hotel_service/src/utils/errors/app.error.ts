// type keyword does not support object oriented implementations
// interface keyword acts as a contract
export interface AppError extends Error {
    statusCode: number;
}

export class InternalServerError implements AppError {
    statusCode: number;
    message: string;
    name: string;
    constructor(messageStr: string){
        this.statusCode = 500;
        this.message = messageStr;
        this.name = "InternalServerError"
    }
}

export class BadRequestError implements AppError {
    statusCode: number;
    message: string;
    name: string;
    constructor(messageStr: string){
        this.statusCode = 400;
        this.message = messageStr;
        this.name = "InternalServerError"
    }
}

export class NotFoundError implements AppError {
    statusCode: number;
    message: string;
    name: string;
    constructor(messageStr: string){
        this.statusCode = 404;
        this.message = messageStr;
        this.name = "InternalServerError"
    }
}

export class ConflictError implements AppError {
    statusCode: number;
    message: string;
    name: string;
    constructor(messageStr: string){
        this.statusCode = 500;
        this.message = messageStr;
        this.name = "InternalServerError"
    }
}

export class UnauthenticatedError implements AppError {
    statusCode: number;
    message: string;
    name: string;
    constructor(messageStr: string){
        this.statusCode = 401;
        this.message = messageStr;
        this.name = "InternalServerError"
    }
}

export class ForbiddenError implements AppError {
    statusCode: number;
    message: string;
    name: string;
    constructor(messageStr: string){
        this.statusCode = 403;
        this.message = messageStr;
        this.name = "InternalServerError"
    }
}