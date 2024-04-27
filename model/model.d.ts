declare module "agent/model" {
    //@ts-ignore
    import {Time} from "go/time"
    //@ts-ignore
    import {NullString, SQLX} from "go/sqlx"
    //@ts-ignore
    import {Context} from "go/context"

    export interface Model {
        // int64
        id: number
        // int32
        version: number
        removed: boolean
        createAt: Time
        // int64
        createBy: number
        modifiedAt: Time
        // int64
        modifiedBy: number
    }

    export interface Entity {
        save(ctx: Context): boolean               //update record
        saveBy(ctx: Context, actor: number): boolean   //update record
        deleteBy(ctx: Context, actor: number): boolean //soft delete record
        delete(ctx: Context): boolean             //soft delete record
        drop(ctx: Context): boolean               //drop record permanently
        refresh(ctx: Context): boolean

        close(ctx: Context): boolean
    }

    export interface User extends Model {
        name: string
        nick: string
        secret: NullString
        totp: NullString


        checkTOTP(code: string): boolean

        checkSecret(raw: string): boolean
    }

    export interface UserEntity extends User, Entity {
        getName(): string

        getNick(): string

        getSecret(): NullString

        getTotp(): NullString

        setName(v: string): boolean

        setNick(v: string): boolean

        setSecret(v: NullString): boolean

        setTotp(v: NullString): boolean
    }

    export interface UserStore {
        byName(name: string): UserEntity | undefined

        byId(id: number): UserEntity | undefined
    }

    export function newUserStore(table: string, db: SQLX): UserStore
}