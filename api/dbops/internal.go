package dbops

import (
	"sync"
	"time"
)

func InsertSession(userId int64) (Session, error) {
	ttl := time.Now().Unix()/1000000 + 30 * 60 * 1000
	session := &Session{
		TTL: ttl,
		UserId: userId,
	}
	_, err = db.Model(session).Insert()
	return *session, err
}

func GetSession(sid int64) (Session, error) {
	var session Session
	err = db.Model(&session).
		Where("id = ?", sid).
		Select()
	return session, err
}

func ListSessions() (*sync.Map, error) {
	m := &sync.Map{}
	var sessions []Session
	err = db.Model(&sessions).Select()

	for _, session := range sessions {
		m.Store(session.Id, session)
	}
	return m, err
}

func DeleteSession(sid int64) error {
	var session Session
	_, err := db.Model(session).Where("id = ?", sid).Delete()
	return err
}
