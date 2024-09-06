
export interface Auth_info {
    id: number;
    email: string;
    name: string;
    profile_id: number;
}

export interface Auth_token {
    access_token: string;
    refresh_token: string;
}