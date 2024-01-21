/*
 * Waterloo OpenData API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// An academic Subject at Waterloo describes an area that a Course can be in
type Subject struct {
	// Code that identifies this Subject
	Code string `json:"code,omitempty"`
	// The Name for this Subject, most often the displayed name
	Name string `json:"name,omitempty"`
	// The short description of this subject, often same as Code
	DescriptionAbbreviated string `json:"descriptionAbbreviated,omitempty"`
	// Description of the Subject
	Description string `json:"description,omitempty"`
	// Code for the Academic Organization that is associated to this Subject
	AssociatedAcademicOrgCode string `json:"associatedAcademicOrgCode,omitempty"`
}
