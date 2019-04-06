package versioning 

import (
    "strconv"
)

var MaturityLevels = map[string]int { 
  "alpha" : 0,
  "beta" : 1, 
  "rc" :  2, 
  "release" : 3,
}

type SemVer20 struct {
  Major int
  Minor int
  Patch int 
  Maturity string 
}


// res = 0, v2 older than v 
// res = 1, v2 newer than v
// res = 2, v equals v2
func (v SemVer20) Newer(v1 Versioning) (res int) {
	
  res = 1	
  v2 := v1.(SemVer20)
  if (MaturityLevels[v.Maturity] < MaturityLevels[v2.Maturity]) {
    res = 0 
  } else {   
    if (MaturityLevels[v.Maturity] == MaturityLevels[v2.Maturity]) {
      if (v.Major < v2.Major) {
        res = 0	 
      } else {
        if (v.Major == v2.Major) {
          if (v.Minor < v2.Minor) {
            res = 0 
          } else {
	    if (v.Minor == v2.Minor) {
	    res = 2 
	    } 
	  }
        }
      } 
    }
  }
  return res 
}

func (v SemVer20) IsNil() bool {
  return v.Major == 0 && v.Minor == 0 && v.Patch == 0
}

func (v SemVer20) Print() string {
  return strconv.Itoa(v.Major) + "." + strconv.Itoa(v.Minor) + "." + strconv.Itoa(v.Patch)
// + "-" + v.Maturity
}
