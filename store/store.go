package store

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/jeonghoikun/gnfull.com/site"
)

const (
	STORE_TYPE_HIGHPUBLIC  string = "하이퍼블릭"
	STORE_TYPE_SHIRTROOM   string = "셔츠룸"
	STORE_TYPE_KARAOKE     string = "가라오케"
	STORE_TYPE_LEGGINGS    string = "레깅스룸"
	STORE_TYPE_DOT5        string = "쩜오"
	STORE_TYPE_HOBBA       string = "호빠"
	STORE_TYPE_CLUB        string = "클럽"
	STORE_TYPE_FULL        string = "풀싸롱"
	STORE_TYPE_MIRRORROOM  string = "미러룸"
	STORE_TYPE_MAGICMIRROR string = "매직미러"
	STORE_TYPE_YAGUJANG    string = "야구장"
)

var stores []*Store = []*Store{}

func Get(do, si, dong, storeType, title string) (o *Store, has bool) {
	for _, s := range stores {
		if s.Location.Do == do && s.Location.Si == si && s.Location.Dong == dong &&
			s.Type == storeType && s.Title == title {
			return s, true
		}
	}
	return nil, false
}

func ListAllStores() []*Store { return stores }

func ListStoresByDoSiAndStoreType(do, si, storeType string) []*Store {
	list := []*Store{}
	for _, s := range stores {
		if s.Location.Do == do && s.Location.Si == si && s.Type == storeType {
			list = append(list, s)
		}
	}
	return list
}

type Category struct {
	Name   string
	Stores []*Store
}

func ListAllCategories() []*Category {
	list := []*Category{}
	for _, s := range ListAllStores() {
		ok := false
		for _, c := range list {
			if s.Type == c.Name {
				ok = true
				break
			}
		}
		if !ok {
			list = append(list, &Category{
				Name:   s.Type,
				Stores: []*Store{s},
			})
			continue
		}
		for _, c := range list {
			if s.Type == c.Name {
				c.Stores = append(c.Stores, s)
			}
		}
	}
	sort.Slice(list, func(i, j int) bool { return list[i].Name < list[j].Name })
	for _, x := range list {
		sort.Slice(x.Stores, func(i, j int) bool {
			return x.Stores[i].DatePublished.UnixNano() > x.Stores[j].DatePublished.UnixNano()
		})
	}
	return list
}

type Location struct {
	// Do: ex) 서울
	Do string
	// Si: ex) 강남구
	Si string
	// Dong: ex) 역삼동
	Dong string
	// Address: ex) 822-5
	Address string
	// GoogleMapSrc: iframe google map의 src속성 값
	GoogleMapSrc string
}

type Keywords []string

func (k *Keywords) String() string { return strings.Join(*k, ",") }

type Active struct {
	// IsPermanentClosed: 영업중=true 폐업=false
	IsPermanentClosed bool
	// Reason: 폐업상태일 경우에만 입력
	Reason string
}

type TimeType struct {
	// Has: 유무
	Has bool
	// Open: 오픈시간. ex) 18:00
	Open string
	// Closed: 마감시간. ex) 00:00
	Closed string
}

type Hour struct {
	// Part1: 1부
	Part1 *TimeType
	// Part2: 2부
	Part2 *TimeType
}

type Menu struct {
	// Part1Whisky: 1부 주대
	Part1Whisky int
	// Part2Whisky: 2부 주대
	Part2Whisky int
	// TC: 아가씨 티시
	TC int
	// RT: 룸비
	RT int
}

type Store struct {
	Location *Location
	// Type: 업종 하드코딩
	Type string
	// Title: 가게이름 하드코딩
	Title string
	// Description: 가게 설명 하드코딩
	Description string
	// Keywords: 하드코딩 X. 서버 시작시 지역명, 가게이름, 업종 등으로 자동 초기화 됨
	Keywords Keywords
	// Active: 영업, 폐업 유무와 폐업사유 하드코딩
	Active *Active
	// Hour: 영업시간 하드코딩
	Hour *Hour
	// Price: 가격 하드코딩
	Menu *Menu
	// PhoneNumber: 하드코딩 X.
	PhoneNumber string
	// 생성일
	DatePublished time.Time
	// 수정일
	DateModified time.Time
}

func (s *Store) IsModified() bool { return s.DatePublished.UnixNano() != s.DateModified.UnixNano() }

func storeDate(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
}

func initKaraoke() {
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "논현동",
			Address:      "151-30",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3164.8479106529085!2d127.03145169999998!3d37.5115051!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca3f05b7c4407%3A0xbb44e0b5425b8a89!2z7ISc7Jq47Yq567OE7IucIOqwleuCqOq1rCDrhbztmITrj5kgMTUxLTMw!5e0!3m2!1sko!2skr!4v1660745693771!5m2!1sko!2skr",
		},
		Type:        STORE_TYPE_KARAOKE,
		Title:       "퍼펙트",
		Description: "강남 퍼펙트 가라오케에서 최상의 음향 시스템과 다양한 노래 선택으로 당신의 노래 실력을 뽐내보세요. 편안하고 고품격한 분위기에서 친구들과 즐거운 노래 시간을 만끽해보세요. 강남 지역에서 가장 완벽한 가라오케 경험을 제공합니다.",
		Active: &Active{
			IsPermanentClosed: true,
			Reason:            "하이퍼블릭으로 업종 변경",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "01:00"},
			Part2: &TimeType{Has: true, Open: "01:00", Closed: "15:00"},
		},
		Menu:          &Menu{Part1Whisky: 350000, Part2Whisky: 160000, TC: 120000, RT: 50000},
		DatePublished: storeDate(2023, 8, 12),
		DateModified:  storeDate(2023, 8, 12),
	})
}

func initShirtRoom() {
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "삼성동",
			Address:      "142-35",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.1043050533926!2d127.05085469999999!3d37.505458!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca411d5a288d7%3A0xca6681460caa4840!2s411%20Teheran-ro%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1662046616801!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_SHIRTROOM,
		Title:       "디씨",
		Description: "서울의 활기찬 밤문화의 중심, 강남 디씨 셔츠룸에서 독특한 한국 문화와 현대적인 럭셔리가 조화를 이루는 경험을 즐겨보세요. 개인적이고 첨단 기술이 갖춰진 방에서 마음껏 노래하고, 다양한 세대의 멜로디로 친목을 다지며, 고급 한국 요리에 흠뻑 빠져보세요. 강남의 역동적인 셔츠룸 문화에서 음악, 우정, 그리고 문화적 몰입의 잊지 못할 밤을 경험하세요.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "01:00"},
			Part2: &TimeType{Has: true, Open: "01:00", Closed: "15:00"},
		},
		Menu:          &Menu{Part1Whisky: 160000, Part2Whisky: 130000, TC: 120000, RT: 50000},
		DatePublished: storeDate(2023, 12, 1),
		DateModified:  storeDate(2023, 12, 1),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "잠원동",
			Address:      "18-9",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3164.7060647283693!2d127.0171104!3d37.514850200000005!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca3dd364c8bc7%3A0x3ab4d058c71d79a8!2s18-9%20Jamwon-dong%2C%20Seocho-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1670862647642!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_SHIRTROOM,
		Title:       "유앤미",
		Description: "강남 유앤미 셔츠룸에서 세련되고 편안한 분위기의 조화를 경험하세요. 독특한 스타일과 친밀한 환경에서 즐거운 시간을 보낼 수 있는 공간으로, 최신 음악과 다양한 엔터테인먼트를 제공합니다. 친구들과 함께하는 밤이든, 특별한 모임이든, 강남 유앤미 셔츠룸은 모든 순간을 더욱 특별하게 만들어 줄 것입니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "01:00"},
			Part2: &TimeType{Has: true, Open: "01:00", Closed: "15:00"},
		},
		Menu:          &Menu{Part1Whisky: 160000, Part2Whisky: 130000, TC: 120000, RT: 50000},
		DatePublished: storeDate(2023, 12, 2),
		DateModified:  storeDate(2023, 12, 2),
	})
}

func initHighPublic() {
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "역삼동",
			Address:      "604-7",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.088372324827!2d127.0311099!3d37.5058338!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca3fb63865cd7%3A0x31427b556da83644!2s604-7%20Yeoksam-dong%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1662056274810!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_HIGHPUBLIC,
		Title:       "달토",
		Description: "강남 달토 하이퍼블릭에서 활기찬 밤문화와 독특한 엔터테인먼트를 경험하세요. 현대적인 인테리어와 분위기에서 최고의 음악과 함께하는 특별한 밤을 즐기실 수 있습니다. 강남의 중심에서 펼쳐지는 매혹적인 나이트라이프를 달토 하이퍼블릭에서 만나보세요.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "01:00"},
			Part2: &TimeType{Has: true, Open: "01:00", Closed: "15:00"},
		},
		Menu:          &Menu{Part1Whisky: 170000, Part2Whisky: 140000, TC: 110000, RT: 50000},
		DatePublished: storeDate(2023, 12, 3),
		DateModified:  storeDate(2023, 12, 3),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "역삼동",
			Address:      "604-7",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.088372324827!2d127.0311099!3d37.5058338!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca3fb63865cd7%3A0x31427b556da83644!2s604-7%20Yeoksam-dong%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1662056274810!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_HIGHPUBLIC,
		Title:       "런닝래빗",
		Description: "강남 런닝래빗 하이퍼블릭에서 독창적이고 역동적인 밤문화의 매력을 경험하세요. 최신 트렌드와 감각적인 음악이 어우러진 공간에서 친구들과 잊지 못할 추억을 만들어 보세요. 강남의 밤을 화려하게 수놓는 런닝래빗 하이퍼블릭에서 색다른 즐거움을 발견하실 수 있습니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "01:00"},
			Part2: &TimeType{Has: true, Open: "01:00", Closed: "15:00"},
		},
		Menu:          &Menu{Part1Whisky: 170000, Part2Whisky: 140000, TC: 110000, RT: 50000},
		DatePublished: storeDate(2023, 12, 4),
		DateModified:  storeDate(2023, 12, 4),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "역삼동",
			Address:      "831-42",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.6005044881886!2d127.03146729999997!3d37.4937527!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca15057aba5c3%3A0x3c39e1c32ad3bd0f!2s831-42%20Yeoksam-dong%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1665731145337!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_HIGHPUBLIC,
		Title:       "트렌드",
		Description: "강남 트렌드 하이퍼블릭에서 최신 유행과 함께하는 화려한 밤문화를 체험하세요. 트렌디한 음악과 스타일리시한 분위기 속에서 즐거운 시간을 보낼 수 있는 이상적인 장소입니다. 강남의 밤을 색다르게 즐기고 싶다면 트렌드 하이퍼블릭이 정답입니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "01:00"},
			Part2: &TimeType{Has: true, Open: "01:00", Closed: "15:00"},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 110000, RT: 50000},
		DatePublished: storeDate(2023, 12, 5),
		DateModified:  storeDate(2023, 12, 5),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "대치동",
			Address:      "890-38",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.150656732908!2d127.05328440000001!3d37.504364699999996!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca41055280155%3A0xc6516a6b77ef70c1!2z7ISc7Jq47Yq567OE7IucIOqwleuCqOq1rCDrjIDsuZjrj5kgODkwLTM4!5e0!3m2!1sko!2skr!4v1660489421580!5m2!1sko!2skr",
		},
		Type:        STORE_TYPE_HIGHPUBLIC,
		Title:       "사라있네",
		Description: "강남 사라있네 하이퍼블릭에서 독특하고 매력적인 밤문화의 세계로 빠져보세요. 현대적인 디자인과 함께하는 최고의 음악과 엔터테인먼트로 가득 찬 이 곳에서 잊지 못할 밤을 보내실 수 있습니다. 강남의 밤을 더욱 특별하게 만들어 줄 사라있네 하이퍼블릭을 경험해보세요.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "01:00"},
			Part2: &TimeType{Has: true, Open: "01:00", Closed: "15:00"},
		},
		Menu:          &Menu{Part1Whisky: 170000, Part2Whisky: 140000, TC: 110000, RT: 50000},
		DatePublished: storeDate(2023, 12, 6),
		DateModified:  storeDate(2023, 12, 6),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "역삼동",
			Address:      "822-5",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.3849794120856!2d127.02926860000001!3d37.4988373!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca159d7d08f47%3A0x19ac7457d361928!2z7ISc7Jq47Yq567OE7IucIOqwleuCqOq1rCDthYztl6TrnoDroZwgMTEx!5e0!3m2!1sko!2skr!4v1661153125692!5m2!1sko!2skr",
		},
		Type:        STORE_TYPE_HIGHPUBLIC,
		Title:       "메이커",
		Description: "강남 메이커 하이퍼블릭에서 현대적인 스타일과 창의적인 분위기의 조화를 경험하세요. 도시의 중심에서 즐기는 최신 음악과 함께 독특한 엔터테인먼트의 세계로 여러분을 초대합니다. 강남의 밤을 새롭게 정의하는 메이커 하이퍼블릭에서 특별한 시간을 보내세요.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "01:00"},
			Part2: &TimeType{Has: true, Open: "01:00", Closed: "15:00"},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 110000, RT: 50000},
		DatePublished: storeDate(2023, 12, 7),
		DateModified:  storeDate(2023, 12, 7),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "역삼동",
			Address:      "823-30",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.4051104401256!2d127.03307020000001!3d37.4983624!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca1565c22d639%3A0x1fcb22298cd33520!2z7ISc7Jq47Yq567OE7IucIOqwleuCqOq1rCDsl63sgrzrj5kgODIzLTMw!5e0!3m2!1sko!2skr!4v1693829638202!5m2!1sko!2skr",
		},
		Type:        STORE_TYPE_HIGHPUBLIC,
		Title:       "수목원",
		Description: "강남 수목원 하이퍼블릭에서 자연과 어우러진 독특한 밤문화를 즐겨보세요. 푸르른 자연 속에서 현대적인 음악과 엔터테인먼트가 조화를 이루는 이곳에서 특별한 경험을 하실 수 있습니다. 강남에서 자연과 문화가 만나는 수목원 하이퍼블릭으로 여러분을 초대합니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "01:00"},
			Part2: &TimeType{Has: true, Open: "01:00", Closed: "15:00"},
		},
		Menu:          &Menu{Part1Whisky: 250000, Part2Whisky: 0, TC: 130000, RT: 50000},
		DatePublished: storeDate(2023, 12, 8),
		DateModified:  storeDate(2023, 12, 8),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "논현동",
			Address:      "151-30",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3164.8479106529085!2d127.03145169999998!3d37.5115051!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca3f05b7c4407%3A0xbb44e0b5425b8a89!2z7ISc7Jq47Yq567OE7IucIOqwleuCqOq1rCDrhbztmITrj5kgMTUxLTMw!5e0!3m2!1sko!2skr!4v1660745693771!5m2!1sko!2skr",
		},
		Type:        STORE_TYPE_HIGHPUBLIC,
		Title:       "퍼펙트",
		Description: "강남 퍼펙트 하이퍼블릭에서 완벽한 밤문화 경험을 하세요. 현대적인 감각과 최고급 서비스가 만나는 곳에서 멋진 음악과 함께하는 환상적인 밤을 보낼 수 있습니다. 강남의 매력적인 밤을 퍼펙트 하이퍼블릭과 함께 즐겨보세요.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "01:00"},
			Part2: &TimeType{Has: true, Open: "01:00", Closed: "15:00"},
		},
		Menu:          &Menu{Part1Whisky: 350000, Part2Whisky: 160000, TC: 110000, RT: 50000},
		DatePublished: storeDate(2023, 12, 9),
		DateModified:  storeDate(2023, 12, 9),
	})
}

func initLeggingsRoom() {
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "삼성동",
			Address:      "144-10",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.0368804441946!2d127.0548939!3d37.5070483!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca413ea3ed99f%3A0xdd0a3d80af8a9047!2s144-10%20Samseong-dong%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1662646930422!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_LEGGINGS,
		Title:       "하이킥",
		Description: "강남 하이킥 레깅스룸에서 역동적이고 혁신적인 밤문화를 경험해보세요. 세련된 인테리어와 함께 최신 음악이 가득한 이곳에서 흥겨운 시간을 보낼 수 있습니다. 강남에서 색다른 즐거움을 찾는다면 하이킥 레깅스룸이 완벽한 선택입니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "01:00"},
			Part2: &TimeType{Has: true, Open: "01:00", Closed: "15:00"},
		},
		Menu:          &Menu{Part1Whisky: 250000, Part2Whisky: 0, TC: 150000, RT: 50000},
		DatePublished: storeDate(2023, 12, 10),
		DateModified:  storeDate(2023, 12, 10),
	})
}

func initDot5() {
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "논현동",
			Address:      "204-4",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.069981650343!2d127.02487893188555!3d37.50626757076464!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca3fb554ff02b%3A0x8d9e573a46ec1b7a!2s204-4%20Nonhyeon-dong%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1679716196560!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_DOT5,
		Title:       "유니크",
		Description: "강남 유니크 쩜오에서 개성 넘치는 밤문화의 진수를 경험해보세요. 독특한 테마와 함께하는 다채로운 엔터테인먼트가 여러분을 기다리고 있습니다. 강남에서만 느낄 수 있는 유니크한 매력과 함께 즐거운 밤을 보내세요.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2023, 12, 1),
		DateModified:  storeDate(2023, 12, 1),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "역삼동",
			Address:      "831",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.5641798788934!2d127.0297203!3d37.4946097!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca1508715f00d%3A0xf4d079a0f225c1b1!2s831%20Yeoksam-dong%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1679397724056!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_DOT5,
		Title:       "831",
		Description: "강남 831 쩜오에서 독특하고 현대적인 밤문화를 체험하세요. 트렌디한 음악과 함께하는 매혹적인 분위기에서 즐겁고 흥미진진한 시간을 보낼 수 있습니다. 강남의 밤을 새롭게 정의하는 831 쩜오에서 특별한 경험을 해보세요.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2023, 12, 2),
		DateModified:  storeDate(2023, 12, 2),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "역삼동",
			Address:      "735-32",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.3760308056935!2d127.0341289!3d37.4990484!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca1560e5d6327%3A0x5c114aeb8260a643!2s735-32%20Yeoksam-dong%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1679397562888!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_DOT5,
		Title:       "에이원",
		Description: "강남 에이원 쩜오에서 현대적인 감각과 트렌디한 스타일이 만나는 독창적인 밤문화를 경험하세요. 최고의 음악과 엔터테인먼트가 어우러진 이곳에서 잊을 수 없는 밤을 보내실 수 있습니다. 강남의 밤을 새로운 시선으로 즐기고 싶다면 에이원이 제공하는 특별한 경험을 놓치지 마세요.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2023, 12, 3),
		DateModified:  storeDate(2023, 12, 3),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "삼성동",
			Address:      "141-33",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.121539211326!2d127.04949690000001!3d37.5050515!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca40fc775ade5%3A0xdd9b10797e776ad1!2s141-33%20Samseong-dong%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1678667592079!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_DOT5,
		Title:       "미라클",
		Description: "강남 미라클 쩜오에서 기적 같은 밤문화를 경험해보세요. 활기찬 음악과 함께하는 매혹적인 분위기에서 독특한 엔터테인먼트를 즐길 수 있습니다. 강남에서만 누릴 수 있는 미라클 쩜오의 특별한 밤을 체험하세요.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2023, 12, 4),
		DateModified:  storeDate(2023, 12, 4),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "역삼동",
			Address:      "701-2",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.204884148887!2d127.0430503!3d37.5030856!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca406fc7ff209%3A0x341d4adf49840962!2s701-2%20Yeoksam-dong%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1678667437305!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_DOT5,
		Title:       "오키도키",
		Description: "강남 오키도키 쩜오에서 즐거움과 활기가 넘치는 밤문화를 만끽하세요. 현대적인 분위기와 트렌디한 음악이 조화를 이루는 공간에서 잊을 수 없는 밤을 보내실 수 있습니다. 강남의 밤을 즐겁고 신나게 만들어줄 오키도키 쩜오에서 특별한 경험을 해보세요.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2023, 12, 5),
		DateModified:  storeDate(2023, 12, 5),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "신사동",
			Address:      "561-30",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3164.5256239750274!2d127.0258308!3d37.5191051!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca3ecf7b91b35%3A0x90e6eb4e73a5644e!2s561-30%20Sinsa-dong%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1678606137375!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_DOT5,
		Title:       "인트로",
		Description: "강남 인트로 쩜오에서 시작되는 환상적인 밤문화를 경험하세요. 현대적인 스타일과 함께 트렌디한 음악이 이끄는 멋진 분위기에서 특별한 시간을 보내보세요. 강남의 밤을 새롭고 흥미로운 방식으로 시작하고 싶다면 인트로 쩜오가 완벽한 장소입니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2023, 12, 5),
		DateModified:  storeDate(2023, 12, 5),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "논현동",
			Address:      "248-7",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3164.741047626004!2d127.03369181564705!3d37.51402523489071!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca3f415b07255%3A0x2162a0d614d3c110!2s640%20Eonju-ro%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1678605759071!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_DOT5,
		Title:       "머니볼",
		Description: "강남 머니볼 쩜오에서 화려하고 역동적인 밤문화를 체험하세요. 최신 트렌드를 반영한 음악과 스타일로 가득 찬 이곳에서 잊지 못할 밤을 보내실 수 있습니다. 강남에서 만날 수 있는 머니볼 쩜오의 특별한 매력을 경험해보세요.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2023, 12, 6),
		DateModified:  storeDate(2023, 12, 6),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "역삼동",
			Address:      "731-11",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.401460674176!2d127.0436794!3d37.498448499999995!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca401a6b8183b%3A0xcbcd58a8b2cb7c50!2s731%20Yeoksam-dong%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1678605118720!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_DOT5,
		Title:       "라이징",
		Description: "강남 라이징 쩜오에서 역동적이고 신나는 밤문화의 정수를 체험해보세요. 현대적인 감각과 생동감 넘치는 음악이 만나는 곳에서 새로운 밤문화를 경험할 수 있습니다. 강남에서 특별한 밤을 원한다면 라이징 쩜오가 제공하는 활기찬 경험을 놓치지 마세요.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2023, 12, 7),
		DateModified:  storeDate(2023, 12, 7),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "역삼동",
			Address:      "736-17",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.3715755621406!2d127.03453809999999!3d37.4991535!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca15607cff005%3A0x9a314c8436603f9e!2s736-17%20Yeoksam-dong%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1677802895674!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_DOT5,
		Title:       "임팩트",
		Description: "강남 임팩트 쩜오에서 강렬하고 독특한 밤문화를 경험해보세요. 혁신적인 음악과 함께하는 역동적인 분위기 속에서 잊을 수 없는 밤을 보내실 수 있습니다. 강남에서 찾을 수 있는 임팩트 쩜오의 매력적인 경험을 놓치지 마세요.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2023, 12, 8),
		DateModified:  storeDate(2023, 12, 8),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "역삼동",
			Address:      "824-7",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.429128349951!2d127.03037690000001!3d37.4977958!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca1576a139921%3A0xda0428a0d46a18b2!2s824-7%20Yeoksam-dong%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1676634190100!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_DOT5,
		Title:       "스테이",
		Description: "강남 스테이 쩜오에서 편안하면서도 매력적인 밤문화를 즐겨보세요. 모던하고 아늑한 분위기에서 최신 음악과 함께 휴식과 즐거움을 동시에 경험할 수 있습니다. 강남의 밤을 잔잔하고 특별하게 만들어줄 스테이 쩜오에서 멋진 시간을 보내세요.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2023, 12, 9),
		DateModified:  storeDate(2023, 12, 9),
	})
}

func initClub() {
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "도산대로",
			Address:      "114",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3164.64159846425!2d127.02127!3d37.5163704!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca3e9a9f07727%3A0x4fcde2f83452e564!2s114%20Dosan-daero%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1681189780772!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_CLUB,
		Title:       "사운드",
		Description: "강남 사운드 클럽에서 최신 음악과 함께하는 고급스러운 밤문화를 경험해보세요. 화려한 조명과 최고 수준의 사운드 시스템이 조화를 이루는 이곳에서는 다양한 장르의 음악과 DJ 공연을 즐길 수 있습니다. 강남의 밤을 대표하는 이 클럽에서 특별한 날을 보내세요.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "23:00", Closed: "11:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2023, 12, 10),
		DateModified:  storeDate(2023, 12, 10),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "잠원동",
			Address:      "21-3",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3164.6962477822185!2d127.0192326!3d37.51508169999999!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca3e80fe94731%3A0xadedf946e74c560c!2s21-3%20Jamwon-dong%2C%20Seocho-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1681189358457!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_CLUB,
		Title:       "레이스",
		Description: "강남 레이스 클럽에서 화려하고 역동적인 밤문화를 경험하세요. 현대적인 사운드와 다채로운 음악 장르의 향연이 펼쳐지는 이곳에서 새로운 밤의 시작을 즐겨보세요. 강남의 중심에서 펼쳐지는 최고의 DJ 공연과 화려한 이벤트를 레이스 클럽에서 만나보세요.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "23:00", Closed: "11:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2023, 12, 11),
		DateModified:  storeDate(2023, 12, 11),
	})
}

func initHobba() {
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "삼성동",
			Address:      "143-35",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.0622628938677!2d127.05028567647602!3d37.5064496275705!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca4118576f5e1%3A0xbc745a3337004851!2s143-35%20Samseong-dong%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1685329649613!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_HOBBA,
		Title:       "어게인",
		Description: "강남 어게인 호빠는 강남에서 인기 있는 호스트바 중 하나로, 여러 호스트바를 통합하여 새롭게 리뉴얼한 곳입니다. 이곳은 다양한 종류의 호빠 서비스를 제공하며, 주로 여성 고객을 대상으로 하는 정빠 서비스와 남성 및 여성 모두를 대상으로 하는 중빠, 준빠 서비스 등을 포함합니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "15:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 180000, Part2Whisky: 0, TC: 60000, RT: 50000},
		DatePublished: storeDate(2023, 12, 12),
		DateModified:  storeDate(2023, 12, 12),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "삼성동",
			Address:      "143-27",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.0354982629583!2d127.0543849!3d37.5070809!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca413c457ed95%3A0x2c8f79900d733d24!2s143-27%20Samseong-dong%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1685329268008!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_HOBBA,
		Title:       "씨엔엔",
		Description: "강남 씨엔엔 호빠는 강남 지역에서 유명한 호스트바 중 하나로, 예술과 엔터테인먼트가 어우러진 독특한 경험을 제공하는 곳입니다. 씨엔엔 호빠의 위치는 서울 강남구 삼성동 143-27로, 쉽게 접근할 수 있는 곳에 위치해 있습니다. 이곳은 고급스러운 인테리어와 함께 다양한 종류의 주류 및 음식을 즐길 수 있는 곳 입니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "15:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 180000, Part2Whisky: 0, TC: 60000, RT: 50000},
		DatePublished: storeDate(2023, 12, 13),
		DateModified:  storeDate(2023, 12, 13),
	})
}

func initFull() {
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "역삼동",
			Address:      "718-14",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.303372126534!2d127.0393423!3d37.5007624!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca3ff706eedd3%3A0xfcdcb626028ce68f!2z7ISc7Jq47Yq567OE7IucIOqwleuCqOq1rCDsl63sgrzrj5kgNzE4LTE0!5e0!3m2!1sko!2skr!4v1702894200532!5m2!1sko!2skr",
		},
		Type:        STORE_TYPE_FULL,
		Title:       "세븐",
		Description: "강남 세븐 풀싸롱은 서울 강남구에 위치한 정통 룸시스템을 제공하는 소프트 풀싸롱입니다. 이곳은 다양한 최상급 미모의 여매니저들을 선택할 수 있는 룸초이스 방식을 채택하고 있으며, 비즈니스 미팅, 중요한 접대, 총각파티, 생일파티, 친목모임 등 다양한 자리에 적합합니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "17:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 400000, RT: 0},
		DatePublished: storeDate(2023, 12, 14),
		DateModified:  storeDate(2023, 12, 14),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "역삼동",
			Address:      "677-19",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.2213246096076!2d127.0403477!3d37.5026978!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca40757253203%3A0x71ccce464c9712c7!2z7ISc7Jq47Yq567OE7IucIOqwleuCqOq1rCDsl63sgrzrj5kgNjc3LTE5!5e0!3m2!1sko!2skr!4v1702894590497!5m2!1sko!2skr",
		},
		Type:        STORE_TYPE_FULL,
		Title:       "심포니",
		Description: "강남 심포니 풀싸롱에서 고급스러움과 세련됨이 어우러진 독특한 밤문화를 경험해보세요. 우아하고 편안한 분위기 속에서 프리미엄 서비스를 즐길 수 있는 이곳은 강남의 중심에서 품격 있는 여가 시간을 제공합니다. 최상의 매니저들과 함께하는 풍성한 선택의 폭, 개인화된 주문형 서비스로 각 고객의 취향을 만족시킬 준비가 되어 있습니다. 강남 심포니 풀싸롱은 여러분의 특별한 밤을 위한 완벽한 선택지입니다. 탁월한 호스피탈리티와 함께하는 최고급 경험을 강남 심포니 풀싸롱에서 만나보세요.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "17:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 400000, RT: 0},
		DatePublished: storeDate(2023, 12, 15),
		DateModified:  storeDate(2023, 12, 15),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "역삼동",
			Address:      "719-18",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.279012273408!2d127.03845047644623!3d37.50133702786316!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca400ba25b03f%3A0x303ca4f57a0d74bd!2z7ISc7Jq47Yq567OE7IucIOqwleuCqOq1rCDsl63sgrzrj5kgNzE5LTE4!5e0!3m2!1sko!2skr!4v1702894960266!5m2!1sko!2skr",
		},
		Type:        STORE_TYPE_FULL,
		Title:       "애플",
		Description: "강남 애플 풀싸롱에서 고급스러운 룸서비스와 매력적인 매니저들과의 만남을 경험하세요. 개인화된 서비스와 다양한 엔터테인먼트 옵션을 통해 강남의 밤문화를 새로운 차원에서 즐길 수 있습니다. 고객 맞춤형 경험을 제공하는 강남 애플 풀싸롱에서 화려한 밤을 보내세요.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "17:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 400000, RT: 0},
		DatePublished: storeDate(2023, 12, 16),
		DateModified:  storeDate(2023, 12, 16),
	})
}

func setStoreKeywords() {
	for _, s := range stores {
		s.Keywords = Keywords([]string{
			fmt.Sprintf("%s %s %s %s %s", s.Location.Do, s.Location.Si, s.Location.Dong, s.Type, s.Title),
			fmt.Sprintf("%s %s", s.Location.Do, s.Type),
			fmt.Sprintf("%s %s %s", s.Location.Do, s.Location.Si, s.Type),
			fmt.Sprintf("%s %s %s %s", s.Location.Do, s.Location.Si, s.Location.Dong, s.Type),
			fmt.Sprintf("%s %s", s.Title, s.Type),
			fmt.Sprintf("%s %s 가격", s.Title, s.Type),
			fmt.Sprintf("%s %s 시스템", s.Title, s.Type),
			fmt.Sprintf("%s %s 주소", s.Title, s.Type),
		})
	}
}

func setPhoneNumbers() {
	for _, s := range stores {
		//		switch s.Type {
		//		case STORE_TYPE_DOT5:
		//			s.PhoneNumber = "010-2170-4981"
		//		case STORE_TYPE_CLUB, STORE_TYPE_HOBBA:
		//			s.PhoneNumber = "010-6590-7589"
		//		default:
		//			s.PhoneNumber = site.Config.PhoneNumber
		//		}
		// 풀싸 폰번호로만
		s.PhoneNumber = site.Config.PhoneNumber
	}
}

func sortStores() {
	sort.Slice(stores, func(i, j int) bool {
		return stores[i].DatePublished.UnixNano() < stores[j].DatePublished.UnixNano()
	})
}

// 서버 시작시 vieiws/store directories 자동 생성
func createViewsDirectories() error {
	for _, s := range stores {
		dir := fmt.Sprintf("views/store/%s/%s/%s/%s",
			s.Location.Do, s.Location.Si, s.Location.Dong, s.Type)
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

// 서버 시작시 views/store/../../{{store.Title}}.html 파일 자동 생성
func createHTMLFiles() error {
	for _, s := range stores {
		filepath := fmt.Sprintf("views/store/%s/%s/%s/%s/%s.html",
			s.Location.Do, s.Location.Si, s.Location.Dong, s.Type, s.Title)
		if _, err := os.Stat(filepath); err == nil {
			continue
		}
		if err := os.WriteFile(filepath, []byte("write me!"), os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

// 서버 시작시 store 이미지 디렉토리 자동 생성
func createStaticImgDirectories() error {
	for _, s := range stores {
		dir := fmt.Sprintf("static/img/store/%s/%s/%s/%s/%s",
			s.Location.Do, s.Location.Si, s.Location.Dong, s.Type, s.Title)
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

func Init() error {
	//initKaraoke()
	initShirtRoom()
	initHighPublic()
	initLeggingsRoom()
	initDot5()
	initClub()
	initHobba()
	initFull()

	sortStores()

	setStoreKeywords()
	setPhoneNumbers()

	if err := createViewsDirectories(); err != nil {
		return err
	}
	if err := createHTMLFiles(); err != nil {
		return err
	}
	if err := createStaticImgDirectories(); err != nil {
		return err
	}
	return nil
}
