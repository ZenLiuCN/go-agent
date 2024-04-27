declare module "agent/secret" {
    /**
     * Do password validation
     * @param raw password
     * @param secret the hashed password
     */
    export function password(raw, secret: string): boolean

    /**
     * Do TOTP validation
     * @param code the one time passcode
     * @param def the url define or just the secret code (use default options)
     */
    export function totp(code, def: string): boolean

    /**
     * generate totp code
     * @param def url or just secret
     */
    export function totpCode(def: string): string

    /**
     * generate totp
     * @param opt setting
     */
    export function totpGen(opt: {
        // Name of the issuing Organization/Company.
        issuer: string
        // Name of the User's Account (eg, email address)
        accountName: string
        // Number of seconds a TOTP hash is valid for. Defaults to 30 seconds.
        period?: number
        // Size in size of the generated Secret. Defaults to 20 bytes.
        secretSize?: number
        // Digits to request. Defaults to 6.
        digits?: 6 | 8
        // Algorithm to use for HMAC. Defaults to SHA1. 0=SHA1,1=SHA256,2=SHA512,3=MD5
        algorithm?: 0 | 1 | 2 | 3
    }): string

    /**
     * BCrypt hash
     * @param raw raw password
     * @param cost the code between 4 ~ 31
     */
    export function bcrypt(raw: string, cost: number): string

    /**
     * Argon2 hash
     * @param raw raw password
     * @param opt options
     */
    export function argon2(raw: string, opt?: {
        memory?: number
        iterations?: number
        parallelism?: number
        saltLength?: number
        keyLength?: number
    }): string
}