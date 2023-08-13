import { WebAuth } from "auth0-js"

export const auth0 = new WebAuth({
    clientID:"8BzcKDUjJFpMAGe8KPgkUor7pHRmkI2x",
    domain:"dev-vrsblas6ves78ir5.us.auth0.com",
    redirectUri:"http://localhost:5173"
})