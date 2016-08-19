package util

import (
"github.com/asdqwex/aws_mgmt/describe"
"encoding/json"
)

func Get_regions() (region_list []string){
	regions := describe.Region()
	for _, value := range regions.Regions {
		region, _ := json.Marshal(value)
		var objmap map[string]string
		if err := json.Unmarshal(region, &objmap); err != nil {
	        panic(err)
	    }
		region_list = append(region_list, string(objmap["RegionName"]))
	}
	return
}