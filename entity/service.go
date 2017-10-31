package entity

import (
	"errors"
	"fmt"
	"regexp"
)

func UserLogin(username string, password string) error {
	usr, err := GetCurUser()
	if err == nil {
		return errors.New("Error. You have login.Your username is:" + usr.Name)
	}
	var myuser = User{Name: username, Password: password}
	var checkuser = func(user *User) bool {
		if myuser.Name == user.Name && myuser.Password == user.Password {
			return true
		} else {
			return false
		}
	}
	user := QueryUser(checkuser)
	if len(user) == 0 {
		return errors.New("Error. Your name or password wrong")
	}

	SetCurUser(&user[0])
	return nil
}


func UserLogout() error {
	SetCurUser(&User{})
	return nil
}


func UserRegister(name string, password string,
	email string, phone string) error {
	_, err := GetCurUser()
	if err == nil {
		return errors.New("Error, can't do this operation. You hava login")
	}

	err = nil
	var myuser = User{name, password, email, phone}
	var checkuser = func(user *User) bool {
		if myuser.Name == user.Name {
			return true
		} else {
			return false
		}
	}
	if len(QueryUser(checkuser)) != 0 {
		err = errors.New("Some errors have taken place.")
		fmt.Println("the username has been used.")
		errLog.Println("the username has been used.")
	}

	matched, _ := regexp.MatchString("[0-9]|[A-Z]|[a-z]", password)
	if !matched {
		err = errors.New("Some errors have taken place")
		fmt.Println("password format error.")
		errLog.Println("password format error.")
	}

	matched, _ = regexp.MatchString("([0-9]|[A-Z]|[a-z])+@([0-9]|[a-z])+\\.+?com", email)
	if !matched {
		err = errors.New("Some errors have taken place")
		fmt.Println("email format error.")
		errLog.Println("email format error.")
	}

	matched, _ = regexp.MatchString("[0-9]{11}", phone)
	if !matched {
		err = errors.New("Some errors have taken place")
		fmt.Println("phone format error.")
		errLog.Println("phone format error.")
	}

	if err != nil {
		return err
	}

	CreateUser(&myuser)
	return nil
}


func UserLogOff() error {
	usr, err := GetCurUser()
	if err != nil {
		return err
	}

	var checkusr = func(user *User) bool {
		if usr.Name == user.Name {
			return true
		}
		return false
	}
	if DeleteUser(checkusr) != 1 {
		return errors.New("Error when log off")
	}

	DeleteAllMeeting()


	var checkmeeting = func(meeting *Meeting) bool {
		for _, participator := range meeting.Participators {
			if participator == usr.Name {
				return true
			}
		}
		return false
	}
	var deletepartic = func(meeting *Meeting) {
		for i, participator := range meeting.Participators {
			if participator == usr.Name {
				meeting.Participators = append(meeting.Participators[:i], meeting.Participators[i+1:]...)
			}
		}

		if len(meeting.Participators) == 0 {
			DeleteMeeting(func(meet *Meeting) bool {
				if meeting.Sponsor == meet.Sponsor {
					return true
				}
				return false
			})
		}
	}
	UpdateMeeting(checkmeeting, deletepartic)


	SetCurUser(&User{})
	return nil
}


func ListAllUser() error {
	var checkusr = func(user *User) bool {
		return true
	}

	fmt.Printf("%-.10q  %-20q  %-11q\n", "Name", "Email", "Phone")
	loginLog.Printf("%-.10q  %-20q  %-11q\n", "Name", "Email", "Phone")
	for i, usr := range QueryUser(checkusr) {
		if i == 0 && len(usr.Name) == 0 {
			return errors.New("Error in list user")
		}
		fmt.Printf("%-.10q  %-20q  %-11q\n", usr.Name, usr.Email, usr.Phone)
		loginLog.Printf("%-.10q  %-20q  %-11q\n", usr.Name, usr.Email, usr.Phone)
	}
	return nil
}


func MeetingCreate(title string, participators []string, sdate string,
	edate string) error {

	user, err := GetCurUser()
	if err != nil {
		return err
	}

	var checkmeeting = func(meeting *Meeting) bool {
		if user.Name == meeting.Sponsor && title == meeting.Title {
			return true
		}
		return false
	}
	if len(QueryMeeting(checkmeeting)) != 0 {
		return errors.New("Error. The meetng has exist")
	}


	if !checkParticipator(participators, Meeting{}) {
		return errors.New("Error. Some Participator not exist")
	}


	meeting := Meeting{user.Name, participators, StringToDate(sdate),
		StringToDate(edate), title}

	CreateMeeting(&meeting)
	return nil
}


func AddMeetingParticipator(title string, participators []string) error {
	user, err := GetCurUser()
	if err != nil {
		return err
	}

	flag, meeting := checkMeeting(title, user.Name)
	if !flag {
		return errors.New("Error The meeting not exist")
	}


	if !checkParticipator(participators, meeting[0]) {
		return errors.New("Error. Some Participators not exist or have been in the meeting")
	}


	var checkmeeting = func(meet *Meeting) bool {
		if meeting[0].Sponsor == meet.Sponsor &&
			meeting[0].Title == meet.Title {
			return true
		}
		return false
	}
	var addparticipator = func(meet *Meeting) {
		for _, participator := range participators {
			meet.Participators = append(meet.Participators, participator)
		}
	}
	UpdateMeeting(checkmeeting, addparticipator)
	return nil
}


func RemoveParticipator(title string, participators []string) error {

	user, err := GetCurUser()
	if err != nil {
		return err
	}


	flag, meeting := checkMeeting(title, user.Name)
	if !flag {
		return errors.New("Error The meeting not exist")
	}


	var all_exist bool = true
	for _, participator := range participators {
		var exist bool = false
		for _, meetparticipator := range meeting[0].Participators {
			if participator == meetparticipator {
				exist = true
				break
			}
		}
		if !exist {
			fmt.Printf("%s not exist in meeting participators.", participator)
			all_exist = false
		}
	}
	if !all_exist {
		return errors.New("Error. Some participator not exisit in meeting")
	}


	var checkmeeting = func(meet *Meeting) bool {
		if meeting[0].Sponsor == meet.Sponsor &&
			meeting[0].Title == meet.Title {
			return true
		}
		return false
	}
	var removeparticipator = func(meeting *Meeting) {
		if len(participators) == len(meeting.Participators) {
			DeleteMeeting(checkmeeting)
		} else {
			for _, participator := range participators {
				for i, meetparticipator := range meeting.Participators {
					if participator == meetparticipator {
						meeting.Participators = append(meeting.Participators[:i], meeting.Participators[i+1:]...)
					}
				}
			}
		}
	}
	UpdateMeeting(checkmeeting, removeparticipator)
	return nil
}

func ListMeeting(tmp_sDate string, tmp_eDate string) error {
	user, err := GetCurUser()
	if err != nil {
		return err
	}

	sDate := StringToDate(tmp_sDate)
	eDate := StringToDate(tmp_eDate)

	var checkMeeting = func(meeting *Meeting) bool {
		sponsor_ok := user.Name == meeting.Sponsor
		var participator_ok bool = false
		for _, participator := range meeting.Participators {
			if user.Name == participator {
				participator_ok = true
			}
		}
		if (sponsor_ok || participator_ok) && sDate.LessOrEqual(eDate) {
			return true
		}

		return false
	}

	meetings := QueryMeeting(checkMeeting)
	fmt.Printf("%.20s  %.16s  %.16s  %.20s  %s\n", "Title", "StartDate", "EndDate",
		"Sponsor", "Participators")
	loginLog.Printf("%.20s  %.16s  %.16s  %.20s  %s\n", "Title", "StartDate", "EndDate",
		"Sponsor", "Participators")
	for _, meeting := range meetings {
		fmt.Printf("%.20s  %.16s  %.16s  %.20s  %s\n",
			meeting.Title, DateToString(meeting.StartDate), DateToString(meeting.EndDate),
			meeting.Sponsor, meeting.Participators)
		loginLog.Printf("%.20s  %.16s  %.16s  %.20s  %s\n",
			meeting.Title, DateToString(meeting.StartDate), DateToString(meeting.EndDate),
			meeting.Sponsor, meeting.Participators)
	}
	return nil
}


func DeleteAMeeting(title string) error {
	user, err := GetCurUser()
	if err != nil {
		return err
	}

	var checkMeeting = func(meeting *Meeting) bool {
		if user.Name == meeting.Sponsor && title == meeting.Title {
			return true
		}
		return false
	}
	DeleteMeeting(checkMeeting)
	return nil
}

func DeleteAllMeeting() error {
	user, err := GetCurUser()
	if err != nil {
		return err
	}

	var checkMeeting = func(meeting *Meeting) bool {
		if user.Name == meeting.Sponsor {
			return true
		}
		return false
	}
	DeleteMeeting(checkMeeting)
	return nil
}


func QuitMeeting(title string) error {
	user, err := GetCurUser()
	if err != nil {
		return err
	}

	var checkmeeting = func(meeting *Meeting) bool {
		if title == meeting.Title {
			for _, participator := range meeting.Participators {
				if participator == user.Name {
					return true
				}
			}
		}

		return false
	}
	var deletepartic = func(meeting *Meeting) {
		for i, participator := range meeting.Participators {
			if participator == user.Name {
				meeting.Participators = append(meeting.Participators[:i], meeting.Participators[i+1:]...)
			}
		}

		if len(meeting.Participators) == 0 {
			DeleteMeeting(func(meet *Meeting) bool {
				if meeting.Sponsor == meet.Sponsor {
					return true
				}
				return false
			})
		}
	}
	UpdateMeeting(checkmeeting, deletepartic)
	return nil
}


func checkParticipator(user []string, meeting Meeting) bool {
	var err bool = false
	if meeting.Title == "" {
		for _, usr := range user {
			var checkuser = func(user *User) bool {
				if usr == user.Name {
					return true
				}
				return false
			}
			if len(QueryUser(checkuser)) == 0 {
				fmt.Printf("the participator %s isn't exist.\n", usr)
				errLog.Printf("the participator %s isn't exist.\n", usr)
				err = true
			}
		}
	} else {
		for _, usr := range user {
			var checkuser = func(user *User) bool {
				if usr == user.Name {
					return true
				}
				return false
			}
			if len(QueryUser(checkuser)) == 0 {
				fmt.Printf("the participator %s isn't exist.\n", usr)
				errLog.Printf("the participator %s isn't exist.\n", usr)
				err = true
			}

			for _, participator := range meeting.Participators {
				if usr == participator {
					fmt.Printf("%s has been a participator.", usr)
					errLog.Printf("%s has been a participator.", usr)
					err = true
				}
			}

			if usr == meeting.Sponsor {
				fmt.Println("Participator can't be sponsor")
				errLog.Println("Participator can't be sponsor")
				err = true
			}
		}
	}
	if err == true {
		return false
	}
	return true
}

func checkMeeting(title string, name string) (bool, []Meeting) {
	var checkmeet = func(meeting *Meeting) bool {
		if title == meeting.Title && name == meeting.Sponsor {
			return true
		}
		return false
	}

	meeting := QueryMeeting(checkmeet)
	if len(meeting) == 0 {
		fmt.Println("Error. The meetng has exist.")
		errLog.Println("Error. The meetng has exist.")
		return false, meeting
	}
	return true, meeting
}
