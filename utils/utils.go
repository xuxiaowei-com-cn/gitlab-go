package utils

import (
	"fmt"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
	"log"
	"strconv"
	"strings"
)

const (
	Separator = "-"
)

func ParseID(id interface{}) (string, error) {
	switch v := id.(type) {
	case int:
		return strconv.Itoa(v), nil
	case string:
		return v, nil
	default:
		return "", fmt.Errorf("invalid ID type %#v, the ID must be an int or a string", id)
	}
}

func Unique(nums []int) []int {
	// 使用 map 来存储不重复的元素
	uniqueMap := make(map[int]bool)

	// 遍历数组，将元素作为键存入 map
	for _, num := range nums {
		uniqueMap[num] = true
	}

	// 从 map 中提取不重复的元素到新的切片
	var result []int
	for num := range uniqueMap {
		result = append(result, num)
	}

	return result
}

func RangeInt(stringSlices []string) []int {
	allowMin := 1
	var rangeInt []int
	for _, stringSlice := range stringSlices {
		if strings.Contains(stringSlice, Separator) {
			stringSplits := strings.Split(stringSlice, Separator)
			stringSplitsLen := len(stringSplits)
			if stringSplitsLen == 2 {
				a := stringSplits[0]
				b := stringSplits[1]

				if a != "" && b != "" {
					numA, err := strconv.Atoi(a)
					if err != nil {
						log.Printf("%s 转数字异常: %s\n", a, err.Error())
						continue
					}

					numB, err := strconv.Atoi(b)
					if err != nil {
						log.Printf("%s 转数字异常: %s\n", b, err.Error())
						continue
					}

					if numA > numB {

					}

					for i := numA; i <= numB; i++ {
						rangeInt = append(rangeInt, i)
					}

				} else if a == "" && b != "" {
					num, err := strconv.Atoi(b)
					if err != nil {
						log.Printf("%s 转数字异常: %s\n", b, err.Error())
					} else {
						if allowMin > num {
							log.Printf("%d 小于允许的最小值 %d 无法处理\n", num, allowMin)
						} else {
							for i := 1; i <= num; i++ {
								rangeInt = append(rangeInt, i)
							}
						}
					}
				} else if a != "" && b == "" {
					num, err := strconv.Atoi(a)
					if err != nil {
						log.Printf("%s 转数字异常: %s\n", a, err.Error())
					} else {
						if allowMin > num {
							log.Printf("%d 小于允许的最小值 %d 无法处理\n", num, allowMin)
						} else {
							for i := num; i <= num+flag.RangeMaxInterval; i++ {
								rangeInt = append(rangeInt, i)
							}
						}
					}
				} else {
					rangeInt = ToInt(b, rangeInt, allowMin)
				}
			} else {
				log.Printf("无法处理: %s\n", stringSlice)
			}
		} else {
			rangeInt = ToInt(stringSlice, rangeInt, allowMin)
		}
	}
	return rangeInt
}

func ToInt(string string, result []int, allowMin int) []int {
	num, err := strconv.Atoi(string)
	if err != nil {
		log.Printf("%s 转数字异常: %s\n", string, err.Error())
	} else {
		if allowMin > num {
			log.Printf("%d 小于允许的最小值 %d 无法处理\n", num, allowMin)
		} else {
			result = append(result, num)
		}
	}
	return result
}
