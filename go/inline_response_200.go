/*
 * Swagger Blog
 *
 * A Simple Blog
 *
 * API version: 1.0.0
 * Contact: apiteam@swagger.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type InlineResponse200 struct {

	Id int64 `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Tags []Tag `json:"tags,omitempty"`

	Author string `json:"author,omitempty"`

	Date string `json:"date,omitempty"`
}
