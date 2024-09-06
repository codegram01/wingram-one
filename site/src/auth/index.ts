import { keyLocalToken } from "@/key/local"
import { auth_info_api } from "@/services/auth"
import { init_store_logged_in } from "@/stores"
import { getToken, setAuthInfo, setToken } from "@/stores/auth"
import type { Auth_token } from "@/types/auth"

import { computed } from "vue"

// init token from local
export const getTokenLocal = (): void => {
    const tokenLocal = localStorage.getItem(keyLocalToken)
    if(tokenLocal){
        setToken(tokenLocal)
    }
}

// Call when user call Login
export const setAuthLogin = async (authTkn: Auth_token):Promise<void> => {
    console.log(authTkn)
    
    setTokenAuth(authTkn.access_token)

    // TODO: handle refresh token

    await get_auth_info()
}

// save token and token local
const setTokenAuth = (tkn: string): void => {
    const token = "Bearer " + tkn
    setToken(token)
    localStorage.setItem(keyLocalToken, token)
}

// init auth when web mount
export const init_auth = async (): Promise<void> => {
    console.log("--> init auth")
    getTokenLocal()

    if(getToken.value) {
        await get_auth_info();
    }
}

// wrap auth_info_api
// it check token is right
// save auth_info - call init store logged in
const get_auth_info = async (): Promise<void> => {
    console.log("get auth info")
    await auth_info_api().then(res => {
        console.log("have auth info return")
        setAuthInfo(res.data)

        init_store_logged_in()
    })
}