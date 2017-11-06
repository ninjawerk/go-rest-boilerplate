/**
 * Created by desha on 10/24/2017
 */

package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string
	Password string
	Email    string
	Avatar   string
}

func (this User)Update() (bool, error) {
	return false, nil
}
func (this User)Create() (bool, error) {
	return false, nil
}
func (this User)Delete() (bool, error) {
	return false, nil
}
