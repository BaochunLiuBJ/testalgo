package weektwo

import (
	"strconv"
	"strings"
)

const (
	space = " "
	dot   = "."
)
func subdomainVisits(cpdomains []string) []string {
	domainVisitsMap := make(map[string]uint64, 0)
	if len(cpdomains) == 0 {
		return []string{}
	}
	for _, cpd := range cpdomains {
		ss := strings.Split(cpd, space)
		if len(ss) != 2 {
			continue
		}
		clickString := ss[0]
		fullDomain := ss[1]
		c, err := strconv.ParseUint(clickString, 10, 32)
		if err != nil {
			continue
		}
		if count, ok := domainVisitsMap[fullDomain]; ok {
			domainVisitsMap[fullDomain] =  count + c 
		} else {
			domainVisitsMap[fullDomain] = c 
		}
	}

	newDomainVisits := make(map[string]uint64, 0)
	for dm, cc := range domainVisitsMap {
		domainSlice := strings.Split(dm, dot)
		temp := make([]string, 0)
		for i := len(domainSlice) - 1; i >= 0; i-- {
			temp = append([]string{domainSlice[i]}, temp...)
			domainName := strings.Join(temp, ".")
			if count, ok := newDomainVisits[domainName]; ok {
				newDomainVisits[domainName] = count + cc
			} else {
				newDomainVisits[domainName] = cc
			}
		}
	}

	results := make([]string, 0)
	for domain, click := range newDomainVisits {
		t := strings.Join([]string{strconv.FormatUint(click, 10), domain}, " ")
		results = append(results, t)
	}

	return results
}
