package openapi

import (
	"bytes"
	"encoding/json"
)

func FromRegisteredCourse(registeredCourse RegisteredCourse) (res PostRegisteredCourses200JSONResponse, err error) {
	res.union, err = json.Marshal(registeredCourse)
	return
}

func FromRegisteredCourses(registeredCourses []RegisteredCourse) (res PostRegisteredCourses200JSONResponse, err error) {
	res.union, err = json.Marshal(registeredCourses)
	return
}

func (body *PostRegisteredCoursesJSONRequestBody) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &body.union)
}

func ToPostRegisteredCoursesJSONBody0(reqBody *PostRegisteredCoursesJSONRequestBody) (ret PostRegisteredCoursesJSONBody0, err error) {
	dec := json.NewDecoder(bytes.NewReader(reqBody.union))
	dec.DisallowUnknownFields()
	err = dec.Decode(&ret)
	return
}

func ToPostRegisteredCoursesJSONBody1(reqBody *PostRegisteredCoursesJSONRequestBody) (ret PostRegisteredCoursesJSONBody1, err error) {
	dec := json.NewDecoder(bytes.NewReader(reqBody.union))
	dec.DisallowUnknownFields()
	err = dec.Decode(&ret)
	return
}

func ToPostRegisteredCoursesJSONBody2(reqBody *PostRegisteredCoursesJSONRequestBody) (ret PostRegisteredCoursesJSONBody2, err error) {
	dec := json.NewDecoder(bytes.NewReader(reqBody.union))
	dec.DisallowUnknownFields()
	err = dec.Decode(&ret)
	return
}
