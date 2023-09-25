下記のGoの構造体に対応するPostgreSQL上のテーブルを作成するためのクエリを生成して下さい。

```
type RegisteredCourse struct {
	ID         uuid.UUID
	UserID     uuid.UUID
	CourseID   *uuid.UUID
	Year       uint16
	Name       *string
	Instructor *string
	Cregit     *uint8
	Methods    *[]CourseMethod
	Schedules  *[]CourseSchedule
	Memo       string
	Attendance int32
	Absence    int32
	Late       int32
	TagIDs     []*uuid.UUID
}
```