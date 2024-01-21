/*
 * Waterloo OpenData API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

// Model representing a WCMS Opportunity (Job)
type SiteOpportunity struct {
	// Unique, numeric site ID
	SiteId int32 `json:"siteId,omitempty"`
	// Unique Id of this opportunity item
	UniqueKey string `json:"uniqueKey,omitempty"`
	// Title of the opportunity
	Title string `json:"title,omitempty"`
	// Username of the user that published this item
	PublisherUsername string `json:"publisherUsername,omitempty"`
	// Last updated date
	UpdatedDate time.Time `json:"updatedDate,omitempty"`
	// Type of opportunity (ie: volunteer, paid)
	OpportunityType string `json:"opportunityType,omitempty"`
	// Employment type (ie: part, full, other)
	EmploymentType string `json:"employmentType,omitempty"`
	// Rate of paay description
	RateOfPay string `json:"rateOfPay,omitempty"`
	// Rate of pay type
	RateOfPayType string `json:"rateOfPayType,omitempty"`
	// Opportunity description/content, usually includes HTML markup
	Content string `json:"content,omitempty"`
	// Posted or open for application date
	PostedOrOpenDate time.Time `json:"postedOrOpenDate,omitempty"`
	// Posted or open date time zone
	PostedOrOpenDateTimeZone string `json:"postedOrOpenDateTimeZone,omitempty"`
	// Opportunity application deadline date
	ApplicationDeadlineDate time.Time `json:"applicationDeadlineDate,omitempty"`
	// Start date for the opportunity
	StartDate time.Time `json:"startDate,omitempty"`
	// End date for the opportunity
	EndDate time.Time `json:"endDate,omitempty"`
	// Number of positions available for this opportunity
	NumberOfPositions string `json:"numberOfPositions,omitempty"`
	// URI for an external applicaton website
	ApplicationUri string `json:"applicationUri,omitempty"`
}