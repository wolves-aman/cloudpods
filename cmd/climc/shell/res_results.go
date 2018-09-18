package shell

import (
	"yunion.io/x/jsonutils"
	"yunion.io/x/onecloud/pkg/mcclient"
	"yunion.io/x/onecloud/pkg/mcclient/modules"
	"yunion.io/x/onecloud/pkg/mcclient/options"
)

func init() {

	type ResResultsListOptions struct {
		options.BaseListOptions
		StatMonth string `help:"stat_month of the query"`
		StartDate string `help:"start_date of the query"`
		EndDate   string `help:"end_date of the query"`
		ResType   string `help:"res_type of the query"`
		Platform  string `help:"platform of the query"`
		ProjectId string `help:"project_id of the query"`
	}
	R(&ResResultsListOptions{}, "resresult-list", "List all res results ", func(s *mcclient.ClientSession, args *ResResultsListOptions) error {
		var params *jsonutils.JSONDict
		{
			var err error
			params, err = args.BaseListOptions.Params()
			if err != nil {
				return err

			}
		}

		if len(args.StatMonth) > 0 {
			params.Add(jsonutils.NewString(args.StatMonth), "stat_month")
		}
		if len(args.StartDate) > 0 {
			params.Add(jsonutils.NewString(args.StartDate), "start_date")
		}
		if len(args.EndDate) > 0 {
			params.Add(jsonutils.NewString(args.EndDate), "end_date")
		}
		if len(args.ResType) > 0 {
			params.Add(jsonutils.NewString(args.ResType), "res_type")
		}
		if len(args.Platform) > 0 {
			params.Add(jsonutils.NewString(args.Platform), "platform")
		}
		if len(args.ProjectId) > 0 {
			params.Add(jsonutils.NewString(args.ProjectId), "project_id")
		}

		result, err := modules.ResResults.List(s, params)
		if err != nil {
			return err
		}

		printList(result, modules.ResResults.GetColumns(s))
		return nil
	})
}
