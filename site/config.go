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
	c.Domain = "gnfull.com"
	c.Author = "박아영"
	c.Title = "박아영의 강남풀싸롱 대모험"
	c.Description = "박아영 실장이 소개하는 강남풀싸롱 정찰제 가격 및 시스템 안내. 풀싸롱, 미러룸, 매직미러, 야구장, 쩜오, 하이퍼블릭, 셔츠룸, 가라오케, 레깅스룸 등 강남의 모든 유흥업소 정보를 한눈에!"
	k := Keywords([]string{"박아영의 강남풀싸롱 대모험", "박아영 실장", "강남 풀싸롱", "강남 미러룸", "강남 매직미러", "강남 야구장", "강남 하이퍼블릭", "강남 레깅스룸", "강남 쩜오", "강남 호빠", "강남 클럽", "강남 셔츠룸", "강남 가라오케"})
	c.Keywords = &k
	c.DatePublished = date(2023, 12, 18)
	c.DateModified = date(2023, 12, 18)
	// 업종마다 전화번호가 다른경우 store/store.go 파일의 setPhoneNumber 함수에서 하드코딩
	c.PhoneNumber = "010-8259-7110"
	c.SearchEngineConnection = &searchEngineConnection{
		Google: "VnHLGCnAiT6dJHiUdLJXFtgbUBeahGlL8-KBT_LCCx4",
	}
	Config = c
}
