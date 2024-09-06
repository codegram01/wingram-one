import { BASE_API } from "@/key/key";
import { getToken } from "@/stores/auth";
import type { Res } from "@/types/api";
import type { METHOD_HTTP } from "@/types/http";

export const gfetch = async (method: METHOD_HTTP, url: string, data: any | undefined): Promise<Res> => {
    const urlApi = BASE_API + url;
    let body;
    if(data){
        body = JSON.stringify(data)
    }

    return await fetch(urlApi, {
        method: method,
        mode: "cors",
        headers: {
            "Content-Type": "application/json",
            "Authorization": getToken.value ? getToken.value : ""
        },
        body: body,
    }).then(async res => {
        const data: Res = await res.json()

        if(res.ok){
            return data;   
        }else {
            throw data;
        }
    })
}