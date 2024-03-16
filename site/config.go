package site

import (
	"strings"
	"time"
)

var Config *config

type Keywords []string

func (k *Keywords) String() string { return strings.Join(*k, ",") }

type searchEngineConnection struct {
	Google string
}

type config struct {
	Port                   uint32
	Domain                 string
	Author                 string
	Title                  string
	Description            string
	Keywords               *Keywords
	DatePublished          time.Time
	DateModified           time.Time
	PhoneNumber            string
	SearchEngineConnection *searchEngineConnection
}

func date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
}

func Init() {
	c := &config{}
	c.Port = uint32(80)
	c.Domain = "jinwoowide.com"
	c.Author = "안예린"
	c.Title = "안예린 실장의 강남풀싸롱 탐방기"
	c.Description = "예린 실장이 강추하는 모든 강남풀싸롱에 탐방기를 전해드려요! 풀싸롱, 미러룸, 매직미러, 야구장, 쩜오, 하이퍼블릭, 셔츠룸, 가라오케, 레깅스룸, 클럽, 호빠 등 강남지역 모든 유흥 정보 검색은 예린실장에게!"
	k := Keywords([]string{"예린 실장", "강남풀싸롱", "풀싸롱", "미러룸", "매직미러", "야구장", "쩜오", "하이퍼블릭", "셔츠룸", "가라오케", "레깅스룸", "클럽", "호빠"})
	c.Keywords = &k
	c.DatePublished = date(2024, 3, 11)
	c.DateModified = date(2024, 3, 11)
	// 업종마다 전화번호가 다른경우 store/store.go 파일의 setPhoneNumber 함수에서 하드코딩
	c.PhoneNumber = "010-7970-9057"
	c.SearchEngineConnection = &searchEngineConnection{
		Google: "z24kLW0whZciFbbZZsg0FkW68ZbJDm3WSgRABdoljNE",
	}
	Config = c
}
