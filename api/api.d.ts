declare module "agent/api" {
    // @ts-ignore
    import * as model from "agent/model";
    // @ts-ignore
    import * as units from "agent/units";
    // @ts-ignore
    import {Request, ResponseWriter} from "go/http";

    export interface Action extends Request, ResponseWriter {
        request: Request
        responseWriter: ResponseWriter
        //current request user
        readonly user?: model.User
        //cache that exists cross actions
        readonly cache?: units.Cache<string, string>

        searchQuery(key: string): string

        jsonBody(): any

        binaryBody(): Uint8Array | undefined

        textBody(): string | undefined

        //response with status 200
        ok()

        //response with status code
        status(code: number)

        //response with json body
        json(v: Record<string, any> | Array<any> | ArrayBuffer | Uint8Array | string)

        //response with json body
        binary(contentType: string, v: Uint8Array | ArrayBuffer)

        /**
         *
         * @param url to redirect to
         * @param status 3xx,eg: 307
         */
        redirect(url: string, status: number)

        /**
         * dump request as string
         */
        dump(): string
    }
}
