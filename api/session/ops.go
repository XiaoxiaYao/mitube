package session

import (
	"mitube/api/dbops"
	"time"
	"sync"
)

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func LoadSessionsFromDB() {
	r, err :=dbops.ListSessions()
	if err != nil {
		return
	}
	r.Range(func(k, v interface{}) bool{
		session := v.(*dbops.Session)
		sessionMap.Store(k, session)
		return true
	})
}

func GenerateNewSession(userId int64) int64 {
	session, _ := dbops.InsertSession(userId)
	sessionMap.Store(session.Id, session)
	return session.Id
}

func IsSessionExpired(sid int64) (int64, bool) {
  session,ok := sessionMap.Load(sid)
  if ok {
  	if session.(*dbops.Session).TTL < time.Now().Unix()/1000000 {
		deleteExpiredSession(sid)
  		return session.(*dbops.Session).Id, true
	}
	return session.(*dbops.Session).Id, false
  } else {
  	return session.(*dbops.Session).Id, true
  }
}

func deleteExpiredSession(sid int64) {
	sessionMap.Delete(sid)
	_ = dbops.DeleteSession(sid)
}
