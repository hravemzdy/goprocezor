package registry_providers

import (
	legalios "github.com/hravemzdy/golegalios"
	"github.com/hravemzdy/goprocezor/internal/types"
)

type IArticleCodeProvider interface {
	Code() types.ArticleCode
}

type IArticleSpecProvider interface {
	IArticleCodeProvider
	GetSpec(period legalios.IPeriod, version types.VersionCode) types.IArticleSpec
}

type ArticleSpecProvider struct {
	code types.ArticleCode
}

func (p ArticleSpecProvider) Code() types.ArticleCode {
	return p.code
}

func NewArticleProvider(code int32) ArticleSpecProvider {
	return ArticleSpecProvider{code: types.GetArticleCode(code)}
}

func NewArticleCodeProvider(code int32) ArticleSpecProvider {
	return ArticleSpecProvider{code: types.GetArticleCode(code)}
}
