package store

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/jeonghoikun/jinwoowide.com/site"
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
			Dong:         "삼성동",
			Address:      "142-35",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.1043050533926!2d127.05085469999999!3d37.505458!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca411d5a288d7%3A0xca6681460caa4840!2s411%20Teheran-ro%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1662046616801!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_KARAOKE,
		Title:       "파티원",
		Description: "강남 파티원 가라오케에서는 화려한 분위기와 최상의 서비스로 당신의 밤을 뜨겁게 만들어 드립니다. 최고의 음향 시설과 환상적인 라이트 쇼로 최상의 파티 경험을 제공합니다.	",
		Active: &Active{
			IsPermanentClosed: true,
			Reason:            "정상폐업",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "01:00"},
			Part2: &TimeType{Has: true, Open: "01:00", Closed: "15:00"},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 150000, TC: 100000, RT: 50000},
		DatePublished: storeDate(2024, 2, 18),
		DateModified:  storeDate(2024, 2, 18),
	})
}

func initShirtRoom() {
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
		Description: "강남 유앤미 셔츠룸은 고급스러운 인테리어와 프라이빗한 분위기에서 여러분을 맞이합니다. 세련된 셔츠 차림의 서비스로 특별한 밤을 보낼 수 있습니다.",
		Active: &Active{
			IsPermanentClosed: true,
			Reason:            "하이퍼블릭으로 업종 변경",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "01:00"},
			Part2: &TimeType{Has: true, Open: "01:00", Closed: "15:00"},
		},
		Menu:          &Menu{Part1Whisky: 160000, Part2Whisky: 130000, TC: 120000, RT: 50000},
		DatePublished: storeDate(2024, 1, 10),
		DateModified:  storeDate(2024, 6, 4),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "삼성동",
			Address:      "143-27",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.0354982629583!2d127.0543849!3d37.5070809!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca413c457ed95%3A0x2c8f79900d733d24!2s143-27%20Samseong-dong%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1685329268008!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_SHIRTROOM,
		Title:       "씨엔엔",
		Description: "강남 씨엔엔 셔츠룸에서는 독특한 테마와 함께 색다른 경험을 제공합니다. 친절한 서비스와 함께 기억에 남는 시간을 보내실 수 있습니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "15:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 180000, Part2Whisky: 0, TC: 60000, RT: 50000},
		DatePublished: storeDate(2024, 2, 15),
		DateModified:  storeDate(2024, 2, 15),
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
		Description: "강남 달토 하이퍼블릭은 차별화된 음료와 함께 현대적인 분위기에서 즐길 수 있는 최고의 장소입니다. 도심 속에서 특별한 밤을 보내고 싶다면 이곳이 정답입니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "01:00"},
			Part2: &TimeType{Has: true, Open: "01:00", Closed: "15:00"},
		},
		Menu:          &Menu{Part1Whisky: 170000, Part2Whisky: 150000, TC: 120000, RT: 50000},
		DatePublished: storeDate(2024, 1, 11),
		DateModified:  storeDate(2024, 1, 11),
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
		Description: "강남 런닝래빗 하이퍼블릭은 독특한 콘셉트와 활기찬 분위기로 여러분의 밤을 즐겁게 만들어 줄 곳입니다. 친구들과 함께라면 더욱 즐거운 시간을 보낼 수 있습니다.	",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "01:00"},
			Part2: &TimeType{Has: true, Open: "01:00", Closed: "15:00"},
		},
		Menu:          &Menu{Part1Whisky: 170000, Part2Whisky: 150000, TC: 120000, RT: 50000},
		DatePublished: storeDate(2024, 1, 11),
		DateModified:  storeDate(2024, 1, 11),
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
		Description: "강남 트렌드 하이퍼블릭에서는 최신 유행을 선도하는 멋진 공간에서 독보적인 밤문화를 경험할 수 있습니다. 모던하고 스타일리시한 분위기가 여러분을 기다립니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "01:00"},
			Part2: &TimeType{Has: true, Open: "01:00", Closed: "15:00"},
		},
		Menu:          &Menu{Part1Whisky: 200000, Part2Whisky: 0, TC: 120000, RT: 50000},
		DatePublished: storeDate(2024, 1, 12),
		DateModified:  storeDate(2024, 1, 12),
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
		Description: "강남 사라있네 하이퍼블릭은 이색적인 분위기와 함께 잊을 수 없는 밤을 선사합니다. 개성 넘치는 인테리어와 뛰어난 서비스로 당신의 밤을 특별하게 만들어 드립니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "01:00"},
			Part2: &TimeType{Has: true, Open: "01:00", Closed: "15:00"},
		},
		Menu:          &Menu{Part1Whisky: 170000, Part2Whisky: 150000, TC: 120000, RT: 50000},
		DatePublished: storeDate(2024, 1, 12),
		DateModified:  storeDate(2024, 1, 12),
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
		Description: "강남 수목원 하이퍼블릭에서는 자연 속 휴식 같은 평온함과 함께 활기를 느낄 수 있습니다. 도심 속 오아시스에서 특별한 시간을 보내보세요.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "01:00"},
			Part2: &TimeType{Has: true, Open: "01:00", Closed: "15:00"},
		},
		Menu:          &Menu{Part1Whisky: 250000, Part2Whisky: 0, TC: 130000, RT: 50000},
		DatePublished: storeDate(2024, 1, 12),
		DateModified:  storeDate(2024, 1, 12),
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
		Description: "강남 퍼펙트 하이퍼블릭은 완벽한 밤을 위한 최적의 장소입니다. 고급스러운 분위기와 탁월한 서비스로 여러분의 기대를 충족시켜 드립니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "01:00"},
			Part2: &TimeType{Has: true, Open: "01:00", Closed: "15:00"},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 150000, TC: 120000, RT: 50000},
		DatePublished: storeDate(2024, 1, 13),
		DateModified:  storeDate(2024, 1, 13),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "역삼동",
			Address:      "824-8",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.434083647925!2d127.0305156!3d37.4976789!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca15741b03c33%3A0xf28611c1cfc94af5!2z7ISc7Jq47Yq567OE7IucIOqwleuCqOq1rCDsl63sgrzrj5kgODI0LTg!5e0!3m2!1sko!2skr!4v1704324043037!5m2!1sko!2skr",
		},
		Type:        STORE_TYPE_HIGHPUBLIC,
		Title:       "워라벨",
		Description: "강남 워라벨 하이퍼블릭은 일과 삶의 균형을 중요시하는 분들을 위한 이상적인 공간입니다. 일상에서 벗어나 재충전할 수 있는 편안하고 즐거운 분위기를 제공합니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "01:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 200000, Part2Whisky: 0, TC: 120000, RT: 50000},
		DatePublished: storeDate(2024, 1, 13),
		DateModified:  storeDate(2024, 1, 13),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "역삼동",
			Address:      "832-7",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.679446741483!2d127.02837221193238!3d37.49189017194145!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca1502738de7b%3A0x65a8ee648278baf2!2z7ISc7Jq47Yq567OE7IucIOqwleuCqOq1rCDsl63sgrzrj5kgODMyLTc!5e0!3m2!1sko!2skr!4v1704324092279!5m2!1sko!2skr",
		},
		Type:        STORE_TYPE_HIGHPUBLIC,
		Title:       "방탄",
		Description: "강남 방탄 하이퍼블릭은 강인함과 활력이 넘치는 공간에서 역동적인 밤문화를 즐길 수 있는 곳입니다. 안전하고 활기찬 분위기에서 잊지 못할 경험을 할 수 있습니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "01:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 250000, Part2Whisky: 0, TC: 130000, RT: 50000},
		DatePublished: storeDate(2024, 1, 14),
		DateModified:  storeDate(2024, 1, 14),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "잠원동",
			Address:      "18-9",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3164.7060647283693!2d127.0171104!3d37.514850200000005!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca3dd364c8bc7%3A0x3ab4d058c71d79a8!2s18-9%20Jamwon-dong%2C%20Seocho-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1670862647642!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_HIGHPUBLIC,
		Title:       "유앤미",
		Description: "강남 유앤미 하이퍼블릭은 고급스러운 인테리어와 최신 음향 시스템을 갖춘 유흥 명소입니다. 다양한 엔터테인먼트와 프리미엄 음료로 특별한 밤을 만들어보세요. 지금 예약하세요!",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "01:00"},
			Part2: &TimeType{Has: true, Open: "01:00", Closed: "15:00"},
		},
		Menu:          &Menu{Part1Whisky: 170000, Part2Whisky: 150000, TC: 120000, RT: 50000},
		DatePublished: storeDate(2024, 6, 4),
		DateModified:  storeDate(2024, 6, 4),
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
		Description: "강남 하이킥 레깅스룸에서는 스포티한 매력과 함께 활발한 밤을 보낼 수 있습니다. 역동적인 음악과 함께 색다른 레깅스룸 문화를 경험해 보세요.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "01:00"},
			Part2: &TimeType{Has: true, Open: "01:00", Closed: "15:00"},
		},
		Menu:          &Menu{Part1Whisky: 250000, Part2Whisky: 0, TC: 150000, RT: 50000},
		DatePublished: storeDate(2024, 1, 14),
		DateModified:  storeDate(2024, 1, 14),
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
		Description: "강남 유니크 쩜오는 그 이름처럼 독특하고 창의적인 분위기에서 여러분을 맞이합니다. 개성 넘치는 인테리어와 특별한 서비스로 당신만의 밤을 만들어 드립니다.	",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2024, 1, 15),
		DateModified:  storeDate(2024, 1, 15),
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
		Description: "강남 831 쩜오는 모던하고 세련된 디자인으로 고급스러운 밤문화의 새로운 기준을 제시합니다. 매력적인 공간에서 고급스러운 시간을 보낼 수 있습니다.	",
		Active: &Active{
			IsPermanentClosed: true,
			Reason:            "썸데이로 상호 변경",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2024, 1, 15),
		DateModified:  storeDate(2024, 9, 13),
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
		Description: "강남 에이원 쩜오는 최상의 서비스와 함께 럭셔리한 경험을 제공하는 곳입니다. 고품격 분위기에서 까다로운 취향을 만족시킬 수 있는 프리미엄 장소입니다.	",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2024, 1, 15),
		DateModified:  storeDate(2024, 1, 15),
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
		Description: "강남 미라클 쩜오에서는 매일이 기적 같은 밤을 선사합니다. 환상적인 분위기와 함께 최상의 서비스로 당신의 밤을 더욱 빛나게 만들어 드립니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2024, 1, 16),
		DateModified:  storeDate(2024, 1, 16),
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
		Description: "강남 오키도키 쩜오는 즐거움이 가득한 곳으로, 긍정적인 에너지와 함께 잊을 수 없는 시간을 보낼 수 있습니다. 밝고 활기찬 분위기에서 특별한 밤을 경험하세요.",
		Active: &Active{
			IsPermanentClosed: true,
			Reason:            "더글로리로 상호 변경",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2024, 1, 16),
		DateModified:  storeDate(2024, 4, 27),
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
		Description: "강남 인트로 쩜오는 새로운 시작을 알리는 곳입니다. 독특한 컨셉과 멋진 인테리어로 첫인상부터 강렬한 인상을 남깁니다. 새로운 만남과 경험을 원한다면 이곳이 정답입니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2024, 1, 16),
		DateModified:  storeDate(2024, 1, 16),
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
		Description: "강남 머니볼 쩜오는 성공과 번영의 상징처럼, 여러분의 밤을 화려하게 장식해 줄 곳입니다. 고급스러움과 함께 독특한 경험을 선사하여 당신의 밤을 더욱 특별하게 만들어 줍니다.",
		Active: &Active{
			IsPermanentClosed: true,
			Reason:            "멀리건, 알파벳으로 상호 변경",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2024, 1, 16),
		DateModified:  storeDate(2024, 4, 27),
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
		Description: "강남 라이징 쩜오는 계속해서 상승하는 분위기 속에서 끝없는 즐거움을 제공합니다. 역동적이고 짜릿한 경험을 통해 여러분의 밤을 더욱 열정적으로 만들어 드립니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2024, 1, 17),
		DateModified:  storeDate(2024, 1, 17),
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
		Description: "강남 임팩트 쩜오에서는 강렬한 인상을 남기는 밤을 경험할 수 있습니다. 독특한 컨셉과 역동적인 분위기로 방문하는 모든 이에게 잊지 못할 추억을 선사합니다.	",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2024, 1, 17),
		DateModified:  storeDate(2024, 1, 17),
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
		Description: "강남 스테이 쩜오는 여러분이 머무르고 싶어 할 만큼 매력적인 곳입니다. 편안하면서도 세련된 분위기에서 최상의 서비스를 제공하여, 여러분의 밤을 완벽하게 만들어 드립니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2024, 1, 17),
		DateModified:  storeDate(2024, 1, 17),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "역삼동",
			Address:      "677-22",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.2307457193765!2d127.03704181193267!3d37.50247557193869!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca3f8acb4cd37%3A0xa46ef02bf086e82c!2z7ISc7Jq47Yq567OE7IucIOqwleuCqOq1rCDsl63sgrzrj5kgNjc3LTIy!5e0!3m2!1sko!2skr!4v1704324149895!5m2!1sko!2skr",
		},
		Type:        STORE_TYPE_DOT5,
		Title:       "킹스맨",
		Description: "강남 킹스맨 쩜오는 고급스러움과 우아함이 어우러진 곳으로, 군더더기 없는 세련된 서비스로 귀하를 맞이합니다. 진정한 왕처럼 특별한 대우를 받을 수 있는 장소입니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2024, 1, 18),
		DateModified:  storeDate(2024, 1, 18),
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
		Title:       "더글로리",
		Description: "강남 더글로리 쩜오에서는 최고의 럭셔리와 프리미엄 서비스를 경험할 수 있습니다. 세련된 분위기와 맞춤형 서비스로 모든 순간을 특별하게 만들어줍니다. 강남에서의 완벽한 선택입니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2024, 3, 23),
		DateModified:  storeDate(2024, 3, 23),
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
		Title:       "멀리건",
		Description: "강남 멀리건 쩜오에서 최고급 칵테일과 전문 바텐더의 서비스를 경험하세요. 화려한 라이브 음악과 함께 강남의 밤을 매혹적으로 만들어 드립니다!",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2024, 4, 27),
		DateModified:  storeDate(2024, 4, 27),
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
		Title:       "알파벳",
		Description: "강남 알파벳 쩜오에서는 현대적인 분위기, 전문 바텐더가 제조한 맞춤형 칵테일, 그리고 다양한 엔터테인먼트로 특별한 밤을 보낼 수 있습니다. 강남의 밤을 즐겨보세요!",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2024, 4, 27),
		DateModified:  storeDate(2024, 4, 27),
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
		Title:       "썸데이",
		Description: "강남 썸데이 쩜오는 고급스러운 인테리어와 세련된 분위기 속에서 다양한 프리미엄 음료를 제공하는 유흥주점입니다. 전문 바텐더가 선보이는 독창적인 칵테일과 주류는 고객의 취향을 만족시키며, 프라이빗한 공간에서 편안하고 특별한 시간을 보낼 수 있습니다. 강남의 밤을 더욱 빛나게 할 특별한 경험을 원하는 이들에게 최적의 장소입니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2024, 9, 13),
		DateModified:  storeDate(2024, 9, 13),
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
		Description: "강남 사운드 클럽에서는 최고의 사운드 시스템과 함께 역동적인 밤을 즐길 수 있습니다. 최신 트렌드의 음악과 함께 활기찬 파티를 경험하세요.	",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "23:00", Closed: "11:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2024, 1, 19),
		DateModified:  storeDate(2024, 1, 19),
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
		Description: "강남 레이스 클럽은 스피드와 열정이 넘치는 곳으로, 빠르게 변화하는 트렌드를 따라잡을 수 있는 최적의 장소입니다. 역동적인 분위기에서 즐거운 시간을 보낼 수 있습니다.	",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "23:00", Closed: "11:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2024, 1, 19),
		DateModified:  storeDate(2024, 1, 19),
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
		Description: "강남 어게인 호빠는 한 번 더 찾고 싶어지는 매력적인 곳입니다. 친근하고 활기찬 분위기 속에서 즐거운 대화와 함께 즐길 수 있는 장소로, 재방문을 유도하는 매력이 있습니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "15:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 180000, Part2Whisky: 0, TC: 60000, RT: 50000},
		DatePublished: storeDate(2024, 1, 20),
		DateModified:  storeDate(2024, 1, 20),
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
		Description: "강남 세븐 풀싸롱에서는 럭셔리한 분위기와 함께 일곱 가지의 특별한 경험을 제공합니다. 고급스러운 인테리어와 최상의 서비스로 당신의 밤을 빛내 줄 것입니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "17:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 400000, RT: 0},
		DatePublished: storeDate(2024, 1, 21),
		DateModified:  storeDate(2024, 1, 21),
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
		Description: "강남 심포니 풀싸롱은 우아하고 고급스러운 음악처럼 여러분의 밤을 아름답게 장식해 줄 곳입니다. 조화로운 분위기와 함께 환상적인 시간을 보낼 수 있습니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "17:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 400000, RT: 0},
		DatePublished: storeDate(2024, 1, 22),
		DateModified:  storeDate(2024, 1, 22),
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
		Description: "강남 애플 풀싸롱은 신선하고 활기찬 분위기로 여러분을 맞이합니다. 새로움과 재미가 가득한 곳에서 높은 퀄리티의 서비스와 함께 즐거운 밤을 보낼 수 있습니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "17:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 400000, RT: 0},
		DatePublished: storeDate(2024, 1, 23),
		DateModified:  storeDate(2024, 1, 23),
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
	initKaraoke()
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
