package models

import (
	"github.com/astaxie/beego/logs"
	"math"
	"strconv"
	"strings"
)

//TemplateGet func
func TemplateGet(page, limit, templates string) ([]Template, int64, error) {
	par := []string{"host", "name", "templateid"}
	hostspar := []string{"host", "name", "hostid"}
	rep, err := API.Call("template.get", Params{"output": par,
		"selectApplications": "count", "selectItems": "count",
		"selectTriggers": "count", "selectGraphs": "count",
		"selectDiscoveries": "count", "selectScreens": "count",
		"selectHosts": hostspar})
	if err != nil {
		logs.Error(err)
		return []Template{}, 0, err
	}
	hba, err := json.Marshal(rep.Result)
	if err != nil {
		logs.Error(err)
		return []Template{}, 0, err
	}

	var hb []Template
	err = json.Unmarshal(hba, &hb)
	if err != nil {
		logs.Error(err)
		return []Template{}, 0, err
	}
	var dt []Template
	var d Template

	if templates != "" {
		for _, v := range hb {
			if strings.Contains(v.Name, templates) {
				d.Host = v.Host
				d.Templateid = v.Templateid
				d.Name = v.Name
				d.Hosts = v.Hosts
				d.Applications = v.Applications
				d.Triggers = v.Triggers
				d.Items = v.Items
				d.Graphs = v.Graphs
				d.Screens = v.Screens
				d.Discoveries = v.Discoveries
				dt = append(dt, d)
			}
		}
	} else {
		for _, v := range hb {
			d.Host = v.Host
			d.Templateid = v.Templateid
			d.Name = v.Name
			d.Hosts = v.Hosts
			d.Applications = v.Applications
			d.Triggers = v.Triggers
			d.Items = v.Items
			d.Graphs = v.Graphs
			d.Screens = v.Screens
			d.Discoveries = v.Discoveries
			dt = append(dt, d)
		}
	}

	IntPage, err := strconv.Atoi(page)
	if err != nil {
		IntPage = 1
	}
	IntLimit, err := strconv.Atoi(limit)
	if err != nil {
		IntLimit = 10
	}
	//如果dt為空直接返回
	if len(dt) == 0 {
		return dt, int64(len(dt)), err
	}
	//分頁
	nums := len(dt)
	//page總數
	totalpages := int(math.Ceil(float64(nums) / float64(IntLimit)))
	if IntPage >= totalpages {
		IntPage = totalpages
	}
	if IntPage <= 0 {
		IntPage = 1
	}
	//結束頁數據
	var end int
	//begin 開始頁數據

	begin := (IntPage - 1) * IntLimit
	if IntPage == totalpages {
		end = nums
	}
	if IntPage < totalpages {
		end = IntPage * IntLimit
	} else {
		end = nums
	}
	//根據開始和結束返回數據列表
	var newtemplates []Template
	for i := begin; i < end; i++ {
		newtemplates = append(newtemplates, dt[i])
	}
	return newtemplates, int64(len(dt)), err
}

//TemplateAllGet func
func TemplateAllGet() ([]Template, int64, error) {
	par := []string{"host", "name", "templateid"}
	hostspar := []string{"host", "name", "hostid"}
	rep, err := API.Call("template.get", Params{"output": par,
		"selectApplications": "count", "selectItems": "count",
		"selectTriggers": "count", "selectGraphs": "count",
		"selectDiscoveries": "count", "selectScreens": "count",
		"selectHosts": hostspar})
	if err != nil {
		return []Template{}, 0, err
	}
	hba, err := json.Marshal(rep.Result)
	if err != nil {
		return []Template{}, 0, err
	}

	var hb []Template
	err = json.Unmarshal(hba, &hb)
	if err != nil {
		return []Template{}, 0, err
	}
	return hb, int64(len(hb)), nil
}
