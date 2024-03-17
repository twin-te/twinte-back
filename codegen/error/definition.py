from typing import NewType

ConnectErrorCode = NewType("ConnectErrorCode", str)

Canceled = ConnectErrorCode("Canceled")
Unknown = ConnectErrorCode("Unknown")
InvalidArgument = ConnectErrorCode("InvalidArgument")
DeadlineExceeded = ConnectErrorCode("DeadlineExceeded")
NotFound = ConnectErrorCode("NotFound")
AlreadyExists = ConnectErrorCode("AlreadyExists")
PermissionDenied = ConnectErrorCode("PermissionDenied")
ResourceExhausted = ConnectErrorCode("ResourceExhausted")
FailedPrecondition = ConnectErrorCode("FailedPrecondition")
Aborted = ConnectErrorCode("Aborted")
OutOfRange = ConnectErrorCode("OutOfRange")
Unimplemented = ConnectErrorCode("Unimplemented")
Internal = ConnectErrorCode("Internal")
Unavailable = ConnectErrorCode("Unavailable")
DataLoss = ConnectErrorCode("DataLoss")
Unauthenticated = ConnectErrorCode("Unauthenticated")

definition: dict[str, list[tuple[str, ConnectErrorCode]]] = {
    "announcement": [
        ("AnnouncementNotFound", NotFound),
    ],
    "auth": [
        ("UserAuthenticationAlreadyExists", AlreadyExists),
        ("UserAuthenticationNotAssociated", FailedPrecondition),
        ("UserHasAtLeastOneAuthentication", FailedPrecondition),
        (
            "UserHasAtMostOneAuthenticationFromSameProvider",
            FailedPrecondition,
        ),
    ],
    "donation": [("SubscriptionNotFound", NotFound)],
    "schoolcalendar": [
        ("ModuleNotFound", NotFound),
    ],
    "shared": [
        ("AlreadyExists", AlreadyExists),
        ("InvalidArgument", InvalidArgument),
        ("NotFound", NotFound),
        ("Unauthenticated", Unauthenticated),
        ("Unauthorized", PermissionDenied),
    ],
    "timetable": [
        ("CourseNotFound", NotFound),
        ("RegisteredCourseAlreadyExists", AlreadyExists),
        ("RegisteredCourseNotFound", NotFound),
        ("TagNotFound", NotFound),
    ],
}
