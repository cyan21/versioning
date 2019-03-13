package versioning

import (
  "fmt"
)

func ExampleNewer() {
  sv01 := SemVer20{Major: 0, Minor: 0, Patch: 1, Maturity: "alpha"}
  sv02 := SemVer20{Major: 0, Minor: 2, Patch: 0, Maturity: "alpha"}
  fmt.Printf("%v\n", sv01.Newer(sv02))

  sv1 := SemVer20{Major: 0, Minor: 2, Patch: 0, Maturity: "alpha"}
  sv2 := SemVer20{Major: 0, Minor: 12, Patch: 0, Maturity: "alpha"}
  fmt.Printf("%v\n", sv1.Newer(sv2))

  sv11 := SemVer20{Major: 1, Minor: 2, Patch: 0, Maturity: "alpha"}
  sv22 := SemVer20{Major: 0, Minor: 12, Patch: 0, Maturity: "alpha"}
  fmt.Printf("%v\n", sv11.Newer(sv22))

  sv111 := SemVer20{Major: 0, Minor: 2, Patch: 0, Maturity: "beta"}
  sv222 := SemVer20{Major: 0, Minor: 12, Patch: 0, Maturity: "alpha"}
  fmt.Printf("%v\n", sv111.Newer(sv222))

  sv1111 := SemVer20{Major: 0, Minor: 12, Patch: 0, Maturity: "release"}
  sv2222 := SemVer20{Major: 0, Minor: 12, Patch: 0, Maturity: "release"}
  fmt.Printf("%v\n", sv1111.Newer(sv2222))

  sv := SemVer20{Major: 0, Minor: 0, Patch: 0, Maturity: ""}
  fmt.Println(sv.IsNil())

  // Output:
  // 0  
  // 0  
  // 1  
  // 1  
  // 2
  // true 

}
