
export const process_error = (e: any): void => {
    if(typeof e === "string") {
        alert(e)
    }
}