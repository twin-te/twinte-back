layer
- domainmodel
- usecase IF
- usecase impl or interactor
- external IF
- external impl
- api

module package
- root
  - domainmodel
  - usecase IF
  - internal
    - usecase impl or interactor
    - external IF
    - external impl

- entity
- usecase


- module (including usecase IF)
  - entity
  - internal
    - port
    - external
    - usecase


- module
  - entity
  - usecaseport
  - usecaseimpl
  - repositoryport
  - repositoryimpl

- auth 認証・認可に関するモジュール
- timetable 講義
- schoolcalendar
- information
- usersetting
- feedback
- donation
