package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"taka-api/internal/svc"
	"taka-api/internal/types"
)

type GetProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProfileLogic {
	return &GetProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProfileLogic) GetProfile(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.Response)
	info := types.BasicInfo{
		Name:             "Raymond",
		ResidingLocation: "Auckland",
		Tags:             []string{"#CEO", "#CFO", "#COO"},
		MatchingScore:    10,
		Personality:      "Unsure",
		EducationHistory: []types.Education{{
			Level:       9,
			Provider:    "The University of Auckland",
			Discription: "Majoring in Computer Science and Information System",
			Gpa:         "9.0/9.0",
		}},
		WorkHistory: []types.Work{{
			CompanyName: "Raymond & Co",
			Title:       "CEO",
			Description: "",
		}},
		OtherExperience: []types.OtherExperience{{Description: "Hong Kong Best Photographer Reward"}},
	}

	resp.BasicInfo = info
	return resp, nil
}
