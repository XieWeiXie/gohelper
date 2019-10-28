package gohelper

import (
	"fmt"
	"testing"
)

func TestNewMd5Action(t *testing.T) {
	m := NewMd5Action("xiewei")
	m.Do()
	fmt.Println(m)
	fmt.Println(m.Random(32))
	m1 := NewMd5Action(
		">Z|w`w9zwxq9Ebd&aid=01A0QclqWHjYDiPccIVGsKkq8VuZpCUIaUQOAW6PCht1I4BUo.&cfrom=2899595010&count=6&cuid=1948244870&cursor=-1&noncestr=L9Qyq1sz08688Yh3ZFl44l7N11a66Q&platform=ANDROID&timestamp=1572250716691&ua=smartisan-OD103__oasis__1.4.17__Android__Android7.1.1&version=1.4.17&wm=9009_90025",
	)
	m1.Do()
	fmt.Println(m1)
}

/*
cursor=-1&ua=smartisan-OD103__oasis__1.4.17__Android__Android7.1.1&wm=9009_90025&aid=01A0QclqWHjYDiPccIVGsKkq8VuZpCUIaUQOAW6PCht1I4BUo.&cuid=1948244870&sign=2fd59e122983e46c1e95074a8da0191d&timestamp=1572250716691&cfrom=2899595010&count=6&version=1.4.17&noncestr=L9Qyq1sz08688Yh3ZFl44l7N11a66Q&platform=ANDROID

aid:01A0QclqWHjYDiPccIVGsKkq8VuZpCUIaUQOAW6PCht1I4BUo.
cfrom:2899595010
count:6
cuid:1948244870
cursor:-1
noncestr:L9Qyq1sz08688Yh3ZFl44l7N11a66Q
platform:ANDROID
timestamp:1572250716691
ua:smartisan-OD103__oasis__1.4.17__Android__Android7.1.1
version:1.4.17
wm:9009_90025
sign:2fd59e122983e46c1e95074a8da0191d
aid=01A0QclqWHjYDiPccIVGsKkq8VuZpCUIaUQOAW6PCht1I4BUo.&cfrom=2899595010&count=6&cuid=1948244870&cursor=-1&noncestr=L9Qyq1sz08688Yh3ZFl44l7N11a66Q&platform=ANDROID&timestamp=1572250716691&ua=smartisan-OD103__oasis__1.4.17__Android__Android7.1.1&version=1.4.17&wm=9009_90025
*/
