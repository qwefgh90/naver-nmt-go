# Naver Neural Machine Translation for Go

Go에서 [Naver Labs의 인공신경망 기반 번역 서비스](http://labspace.naver.com/nmt/)를 사용하기 위한 라이브러리.

## Install

```bash
$ go get -u github.com/meinside/naver-nmt-go
```

## Usage

```go
// sample code

package main

import (
	"fmt"

	nmt "github.com/meinside/naver-nmt-go"
)

func main() {
	var str string

	// kor => eng
	str = "테스트입니다"

	if translated, err := nmt.Translate(str, nmt.Korean, nmt.English); err == nil {
		fmt.Printf("> %s => %s\n", str, translated)
	} else {
		fmt.Printf("* error = %s\n", err)
	}

	// eng => kor
	str = "This is for testing"

	if translated, err := nmt.Translate(str, nmt.English, nmt.Korean); err == nil {
		fmt.Printf("> %s => %s\n", str, translated)
	} else {
		fmt.Printf("* error = %s\n", err)
	}
}
```
