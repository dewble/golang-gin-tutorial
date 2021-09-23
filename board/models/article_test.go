package models

import "testing"

// Test for getAllarticles
func TestGetAllArticles(t *testing.T) {
	alist := getAllArticles()

	// 1. 가져온 게시글 리스트와 전역 변수 게시글 리스트의 길이가 동일한지 확인
	// 2. 각 게시글이 동일한지 확인하기 위해 반복문 사용. 1,2 둘중 어느것이라도 통과못할 경우 테스트 실패
	if len(alist) != len(articleList) {
		t.Fail()
	}

	for i, v := range alist {
		if v.Content != articleList[i].Content ||
			v.ID != articleList[i].ID ||
			v.Title != articleList[i].Title {

			t.Fail()
			break
		}
	}
}
