package templaterender

import (
	"github.com/answerdev/answer/internal/base/pager"
	"github.com/answerdev/answer/internal/schema"
	"golang.org/x/net/context"
)

func (q *TemplateRenderController) TagList(ctx context.Context, req *schema.GetTagWithPageReq) (resp *pager.PageModel, err error) {
	resp, err = q.tagService.GetTagWithPage(ctx, req)
	return resp, err
}
