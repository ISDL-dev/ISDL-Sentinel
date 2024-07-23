/*
 * ISDL Sentinel API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package schema

type GetUserById200Response struct {

	UserId int32 `json:"user_id"`

	UserName string `json:"user_name"`

	MailAddress string `json:"mail_address"`

	NumberOfCoin int32 `json:"number_of_coin"`

	AttendanceDays int32 `json:"attendance_days"`

	StayTime string `json:"stay_time"`

	Status string `json:"status"`

	Place string `json:"place"`

	Grade string `json:"grade"`

	AvatarId int32 `json:"avatar_id"`

	AvatarImgPath string `json:"avatar_img_path"`

	AvatarList []GetUserById200ResponseAvatarListInner `json:"avatar_list"`
}
