package versioning 

var MvnMaturityLevels = map[string]int { 
  "SNAPSHOT" : 0,
  "RELEASE" : 1,
}

type MavenVer struct {
  Major int
  Minor int
  Patch int 
  Maturity string 
}


// res = 0, v2 older than v 
// res = 1, v2 newer than v
// res = 2, v equals v2
func (v MavenVer) Newer(v2 MavenVer) (res int) {
	
  res = 1	

  if (MvnMaturityLevels[v.Maturity] < MvnMaturityLevels[v2.Maturity]) {
    res = 0 
  } else {   
    if (MvnMaturityLevels[v.Maturity] == MvnMaturityLevels[v2.Maturity]) {
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

func (v MavenVer) IsNil() bool {
  return v.Major == 0 && v.Minor == 0 && v.Patch == 0
}

