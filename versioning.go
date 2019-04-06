package versioning 

type Versioning interface {
  Newer(s Versioning) int
  IsNil() bool
  Print() string
}
