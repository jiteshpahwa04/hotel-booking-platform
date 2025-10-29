import dotenv from 'dotenv';

type ServerConfigType = {
    PORT: number
    REDIS_HOST: string
    REDIS_PORT: number
}

type DbConfigType = {
    DB_HOST: string,
    DB_USER: string,
    DB_PASSWORD: string,
    DB_NAME: string
}

function loadEnv(){
    dotenv.config();
    console.log("Environment variables loaded");
}

loadEnv();

export const serverConfig : ServerConfigType = {
    PORT: Number(process.env.PORT) || 3001,
    REDIS_HOST: process.env.REDIS_HOST || 'localhost',
    REDIS_PORT: Number(process.env.REDIS_PORT) || 6379
};

export const dbConfig : DbConfigType = {
    DB_HOST: process.env.DB_HOST || 'localhost',
    DB_USER: process.env.DB_USER || 'root',
    DB_PASSWORD: process.env.DB_PASSWORD || 'my-password',
    DB_NAME: process.env.DB_NAME || 'airbnb'
}