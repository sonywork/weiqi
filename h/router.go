package h

import "strings"

const (
	MatchDefault = ""
	MatchParam = "*"
	PathSplit = "/"
)

type route struct {
	Pattern string
	Methods []string
	Handler Handler

	DefRoute   *route
	ParamRoute *route
	Routes     map[string]*route
}

func newRoute() *route {
	return &route{Routes: make(map[string]*route)}
}

func (this *route) Set(handler Handler, pattern string, methods ...string) {
	listPath := strings.Split(pattern, PathSplit)
	cr := this
	for i := range listPath {
		switch listPath[i] {
		//默认
		case MatchDefault:
			if cr.DefRoute == nil {
				cr.DefRoute = newRoute()
			}
			cr = cr.DefRoute
		//参数
		case MatchParam:
			if cr.ParamRoute == nil {
				cr.ParamRoute = newRoute()
			}
			cr = cr.ParamRoute
		//静态
		default:
			r, ok := cr.Routes[listPath[i]]
			if !ok {
				r = newRoute()
				cr.Routes[listPath[i]] = r
			}
			cr = r
		}
	}
	cr.Pattern = pattern
	cr.Methods = methods
	cr.Handler = handler
}

func (this *route) Match(pattern string) (*route, []string) {
	listPath := strings.Split(pattern, PathSplit)
	listParam := make([]string, 0)
	cr := this
	for i := range listPath {
		if listPath[i] == MatchDefault {
			cr = cr.DefRoute
		} else {
			r, ok := cr.Routes[listPath[i]]
			if ok {
				cr = r
			} else {
				listParam = append(listParam, listPath[i])
				cr = cr.ParamRoute
			}
		}
		if cr == nil {
			return nil, nil
		}
	}
	return cr, listParam
}