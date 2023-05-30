export function clearStorage() {
    localStorage.clear()
}

export function setUUID(uuid) {
    localStorage.setItem("uuid", uuid)
}

export function getUUID() {
    return localStorage.getItem("uuid")
}
