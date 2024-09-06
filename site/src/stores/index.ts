

// call everytime web mount
export const init_store = async (): Promise<void> => {
    console.log("init store")
}

// call everytime web mount if logged in 
export const init_store_logged_in = async (): Promise<void> => {
    console.log("init store logged in")
}