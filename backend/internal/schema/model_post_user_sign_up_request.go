package schema

type PostUserSignUpRequest struct {

    Name         string `json:"name" form:"name"`

    AuthUserName string `json:"auth_user_name" form:"auth_user_name"`

    MailAddress  string `json:"mail_address" form:"mail_address"`

    Password     string `json:"password" form:"password"`
	
    GradeID      int    `json:"grade_id" form:"grade_id"`
}