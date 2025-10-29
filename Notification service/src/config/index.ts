// This file contains all the basic configuration logic for the app server to work
import dotenv from 'dotenv';

type ServerConfig = {
    PORT: number,
    REDIS_HOST: string,
    REDIS_PORT: number,
    MAIL_PASSWORD: string,
    MAIL_USER: string
}

function loadEnv() {
    dotenv.config();
    console.log(`Environment variables loaded`);
}

loadEnv();

export const serverConfig: ServerConfig = {
    PORT: Number(process.env.PORT) || 3001,
    REDIS_HOST: process.env.REDIS_HOST || 'localhost',
    REDIS_PORT: Number(process.env.REDIS_PORT) || 6379,
    MAIL_PASSWORD: process.env.MAIL_PASS || '',
    MAIL_USER: process.env.MAIL_USER || ''
};