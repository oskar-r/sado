package utility

import (
	"testing"
	"time"
)

type User struct {
	Username          string     `json:"username,omitempty"`
	UserID            int        `db:"ID" json:"user_id"`
	EncryptedPassword string     `db:"user_pass" json:"-"`
	UserLogin         string     `db:"user_login" json:"-"`
	Email             string     `db:"user_email" json:"email,omitempty"`
	School            string     `db:"school" json:"school,omitempty"`
	SchoolType        string     `db:"school_type" json:"school_type,omitempty"`
	AccountExpDate    *time.Time `db:"account_exp_date" json:"account_exp_date,omitempty"`
	EulaApproved      bool       `db:"eula_approved" json:"eula_approved,omitempty"`
	Deleted           int        `db:"deleted" json:"-"`
	Roles             []string   `json:"roles,omitempty"` //The roles that the use can have
	Role              string     `db:"role,omitempty"`    //Currently active role
}

type UserAccess struct {
	ClassroomUser     bool   `db:"classroom_user" json:"classroom_user,omitempty"`
	MatrixAccess      bool   `db:"matrix_access" json:"matrix_access,omitempty"`
	SFIAccess         bool   `db:"sfi" json:"sfi_access,omitempty"`
	AccountType       string `db:"account_type" json:"account_type,omitempty"`
	SubscriptionTopic string `db:"subscription_topic" json:"subscription_topic,omitempty"`
}

func TestMarshalToJSON(t *testing.T) {
	type args struct {
		wrapper []string
		i       []interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"Test_1",
			args{
				wrapper: []string{"profile", "access"},
				i: []interface{}{
					User{
						Username: "test",
						UserID:   4,
					},
					UserAccess{
						ClassroomUser: true,
						AccountType:   "premium",
					},
				},
			},
			`{"profile":{"username":"test","user_id":4,"Role":""},"access":{"classroom_user":true,"account_type":"premium"}}`,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := MarshalToJSON(tt.args.wrapper, tt.args.i...); got != tt.want {
				if (err != nil) != tt.wantErr {
					t.Errorf("MarshalToJSON() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				t.Errorf("MarshalToJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
