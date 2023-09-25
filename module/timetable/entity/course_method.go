package timetableentity

type CourseMethod int

const (
	CourseMethodOnlineAsynchronous CourseMethod = iota + 1
	CourseMethodOnlineSynchronous
	CourseMethodFaceToFace
	CourseMethodOthers
)
