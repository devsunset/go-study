package test_test

import (
	"example/test"
	"testing"
)

func TestSum(t *testing.T) {
	s := test.Sum(1, 2, 3)
	if s != 6 {
		t.Error("Fail")
	}
}

// go test

// 1. go test 는 현재 명령어 실행 폴더에 있는 *_test.go 파일들을 인식 후 일괄 실행
// 2. package *_test
// 3. import testing
// 4. TestXxx 형식의 Method 인자는 (t *testing.T) 테스트 메소드 작성 출력은 없음
// 5. 테스트 에러를 표시하기 위해 testing.T 의 Error(), Fail() 등의 메서드들을 사용한다.
