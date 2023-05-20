// used for token management
export function setCookie(tokenKey, token) {
    return localStorage.setItem(tokenKey, token)
}
export function getCookie(tokenKey) {
    return localStorage.getItem(tokenKey)
}
export function removeCookie(tokenKey) {
    return localStorage.removeItem(tokenKey)
}