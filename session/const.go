package session

import "time"

const defaultSessionKeyName = "y"
const defaultDomain = ""
const defaultExpires = 5 * time.Hour
const defaultDemoExpires = 5 * time.Minute // 试玩登录默认过期时间


const defaultGCLifetime = 60 * time.Second
const defaultSecure = false
const defaultSessionIDInURLQuery = false
const defaultSessionIDInHTTPHeader = false
const defaultCookieLen uint32 = 32

const expirationAttributeKey = "_sid_"
