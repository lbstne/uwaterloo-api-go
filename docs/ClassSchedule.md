# ClassSchedule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CourseId** | **string** | Course identifier number, not unique | [optional] [default to null]
**CourseOfferNumber** | **int32** | Course offer number identifier for this class | [optional] [default to null]
**SessionCode** | **string** | Session code number for this class | [optional] [default to null]
**ClassSection** | **int32** | Number identifying the section of this class | [optional] [default to null]
**TermCode** | **string** | Waterloo academic term code | [optional] [default to null]
**ClassMeetingNumber** | **int32** | Identifier for the class meeting this schedule data relates to | [optional] [default to null]
**ScheduleStartDate** | [**time.Time**](time.Time.md) | The date this schedule begins | [optional] [default to null]
**ScheduleEndDate** | [**time.Time**](time.Time.md) | The date this schedule ends | [optional] [default to null]
**ClassMeetingStartTime** | [**time.Time**](time.Time.md) | The start time of this class | [optional] [default to null]
**ClassMeetingEndTime** | [**time.Time**](time.Time.md) | The end time of this class | [optional] [default to null]
**ClassMeetingDayPatternCode** | **string** | A code representing the days the class schedule takes place during the week | [optional] [default to null]
**ClassMeetingWeekPatternCode** | **string** | A Y|N per day representation of the class schedule taking place during the week | [optional] [default to null]
**LocationName** | **string** | REMOVED for privacy. Building and room name for the location of this class schedule | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

