package pkg

import (
	"fmt"
	"testing"
)

func TestParseToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBY2NvdW50IjoiMjAyMTIxMzk5MCIsImV4cCI6MTY4ODQ3MDM3MCwiaXNzIjoiQ0NOVS1Jbm4ifQ.lboPXTIdBoYSAL3_4so7ocppy2GQS-N7mo4iCEaRvq4"
	fmt.Println(ParseToken(token))
}
