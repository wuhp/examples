package handler

const (
    RequestBodyDecodeError = "invalid request body"
    RequestBodyError       = "invalid request body field"
    DuplicateResource      = "duplicate resource"

    TokenExpireTime = 3600000

    ManagedAuth     = "Managed"
)

func encodeUserToken(username string) string {
    // TBD
    return ""
}

func decodeUserToken(key string) (string, bool) {
    // TBD
    // username, expiered
    return "", false
}

func hashPassword(password string) string {
    // TBD
    return ""
}
