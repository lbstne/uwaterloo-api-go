# Class

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CourseId** | **string** | Course identifier number, not unique | [optional] [default to null]
**CourseOfferNumber** | **int32** | Course offer number identifier for this class | [optional] [default to null]
**SessionCode** | **string** | Session code for this class | [optional] [default to null]
**ClassSection** | **int32** | Number identifying the section of this class | [optional] [default to null]
**TermCode** | **string** | Waterloo academic term code | [optional] [default to null]
**ClassNumber** | **int32** | Class number identifier for this class | [optional] [default to null]
**CourseComponent** | **string** | Course component code for this class | [optional] [default to null]
**AssociatedClassCode** | **int32** | Associated class code for this class | [optional] [default to null]
**MaxEnrollmentCapacity** | **int32** | Indicates the maximum number of students that can enroll in this class | [optional] [default to null]
**EnrolledStudents** | **int32** | Indicates the current number of students enrolled in this class | [optional] [default to null]
**EnrollConsentCode** | **string** | Code describing whether No, Instructor, or Department consent to enroll is required. Overrides Course level information if different. | [optional] [default to null]
**EnrollConsentDescription** | **string** | Description of the enroll requirement. Overrides Course level information if different. | [optional] [default to null]
**DropConsentCode** | **string** | Code describing whether No, Instructor, or Department consent to drop is required. Overrides Course level information if different. | [optional] [default to null]
**DropConsentDescription** | **string** | Description of the drop requirement. Overrides Course level information if different. | [optional] [default to null]
**ScheduleData** | [**[]ClassSchedule**](ClassSchedule.md) | Schedule data for this class | [optional] [default to null]
**InstructorData** | [**[]ClassInstructor**](ClassInstructor.md) | Instructor data for this class | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

