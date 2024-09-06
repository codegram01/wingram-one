import type { Auth_info } from "@/types/auth";
import { computed, ref } from "vue";

const auth_info = ref<Auth_info | undefined>()
export const setAuthInfo = (a: Auth_info) => {
    auth_info.value = a
}

export const getAuthInfo = computed(()=>{
    return auth_info.value;
})

export const isLoggedIn = computed(()=> {
    if(auth_info.value){
        return true
    }
    return false
})


const token = ref<string | undefined>()
export const setToken = (tkn: string) => {
    token.value = tkn
}

export const getToken = computed(()=>{
    return token.value;
})