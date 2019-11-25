# Code Model

Core: Function, Parameter List, Return Type, Refs

Function (Ref) -> File -> Dir

Function (mapping) <-> Data Struct

Function (mapping) <-> Member

## Model

Entity: File
  VO: Full Name

Entity: Member
  VO: Name
  VO: Line Number
  VO: Data Struct ID
  VO: File ID
  VO: NameSpace / Package
  VOs: Position
    VO: Start
    VO: End

Entity: Data Struct
  VO: Member ID

Entity: Commit
  VO: File ID

Entity: Function
  VO: MemberId
  VOs: Parameter
    VO: Type
  VOs: Return Type
    VO: Type
  VOS: Refs
    VO: member ID
