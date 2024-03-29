/*
 * Waterloo OpenData API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Represents a scheduled Exam for a Class
type Exam struct {
	// The name of the Exam, representative of the Course name
	ExamDisplayName string `json:"examDisplayName,omitempty"`
	// Sections of the Class this Exam schedule is applicable to
	Sections string `json:"sections,omitempty"`
	// Indicates whether this Exam is held online, or provides an alternate description
	IsOnlineDescription string `json:"isOnlineDescription,omitempty"`
	// Day name on which this Exam is scheduled to take place
	Day string `json:"day,omitempty"`
	// Description of the location of the Exam
	Location string `json:"location,omitempty"`
	// The date the Exam is scheduled for
	ExamWindowStartDate string `json:"examWindowStartDate,omitempty"`
	// The time the Exam is scheduled to start
	ExamWindowStartTime string `json:"examWindowStartTime,omitempty"`
	// The scheduled duration of the Exam
	ExamDuration string `json:"examDuration,omitempty"`
	// Additional notes about this Exam schedule
	Notes string `json:"notes,omitempty"`
	// Term Code for the Term this Exam is associated with
	TermCode string `json:"termCode,omitempty"`
}
