import { CreationAttributes, Model, ModelStatic, WhereOptions } from "sequelize";
import { NotFoundError } from "../utils/errors/app.error";

abstract class BaseRepository<T extends Model> {

    protected model: ModelStatic<T>;

    constructor(model: ModelStatic<T>) {
        this.model = model;
    }

    async create(data: CreationAttributes<T>): Promise<T> {
        return await this.model.create(data);
    }

    async findById(id: number): Promise<T | null> {
        const record = await this.model.findByPk(id);
        if(!record) {
            return null;
        }
        return record;
    }

    async findAll(): Promise<T[]> {
        const records = await this.model.findAll();
        if(!records) {
            return [];
        }
        return records;
    }

    async update(whereOptions: WhereOptions<T>, data: Partial<T>): Promise<[affectedCount: number]> {
        return await this.model.update(data, { where: whereOptions });
    }

    async deleteById(whereOptions: WhereOptions<T>): Promise<number> {
        const record = await this.model.destroy({ where: {
            ...whereOptions
        } });
        if (!record) {
            throw new NotFoundError(`Record not found for deletion with options: ${JSON.stringify(whereOptions)}`);
        }
        return record;
    }

}

export default BaseRepository;