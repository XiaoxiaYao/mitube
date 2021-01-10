package dbops

import "time"

func AddUserCredential(userName string, pwd string) error {
	user := &User{
		UserName: userName,
		Pwd: pwd,
	}
	_, err = db.Model(user).Insert()
	return err
}

func GetUserCredential(userName string) (User, error) {
	var user User
	err = db.Model(&user).
		Where("username = ?", userName).
		Select()
	return user, err
}

func AddVideo(userId int64, name string) error {
	video := &Video{
		Name: name,
		AuthorId: userId,
		CreatedAt: time.Now(),
	}
	_, err = db.Model(video).Insert()
	return err
}

func ListVideos() ([]Video, error) {
	var videos []Video
	err = db.Model(&videos).Select()
	return videos, err
}
