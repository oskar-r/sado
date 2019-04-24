package models

import "time"

type User struct {
	Username          string   `db:"username" json:"username,omitempty"`
	UserID            string   `db:"user_id" json:"-"`
	EncryptedPassword string   `db:"user_pass" json:"-"`
	CurrentRole       string   `json:"current_role,omitempty"`
	Roles             []string `json:"roles,omitempty"`
	MyBucket          string   `json:"my_bucket,omitempty"`
}

type NewUser struct {
	Username       string `json:"username,omitempty"`
	Password       string `json:"password,omitempty"`
	MyBucket       string `json:"my_bucket,omitempty"`
	UserID         string `json:"user_id,omitempty"`
	BCryptPassword string `json:"-"`
	AESPassword    string `json:"-"`
}

type AppConfig struct {
	Routes []struct {
		Order        int    `json:"order"`
		Icon         string `json:"icon"`
		To           string `json:"to"`
		Text         string `json:"text"`
		Dropdown     bool   `json:"dropdown"`
		DropdownData string `json:"dropdown_data"`
	} `json:"routes,omitempty"`
	ActiveRole string   `json:"active_role,omitempty"`
	Roles      []string `json:"roles,omitempty"`
}

type DataSet struct {
	Name         string    `json:"name,omitempty"`
	ContentType  string    `json:"content_type,omitempty"`
	Size         int64     `json:"size,omitempty"`
	Category     string    `json:"category,omitempty"`
	LastModified time.Time `json:"last_modified,omitempty"`
}
