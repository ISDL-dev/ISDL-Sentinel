/*
 * ISDL Sentinel API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package schema

type PutUserByIdRequest struct {

	UserName string `json:"user_name"`

	MailAddress string `json:"mail_address"`

	Grade string `json:"grade"`

	RoleList []string `json:"role_list"`
}
