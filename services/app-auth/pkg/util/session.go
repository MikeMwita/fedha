package util

//
//func CreateSessionCookie(cfg *config.Config, sess *entity.Session) *http.Cookie {
//	value := url.QueryEscape(sess.SessionID)
//	cookie := &http.Cookie{
//		Name:     cfg.Session.Name,
//		Value:    value,
//		Path:     "/",
//		HttpOnly: true,
//		//Secure:   cfg.Session.Secure,
//		SameSite: http.SameSiteStrictMode,
//	}
//	if cfg.Session.Expire > 0 {
//		expire := time.Now().Add(time.Duration(cfg.Session.Expire) * time.Second)
//		cookie.Expires = expire
//	}
//	return cookie
//}
