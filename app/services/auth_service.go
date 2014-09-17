package services

import (
	"encoding/json"
	"errors"
	"github.com/stsmedia/thingbricks/app/models"
	"github.com/stsmedia/thingbricks/app/persistence"
	"net/http"
	"strings"
)

type AuthService struct {
}

func (l *AuthService) FindByAccountGroup(accountGroupId int64) []*models.Login {
	var logins []*models.Login
	_, err := persistence.Dbm.Select(&logins, "select * from logins where account_group_id = $1", accountGroupId)
	checkErr(err, "no logins found")
	return logins
}

func (l *AuthService) FindByEmail(email string) []*models.Login {
	var logins []*models.Login
	persistence.Dbm.Select(&logins, "select * from logins where email = $1", email)
	return logins
}

func (l *AuthService) Verify(login *models.Login) (*models.Login, error) {
	if login.AccessToken != "" && login.Network != "" {
		//assume its an oauth source
		err := populateUserDetails(login)
		if err != nil {
			return nil, err
		}
	}
	if login.Email == "" {
		return nil, errors.New("No email available")
	}
	login.Email = strings.ToLower(login.Email)
	existing := l.FindByEmail(login.Email)
	if len(existing) == 0 {
		persistence.Dbm.Insert(login)
	}
	return login, nil
}

func (l *AuthService) Delete(accountGroupId int64, id int64) bool {
	_, err := persistence.Dbm.Exec("delete from logins where account_group_id = $1 and id=$2", accountGroupId, id)
	if err != nil {
		checkErr(err, "login not deleted")
		return false
	}
	return true
}

func populateUserDetails(login *models.Login) error {
	switch login.Network {
	case "facebook":
		{
			response, err := http.Get("https://graph.facebook.com/me?access_token=" + login.AccessToken)
			if err != nil {
				return err
			} else {
				type facebook struct {
					Email     string `json:"email"`
					FirstName string `json:"first_name"`
					LastName  string `json:"last_name"`
					Gender    string `json:"gender"`
				}
				var message facebook
				err := json.NewDecoder(response.Body).Decode(&message)
				if err != nil {
					return err
				}
				login.Email = message.Email
				login.FirstName = message.FirstName
				login.LastName = message.LastName
				login.Gender = message.Gender
			}
			response, err = http.Get("https://graph.facebook.com/me?fields=picture.type(small)&access_token=" + login.AccessToken)
			if err != nil {
				return err
			} else {
				type fb_data struct {
					Url string `json:"url"`
				}
				type fb_picture struct {
					Picture fb_data `json:"data"`
				}
				var message fb_picture
				err := json.NewDecoder(response.Body).Decode(&message)
				if err != nil {
					return err
				}
				login.Picture = message.Picture.Url
			}
			return nil
		}
	case "google":
		{
			type google struct {
				Email     string `json:"email"`
				FirstName string `json:"given_name"`
				LastName  string `json:"family_name"`
				Gender    string `json:"gender"`
				Picture   string `picture:"picture"`
			}
			response, err := http.Get("https://www.googleapis.com/oauth2/v1/userinfo?alt=json&access_token=" + login.AccessToken)
			if err != nil {
				return err
			} else {
				var message google
				err := json.NewDecoder(response.Body).Decode(&message)
				if err != nil {
					return err
				}
				login.Email = message.Email
				login.FirstName = message.FirstName
				login.LastName = message.LastName
				login.Gender = message.Gender
				login.Picture = message.Picture
				return nil
			}
		}
	case "github":
		{
			type github struct {
				Login   string `json:"login"`
				Name    string `json:"name"`
				Email   string `json:"email"`
				Picture string `json:"avatar_url"`
			}
			response, err := http.Get("https://api.github.com/user?access_token=" + login.AccessToken)
			if err != nil {
				return err
			} else {
				var message github
				err := json.NewDecoder(response.Body).Decode(&message)
				if err != nil {
					return err
				}
				names := strings.Fields(message.Name)
				if len(names) > 0 {
					login.FirstName = names[0]
				}
				if len(names) > 1 {
					login.LastName = names[1]
				}
				if message.Email == "" {
					login.Email = message.Login
				} else {
					login.Email = message.Email
				}
				login.Picture = message.Picture
				return nil
			}
		}
	}
	return errors.New("network not supported")
}
