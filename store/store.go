package store

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/jeonghoikun/woorifull.com/site"
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
		Description: "강남 디씨 셔츠룸은 강남의 활기찬 유흥주점 중에서도 고급스러움과 프라이버시를 강조하는 곳으로 알려져 있으며, 비즈니스 미팅이나 친구들과의 사적인 모임에 이상적인 고급 서비스를 제공하여 각 방문객에게 맞춤형 즐거움을 선사합니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "01:00"},
			Part2: &TimeType{Has: true, Open: "01:00", Closed: "15:00"},
		},
		Menu:          &Menu{Part1Whisky: 160000, Part2Whisky: 130000, TC: 120000, RT: 50000},
		DatePublished: storeDate(2024, 1, 10),
		DateModified:  storeDate(2024, 1, 10),
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
		Description: "강남 유앤미 셔츠룸은 우아한 인테리어와 뛰어난 서비스로 유흥주점을 방문하는 이들에게 깊은 인상을 남기며, 강남의 밤을 더욱 빛나게 하는 특별한 경험을 제공하여 방문객들이 언제나 잊지 못할 추억을 만들 수 있도록 합니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "01:00"},
			Part2: &TimeType{Has: true, Open: "01:00", Closed: "15:00"},
		},
		Menu:          &Menu{Part1Whisky: 160000, Part2Whisky: 130000, TC: 120000, RT: 50000},
		DatePublished: storeDate(2024, 1, 10),
		DateModified:  storeDate(2024, 1, 10),
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
		Description: "강남 달토 하이퍼블릭은 혁신적인 엔터테인먼트와 다양한 음료 선택으로 유흥을 찾는 방문객들에게 새로운 유형의 즐거움을 제공하는 것으로 유명하며, 강남의 밤문화를 새롭게 경험하고 싶은 이들에게 추천되는 장소입니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "01:00"},
			Part2: &TimeType{Has: true, Open: "01:00", Closed: "15:00"},
		},
		Menu:          &Menu{Part1Whisky: 170000, Part2Whisky: 140000, TC: 110000, RT: 50000},
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
		Description: "강남 런닝래빗 하이퍼블릭은 역동적인 분위기와 다채로운 이벤트로 가득 차 있어 강남에서 유흥을 즐기고자 하는 방문객들에게 지루할 틈 없는 밤을 선사하며, 언제나 새로운 것을 찾는 이들에게 에너지 넘치는 경험을 제공합니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "01:00"},
			Part2: &TimeType{Has: true, Open: "01:00", Closed: "15:00"},
		},
		Menu:          &Menu{Part1Whisky: 170000, Part2Whisky: 140000, TC: 110000, RT: 50000},
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
		Description: "강남 트렌드 하이퍼블릭은 최신 트렌드를 반영한 모던한 디자인과 엔터테인먼트로 강남의 젊은 층에게 큰 인기를 끌며, 유흥을 즐기려는 방문객들에게 강남에서의 밤을 멋지게 즐길 수 있는 다양한 프로그램과 이벤트를 제공합니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "01:00"},
			Part2: &TimeType{Has: true, Open: "01:00", Closed: "15:00"},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 110000, RT: 50000},
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
		Description: "강남 사라있네 하이퍼블릭은 독특한 콘셉트와 아늑한 분위기로 유흥주점을 찾는 방문객들에게 편안한 휴식과 즐거움을 동시에 제공하며, 강남에서 조용히 여유를 즐기고 싶은 이들에게 적합한 공간으로 자리잡고 있습니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "01:00"},
			Part2: &TimeType{Has: true, Open: "01:00", Closed: "15:00"},
		},
		Menu:          &Menu{Part1Whisky: 170000, Part2Whisky: 140000, TC: 110000, RT: 50000},
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
		Description: "강남 수목원 하이퍼블릭은 자연 친화적인 인테리어와 평온한 분위기로 유흥을 즐기면서도 휴식을 찾는 방문객들에게 도심 속에서도 자연과 함께하는 듯한 특별한 경험을 제공합니다, 이곳은 바쁜 일상에서 벗어나 진정한 휴식을 찾는 이들에게 안성맞춤입니다.",
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
		Description: "강남 퍼펙트 하이퍼블릭은 완벽한 서비스와 고급스러운 분위기로 유명한 유흥주점으로, 강남에서의 밤을 더욱 특별하게 만들고자 하는 방문객들에게 매력적인 선택이 됩니다. 이곳에서는 각종 파티와 모임을 위한 프라이빗한 공간과 함께 최상의 경험을 제공하여 방문객들이 만족스러운 시간을 보낼 수 있도록 합니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "01:00"},
			Part2: &TimeType{Has: true, Open: "01:00", Closed: "15:00"},
		},
		Menu:          &Menu{Part1Whisky: 200000, Part2Whisky: 160000, TC: 120000, RT: 50000},
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
		Description: "강남 워라벨 하이퍼블릭은 일과 삶의 균형을 중시하는 고급 유흥주점으로, 편안한 분위기에서 즐길 수 있는 다양한 엔터테인먼트와 서비스를 제공합니다. 직장인들에게 특히 인기가 있으며, 업무 후 휴식과 즐거움을 동시에 찾고 싶은 이들에게 이상적인 장소입니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "01:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 250000, Part2Whisky: 0, TC: 130000, RT: 50000},
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
		Description: "강남 방탄 하이퍼블릭은 최첨단 음향 시스템과 현대적인 디자인이 특징인 유흥주점으로, 방문객들에게 강남의 역동적인 밤문화를 경험할 수 있는 환경을 제공합니다. 이곳에서는 다양한 음악과 엔터테인먼트를 즐길 수 있어 강남에서의 밤을 더욱 즐겁게 만들어줍니다.",
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
		Description: "강남 하이킥 레깅스룸은 스포티한 매력과 활동적인 분위기로 강남의 유흥주점 중에서도 독특한 콘셉트를 가지고 있으며, 새로운 스타일의 밤문화를 경험하고 싶은 젊은 층에게 인기가 많습니다. 이곳에서는 친구들과 함께 신나는 음악과 함께 다이나믹한 밤을 보낼 수 있습니다.",
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
		Description: "강남 유니크 쩜오는 그 이름처럼 유니크한 인테리어와 맞춤형 서비스로 방문객들에게 개성적인 밤문화 경험을 제공하는 유흥주점입니다. 강남에서 찾기 힘든 독특한 분위기와 프로그램으로 방문객들에게 새롭고 특별한 밤을 선사합니다.",
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
		Description: "강남 831 쩜오는 현대적인 감각이 돋보이는 인테리어와 함께 다양한 엔터테인먼트를 제공하는 유흥주점으로, 강남에서 활기찬 밤을 보내고 싶은 이들에게 적합한 공간입니다. 이곳에서는 트렌디한 음악과 함께 멋진 밤을 즐길 수 있습니다.",
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
			Address:      "735-32",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.3760308056935!2d127.0341289!3d37.4990484!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca1560e5d6327%3A0x5c114aeb8260a643!2s735-32%20Yeoksam-dong%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1679397562888!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_DOT5,
		Title:       "에이원",
		Description: "강남 에이원 쩜오는 고급스러움과 세련미를 겸비한 유흥주점으로 강남에서 품격 있는 밤문화를 추구하는 이들에게 최적의 장소입니다. 이곳에서는 프리미엄 음료와 함께 전문적인 서비스를 경험할 수 있어, 기억에 남는 밤을 보내고 싶은 방문객들에게 적극 추천됩니다.",
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
		Description: "강남 미라클 쩜오는 환상적인 분위기와 매혹적인 음악이 조화를 이루는 유흥주점으로, 강남에서 마법 같은 밤을 경험하고 싶은 이들에게 완벽한 선택이 됩니다. 이곳에서는 독창적인 이벤트와 특별한 엔터테인먼트를 통해 방문객들에게 즐거움과 흥분을 선사합니다.",
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
		Description: "강남 오키도키 쩜오는 친근하고 즐거운 분위기가 특징인 유흥주점으로, 편안하게 즐길 수 있는 엔터테인먼트와 다양한 음료로 방문객들에게 휴식과 즐거움을 동시에 제공합니다. 강남에서 친구들과 함께 가볍게 즐기고 싶은 밤에 이상적인 장소입니다.",
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
			Dong:         "신사동",
			Address:      "561-30",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3164.5256239750274!2d127.0258308!3d37.5191051!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca3ecf7b91b35%3A0x90e6eb4e73a5644e!2s561-30%20Sinsa-dong%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1678606137375!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_DOT5,
		Title:       "인트로",
		Description: "강남 인트로 쩜오는 새로운 만남과 경험을 위한 현대적인 공간으로, 유흥주점을 찾는 방문객들에게 다양한 즐거움과 새로운 인사이트를 제공합니다. 이곳에서는 혁신적인 콘셉트와 엔터테인먼트로 강남의 밤을 더욱 풍부하게 만들어줍니다.",
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
		Description: "강남 머니볼 쩜오는 화려함과 사치가 넘치는 유흥주점으로, 독특한 엔터테인먼트와 고급 음료를 제공하여 강남에서 럭셔리한 밤문화를 경험하고 싶은 이들에게 적합합니다. 이곳에서는 고급스러운 분위기 속에서 특별한 시간을 보낼 수 있습니다.",
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
			Address:      "731-11",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.401460674176!2d127.0436794!3d37.498448499999995!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca401a6b8183b%3A0xcbcd58a8b2cb7c50!2s731%20Yeoksam-dong%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1678605118720!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_DOT5,
		Title:       "라이징",
		Description: "강남 라이징 쩜오는 강남에서 떠오르는 인기 유흥주점으로, 역동적인 분위기와 최신 트렌드를 반영한 엔터테인먼트를 제공합니다. 젊은 층에게 특히 인기가 많으며, 강남에서 활기찬 밤을 보내고 싶은 방문객들에게 추천되는 공간입니다.",
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
		Description: "강남 임팩트 쩜오는 강렬한 인상을 주는 유흥주점으로, 강남에서 독특하고 강렬한 밤문화를 경험하고 싶은 이들에게 적합한 곳입니다. 이곳에서는 파워풀한 음악과 함께 열정적인 밤을 보낼 수 있습니다.",
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
		Description: "강남 스테이 쩜오는 오랫동안 머물고 싶은 아늑한 분위기의 유흥주점으로, 편안한 분위기에서 즐길 수 있는 다양한 엔터테인먼트와 음료로 방문객들에게 잊지 못할 밤을 제공합니다. 강남에서 휴식과 즐거움을 동시에 찾는 이들에게 이상적입니다.",
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
		Description: "강남 킹스맨 쩜오는 고급스러운 서비스와 세련된 인테리어로 유명한 유흥주점으로, 강남에서 품격 있는 밤문화를 즐기고 싶은 이들에게 매력적인 선택입니다. 프리미엄 음료와 함께 고급스러운 밤을 보낼 수 있는 곳입니다.",
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
		Description: "강남 사운드 클럽은 최신 음악과 화려한 조명이 어우러진 강남의 인기 클럽으로, 다이내믹한 밤문화를 즐기고 싶은 방문객들에게 에너지 넘치는 경험을 제공합니다. 이곳에서는 최고의 DJ와 함께 신나는 밤을 보낼 수 있습니다.",
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
		Description: "강남 레이스 클럽은 스피디한 분위기와 함께 모터스포츠 테마를 즐길 수 있는 유니크한 클럽으로, 강남에서 색다른 밤문화를 경험하고 싶은 이들에게 적합한 곳입니다. 이곳에서는 역동적인 음악과 함께 흥미로운 밤을 보낼 수 있습니다.",
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
		Description: "강남 어게인 호빠는 친근하고 즐거운 분위기에서 다양한 엔터테인먼트를 즐길 수 있는 호스트 바로, 방문객들에게 개성 넘치는 서비스와 함께 즐거운 시간을 보낼 수 있는 장소입니다. 강남에서 친구들과의 모임이나 파티에 적합한 곳입니다.",
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
		Description: "강남 씨엔엔 호빠는 현대적인 디자인과 프로페셔널한 서비스가 특징인 호스트 바로, 강남에서 고급스러운 호스팅 문화를 경험하고 싶은 방문객들에게 추천되는 장소입니다. 이곳에서는 다양한 음료와 함께 즐거운 대화를 나눌 수 있습니다.",
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
		Description: "강남 세븐 노래방은 강남 지역에서 현대적인 인테리어와 최신 음향 시스템으로 유명한 곳입니다. 여기서는 친구들과 함께 즐거운 시간을 보내며 다양한 장르의 노래를 즐길 수 있어, 모든 연령대에 걸쳐 인기가 높습니다.",
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
		Description: "강남 심포니 노래방은 고급스러운 분위기와 프리미엄 서비스를 자랑하는 노래방으로, 강남에서 특별한 밤을 보내고 싶은 이들에게 완벽한 장소입니다. 이곳에서는 최신 곡부터 클래식까지 폭넓은 음악 선택과 함께 품격 있는 노래 경험을 할 수 있습니다.",
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
		Description: "강남 애플 노래방은 트렌디하고 밝은 분위기가 특징인 곳으로, 젊은 층에게 특히 인기가 많습니다. 최신 히트곡들로 항상 업데이트되는 음악 리스트와 함께 친구들과의 모임이나 파티에 이상적인 공간으로, 즐거운 시간을 보낼 수 있는 장소입니다.",
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
