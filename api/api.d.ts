declare module "agent/api" {
    // @ts-ignore
    import * as model from "agent/model";

    interface Cache<K,V> {

    }

    interface Action {
        readonly user?: model.User
    }
}