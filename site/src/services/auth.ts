import { gfetch } from "@/modules/api";
import type { Auth_info, Auth_token } from "@/types/auth";
import { process_error } from ".";
import type { Res } from "@/types/api";

export interface Register_req{
    email: string;
    password: string;
    name: string;
}

interface Auth_info_res extends Res {
    data: Auth_info
};

export const register_api = async (data: Register_req): Promise<Auth_info_res> => {
    try {
        return await gfetch("POST", "/accounts/register", data)
    } catch (error) {
        process_error(error)
        throw error
    }
}

export interface Login_req {
    email: string;
    password: string;
}

interface Login_res extends Res {
    data: Auth_token
}

export const login_api = async (data: Login_req): Promise<Login_res> => {
    try {
        return await gfetch("POST", "/accounts/login", data)
    } catch (error) {
        process_error(error)
        throw error
    }
}

export const auth_info_api = async (): Promise<Auth_info_res> => {
    try {
        return await gfetch("GET", "/accounts/info", null)
    } catch (error) {
        process_error(error)
        throw error
    }
}