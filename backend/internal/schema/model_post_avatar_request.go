/*
 * ISDL Sentinel API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package schema

type PostAvatarRequest struct {

	UserId int32 `json:"user_id"`

	AvatarImgPath string `json:"avatar_img_path"`
}
